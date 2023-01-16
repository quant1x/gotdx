package quotes

import (
	"net"
	"strconv"
	"strings"
	"time"
)

type TcpClient struct {
	conn     net.Conn
	opt      *Opt
	complete chan bool
	sending  chan bool
}

type Opt struct {
	Host          string
	Port          int
	Timeout       time.Duration
	MaxRetryTimes int
	RetryDuration time.Duration
}

func NewClient(opt *Opt) *TcpClient {
	client := &TcpClient{}
	if opt.MaxRetryTimes <= 0 {
		opt.MaxRetryTimes = DefaultRetryTimes
	}
	if opt.Timeout <= 0 {
		opt.Timeout = 1 * time.Second
	}

	client.opt = opt
	client.sending = make(chan bool, 1)
	client.complete = make(chan bool, 1)

	return client
}

// Connect 连接服务器
func (client *TcpClient) Connect() error {
	addr := strings.Join([]string{client.opt.Host, strconv.Itoa(client.opt.Port)}, ":")
	conn, err := net.DialTimeout("tcp", addr, client.opt.Timeout) // net.DialTimeout()
	if err != nil {
		return err
	}
	client.conn = conn
	return err
}

// Disconnect 断开服务器
func (client *TcpClient) Close() error {
	return client.conn.Close()
}

func (client *TcpClient) GetConn() net.Conn {
	return client.conn
}

func (client *TcpClient) GetOpt() Opt {
	return *client.opt
}
