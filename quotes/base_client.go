package quotes

import (
	"io"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

type TcpClient struct {
	sync.Mutex
	conn     net.Conn
	opt      *Opt
	complete chan bool
	sending  chan bool
	timer    *time.Timer
}

type Opt struct {
	Servers       []Server      // 服务器组
	index         int           // 索引
	Timeout       time.Duration // 超时
	MaxRetryTimes int           // 最大重试次数
	RetryDuration time.Duration // 重试时间
}

func NewClient(opt *Opt) *TcpClient {
	client := &TcpClient{}
	if opt.MaxRetryTimes <= 0 {
		opt.MaxRetryTimes = DefaultRetryTimes
	}
	if opt.Timeout <= 0 {
		opt.Timeout = CONN_TIMEOUT * time.Second
	}

	client.opt = opt
	client.sending = make(chan bool, 1)
	client.complete = make(chan bool, 1)
	client.timer = time.NewTimer(opt.Timeout)
	return client
}

// Command 执行通达信指令
func (client *TcpClient) Command(msg Message) error {
	defer client.Unlock()
	client.Lock()
	opt := client.GetOpt()
	conn := client.GetConn()
	if conn == nil {
		return io.EOF
	}
	err := process(conn, msg, opt)
	if err != nil {
		//_ = this.poolClose(client)
		return err
	}
	return nil
}

func (client *TcpClient) heartbeat() {
	for {
		select {
		case <-client.timer.C:
			{
				msg := NewSecurityCountPackage()
				msg.SetParams(&SecurityCountRequest{
					Market: uint16(1),
				})
				_ = client.Command(msg)
				//fmt.Println(msg)
			}
			client.timer.Reset(client.opt.Timeout) // 每次使用完后需要人为重置下
		}
	}
}

// Connect 连接服务器
func (client *TcpClient) Connect() error {
	//defer func() {
	//	<-client.sending
	//}()
	//client.sending <- true
	defer client.Unlock()
	client.Lock()
	opt := client.opt
	total := len(opt.Servers)
	for i := opt.index; i < total; i++ {
		serv := opt.Servers[i]
		//if i < total {
		//	serv.Host = "127.0.0.1"
		//}
		addr := strings.Join([]string{serv.Host, strconv.Itoa(serv.Port)}, ":")
		conn, err := net.DialTimeout("tcp", addr, client.opt.Timeout) // net.DialTimeout()
		if err == nil {
			client.conn = conn
			break
		} else if i+1 >= total {
			opt.index = 0
			return err
		} else {
			opt.index += 1
		}
	}
	go client.heartbeat()
	return nil
}

// Close 断开服务器
func (client *TcpClient) Close() error {
	client.timer.Stop()
	close(client.sending)
	close(client.complete)
	return client.conn.Close()
}

func (client *TcpClient) GetConn() net.Conn {
	return client.conn
}

func (client *TcpClient) GetOpt() Opt {
	return *client.opt
}
