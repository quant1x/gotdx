package quotes

import (
	"gitee.com/quant1x/gox/runtime"
	"io"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"gitee.com/quant1x/gox/api"
	"gitee.com/quant1x/gox/exception"
	"gitee.com/quant1x/gox/logger"
)

var (
	ErrTimeOut = exception.New(1, "connect timeout")
)

type TcpClient struct {
	sync.Mutex              // 匿名属性
	conn          net.Conn  // tcp连接
	server        *Server   // 服务器信息
	opt           *Options  // 参数
	done          chan bool // connection done
	completedTime time.Time // 时间戳
	closed        uint32    // 关闭次数
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
		logger.Error("tcp连接失效")
		return io.EOF
	}
	err := process(client, msg)
	//errors.Is(err, net.OpError)
	//if _,ok:=err.( *net.OpError) ;ok{
	//	return nil,err
	//}
	if err != nil {
		logger.Error("业务处理失败", err)
		return err
	}
	client.updateCompletedTimestamp()
	return nil
}

func (client *TcpClient) heartbeat() {
	defer runtime.IgnorePanic("heartbeat.done")
	ticker := time.NewTicker(time.Second * 1)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			client.Lock()
			timedOut := client.hasTimedOut()
			client.Unlock()
			if timedOut {
				msg := NewSecurityCountPackage()
				msg.SetParams(&SecurityCountRequest{
					Market: uint16(1),
				})
				err := client.Command(msg)
				if err != nil {
					logger.Warnf("client -> server[%s]: error > shutdown", client.server)
					_ = client.Close()
					return
				} else {
					client.updateCompletedTimestamp()
					logger.Warnf("client -> server[%s]: heartbeat", client.server)
				}
				// 模拟服务器主动断开或者网络断开
				//logger.Warnf("client -> server[%s]: test force > shutdown", client.Addr)
				//_ = client.Close()
				//return
			}
		case <-client.done:
			logger.Warnf("client -> server[%s]: done > shutdown", client.server)
			return
		}
	}
}

// Connect 连接服务器
func (client *TcpClient) Connect(server *Server) error {
	addr := server.Addr()
	conn, err := net.DialTimeout("tcp", addr, client.opt.ConnectionTimeout) // net.DialTimeout()
	state := "connected"
	if err != nil {
		state = err.Error()
	}
	logger.Warnf("client -> server[%s]: %s", addr, state)
	if err == nil {
		client.conn = conn
		client.server = server
		client.updateCompletedTimestamp()
		go client.heartbeat()
	}

	if client.conn == nil {
		return ErrTimeOut
	}
	return nil
}

// Close 断开服务器
func (client *TcpClient) Close() error {
	defer runtime.IgnorePanic("TcpClient.Close")
	if atomic.LoadUint32(&client.closed) > 0 {
		return io.EOF
	}
	client.done <- true
	close(client.done)
	client.opt.releaseAddress(client.server)
	api.CloseQuietly(client.conn)
	atomic.AddUint32(&client.closed, 1)
	return nil
}
