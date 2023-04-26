package quotes

import (
	"github.com/mymmsc/gox/api"
	"github.com/mymmsc/gox/logger"
	"io"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

type TcpClient struct {
	sync.Mutex
	conn          net.Conn
	Addr          string    // 当前连接成功的服务器地址
	opt           Opt       // 参数
	complete      chan bool // 完成状态
	sending       chan bool // 正在发送状态
	done          chan bool // connection done
	completedTime time.Time // 时间戳
	timeMutex     sync.Mutex
}

type Opt struct {
	Servers           []Server      // 服务器组
	index             int           // 索引
	ConnectionTimeout time.Duration // 连接超时
	ReadTimeout       time.Duration // 读超时
	WriteTimeout      time.Duration // 写超时
	MaxRetryTimes     int           // 最大重试次数
	RetryDuration     time.Duration // 重试时间
}

func NewClient(opt Opt) *TcpClient {
	client := &TcpClient{}
	if opt.MaxRetryTimes <= 0 {
		opt.MaxRetryTimes = DefaultRetryTimes
	}
	if opt.ConnectionTimeout <= 0 {
		opt.ConnectionTimeout = CONN_TIMEOUT * time.Second
	}
	if opt.ReadTimeout <= 0 {
		opt.ReadTimeout = RECV_TIMEOUT * time.Second
	}
	if opt.WriteTimeout <= 0 {
		opt.WriteTimeout = RECV_TIMEOUT * time.Second
	}

	client.opt = opt
	client.sending = make(chan bool, 1)
	client.complete = make(chan bool, 1)
	client.updateCompletedTimestamp()
	return client
}

// 更新最后一次成功send/recv的时间戳
func (client *TcpClient) updateCompletedTimestamp() {
	client.completedTime = time.Now()
}

// 过去了多少秒
func (client *TcpClient) crossTime() (elapsedTime float64) {
	seconds := time.Since(client.completedTime).Seconds()
	return seconds
}

// 是否超时
func (client *TcpClient) hasTimedOut() bool {
	elapsedTime := client.crossTime()
	timeout := client.opt.ConnectionTimeout.Seconds()
	return elapsedTime >= timeout
}

// Command 执行通达信指令
func (client *TcpClient) Command(msg Message) error {
	defer client.Unlock()
	client.Lock()
	if client.conn == nil {
		return io.EOF
	}
	err := process(client.conn, msg, client.opt)
	if err != nil {
		return err
	}
	client.updateCompletedTimestamp()
	return nil
}

func (client *TcpClient) heartbeat() {
	defer func() {
		// 解析失败以后输出日志, 以备检查
		if err := recover(); err != nil {
			logger.Errorf("heartbeat.done error=%+v\n", err)
		}
	}()
	for {
		select {
		case <-time.After(time.Second):
			client.timeMutex.Lock()
			timeouted := client.hasTimedOut()
			client.timeMutex.Unlock()

			if timeouted {
				msg := NewSecurityCountPackage()
				msg.SetParams(&SecurityCountRequest{
					Market: uint16(1),
				})
				err := client.Command(msg)
				_ = err
				client.updateCompletedTimestamp()
				logger.Warnf("client -> server[%s]: heartbeat", client.Addr)
			}
		case <-client.done:
			logger.Warnf("client -> server[%s]: shutdown", client.Addr)
			return
		}
	}
}

// Connect 连接服务器
func (client *TcpClient) Connect() error {
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
		conn, err := net.DialTimeout("tcp", addr, client.opt.ConnectionTimeout) // net.DialTimeout()
		if err == nil {
			client.conn = conn
			client.Addr = addr
			client.updateCompletedTimestamp()
			opt.index += 1
			go client.heartbeat()
			break
		} else if i+1 >= total {
			opt.index = 0
			return err
		} else {
			opt.index += 1
		}
	}
	return nil
}

// Close 断开服务器
func (client *TcpClient) Close() error {
	close(client.sending)
	close(client.complete)
	api.CloseQuietly(client.conn)
	return nil
}
