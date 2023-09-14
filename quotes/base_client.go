package quotes

import (
	"io"
	"net"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"gitee.com/quant1x/gox/api"
	"gitee.com/quant1x/gox/exception"
	"gitee.com/quant1x/gox/logger"
)

type TcpClient struct {
	sync.Mutex
	conn          net.Conn
	Addr          string     // 当前连接成功的服务器地址
	opt           *Options   // 参数
	complete      chan bool  // 完成状态
	sending       chan bool  // 正在发送状态
	done          chan bool  // connection done
	completedTime time.Time  // 时间戳
	timeMutex     sync.Mutex // 时间锁
	closed        uint32     // 关闭次数
}

// Server 主机信息
type Server struct {
	Name      string `json:"name"`
	Host      string `json:"host"`
	Port      int    `json:"port"`
	CrossTime int64  `json:"crossTime"`
}

func (s Server) Addr() string {
	return strings.Join([]string{s.Host, strconv.Itoa(s.Port)}, ":")
}

type Options struct {
	sync.Mutex
	Servers           []Server      // 服务器组
	index             int           // 索引
	ConnectionTimeout time.Duration // 连接超时
	ReadTimeout       time.Duration // 读超时
	WriteTimeout      time.Duration // 写超时
	MaxRetryTimes     int           // 最大重试次数
	RetryDuration     time.Duration // 重试时间
}

func NewClient(opt *Options) *TcpClient {
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
	client.done = make(chan bool, 1)
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
	client.Lock()
	defer client.Unlock()
	if client.conn == nil {
		logger.Errorf("tcp连接失效")
		return io.EOF
	}
	err := process(client, msg)
	//errors.Is(err, net.OpError)
	//if _,ok:=err.( *net.OpError) ;ok{
	//	return nil,err
	//}
	if err != nil {
		logger.Errorf("业务处理失败", err)
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
			timedOut := client.hasTimedOut()
			client.timeMutex.Unlock()
			if timedOut {
				msg := NewSecurityCountPackage()
				msg.SetParams(&SecurityCountRequest{
					Market: uint16(1),
				})
				err := client.Command(msg)
				if err != nil {
					logger.Warnf("client -> server[%s]: error > shutdown", client.Addr)
					_ = client.Close()
					return
				} else {
					client.updateCompletedTimestamp()
					logger.Warnf("client -> server[%s]: heartbeat", client.Addr)
				}
				// 模拟服务器主动断开或者网络断开
				//logger.Warnf("client -> server[%s]: test force > shutdown", client.Addr)
				//_ = client.Close()
				//return
			}
		case <-client.done:
			logger.Warnf("client -> server[%s]: done > shutdown", client.Addr)
			return
		}
	}
}

// Connect 连接服务器
func (client *TcpClient) Connect() error {
	client.opt.Lock()
	defer client.opt.Unlock()
	total := len(client.opt.Servers)
	if client.opt.index >= total {
		client.opt.index = 0
	}
	for i := client.opt.index; i < total; i++ {
		serv := client.opt.Servers[i]
		addr := strings.Join([]string{serv.Host, strconv.Itoa(serv.Port)}, ":")
		conn, err := net.DialTimeout("tcp", addr, client.opt.ConnectionTimeout) // net.DialTimeout()
		state := "connected"
		if err != nil {
			state = err.Error()
		}
		logger.Warnf("client -> server[%s]: %s", addr, state)
		if err == nil {
			client.conn = conn
			client.Addr = addr
			client.updateCompletedTimestamp()
			client.opt.index += 1
			go client.heartbeat()
			break
		} else if i+1 >= total {
			client.opt.index = 0
			return err
		} else {
			client.opt.index += 1
		}
	}
	if client.conn == nil {
		return exception.New(1, "connect timeout")
	}
	return nil
}

// Close 断开服务器
func (client *TcpClient) Close() error {
	defer func() {
		// 解析失败以后输出日志, 以备检查
		if err := recover(); err != nil {
			logger.Errorf("TcpClient.Close error=%+v\n", err)
		}
	}()
	if atomic.LoadUint32(&client.closed) > 0 {
		return io.EOF
	}
	client.done <- true
	close(client.done)
	close(client.sending)
	close(client.complete)
	api.CloseQuietly(client.conn)
	atomic.AddUint32(&client.closed, 1)
	return nil
}
