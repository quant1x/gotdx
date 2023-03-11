package std

import (
	"errors"
	"gitee.com/quant1x/gotdx/util"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

type Client struct {
	conn net.Conn
	addr string
	//Host          string
	//Port          int
	Timeout       time.Duration
	MaxRetryTimes int
	RetryDuration time.Duration
}

// NewClient 创建BaseClient实例
func NewClient(host string, port int) (*Client, error) {
	addr := strings.Join([]string{host, strconv.Itoa(port)}, ":")

	return NewClient2(addr)
}

func NewClient2(addr string) (*Client, error) {
	conn, err := net.Dial("tcp", addr) // net.DialTimeout()

	if err != nil {
		//log.Println(err)
		return nil, err
	}
	return &Client{
		conn: conn,
		addr: addr,
		//Host:          host,
		//Port:          port,
		MaxRetryTimes: 5,
		Timeout:       1 * time.Second,
		RetryDuration: time.Millisecond * 200,
	}, nil
}

func NewClientForTest(addr string) (*Client, error) {
	//defer func() {
	//	if info := recover(); info != nil {
	//		fmt.Println("触发了宕机", info)
	//	} else {
	//		fmt.Println("程序正常退出")
	//	}
	//}()
	conn, err := net.DialTimeout("tcp", addr, 1*time.Second) // net.DialTimeout()

	if err != nil {
		//log.Fatalln(err)
		log.Println(err)
		return nil, err
	}
	return &Client{
		conn: conn,
		addr: addr,
		//Host:          host,
		//Port:          port,
		MaxRetryTimes: 5,
		Timeout:       1 * time.Second,
		RetryDuration: time.Millisecond * 200,
	}, nil
}

func (cli *Client) Do(request Marshaler, response Unmarshaler) error {
	// 序列化请求
	req, err := request.Marshal()
	if err != nil {
		return err
	}
	// 发送请求
	retryTimes := 0
SEND:
	n, err := cli.conn.Write(req)
	// 重试
	if n < len(req) {
		retryTimes += 1
		if retryTimes <= cli.MaxRetryTimes {
			log.Printf("第%d次重试\n", retryTimes)
			goto SEND
		} else {
			return errors.New("数据未完整发送")
		}
	}
	if err != nil {
		return err
	}
	// 解析响应包头
	var header PacketHeader
	// 读取包头 大小为16个字节
	// 单次获取的字列流
	headerLength := 0x10
	headerBytes := make([]byte, headerLength)
	// 调用socket获取字节流并保存到data中
	headerBytes, err = cli.receive(headerLength)
	if err != nil {
		//log.Println(err)
		return err
	}
	err = header.Unmarshal(headerBytes)
	if err != nil {
		return err
	}
	// 根据获取响应体结构
	// 调用socket获取字节流并保存到data中
	bodyBytes, err := cli.receive(header.ZipSize)
	if err != nil {
		return err
	}
	// zlib解压缩
	if header.Compressed() {
		bodyBytes, err = util.ZlibUnCompress(bodyBytes)
	}
	// 反序列化为响应体结构
	err = response.Unmarshal(bodyBytes)
	if err != nil {
		return err
	}
	return nil
}

func (cli *Client) receive(length int) (data []byte, err error) {
	var (
		receivedSize int
	)
READ:
	tmp := make([]byte, length)
	// 设置读timeout
	err = cli.conn.SetReadDeadline(time.Now().Add(cli.Timeout))
	if err != nil {
		log.Println("setReadDeadline failed:", err)
	}
	// 调用socket获取字节流并保存到data中
	receivedSize, err = cli.conn.Read(tmp)
	// socket错误,可能为EOF
	if err != nil {
		return nil, err
	}
	// 数据添加到总输出,由于tmp申请内存时使用了length的长度，
	// 所以直接全部复制到data中会使得未完全传输的部分被填充为0导致数据获取不完整，
	// 故使用tmp[:receivedSize]
	data = append(data, tmp[:receivedSize]...)
	// 数据读满就可以返回了
	if len(data) == length {
		return
	}
	// 读取小于标准尺寸，说明到文件尾或者读取出现了问题没读满，可以返回了
	if receivedSize < length {
		goto READ
	}
	return
}

func (cli *Client) Close() error {
	return cli.conn.Close()
}
