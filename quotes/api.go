package quotes

import (
	"strconv"
	"strings"
	"sync"
	"time"
)

// Server 主机信息
type Server struct {
	Source    string `json:"source"`
	Name      string `json:"name"`
	Host      string `json:"host"`
	Port      int    `json:"port"`
	CrossTime int64  `json:"crossTime"`
}

func (s Server) Addr() string {
	return strings.Join([]string{s.Host, strconv.Itoa(s.Port)}, ":")
}

func (s Server) String() string {
	//return fmt.Sprintf("%s[%s]: host=%s, port=%d", s.Source, s.Name, s.Host, s.Port)
	return s.Addr()
}

type Options struct {
	ConnectionTimeout time.Duration // 连接超时
	ReadTimeout       time.Duration // 读超时
	WriteTimeout      time.Duration // 写超时
	MaxRetryTimes     int           // 最大重试次数
	RetryDuration     time.Duration // 重试时间
	//index             int           // 索引
}

// StdApi 标准行情API接口
type StdApi struct {
	connPool   *ConnPool   // 连接池
	opt        *Options    // 选项
	sync.Mutex             // 锁
	sync.Once              // 一次性初始化锁
	Servers    []Server    // 服务器组
	ch         chan Server // 服务器地址channel
}

// NewStdApi 创建一个标准接口
func NewStdApi() (*StdApi, error) {
	server := GetFastHost(TDX_HOST_HQ)
	return NewStdApiWithServers(server)
}

// NewStdApiWithServers 通过服务器组创建一个标准接口
func NewStdApiWithServers(srvs []Server) (*StdApi, error) {
	opt := Options{
		ConnectionTimeout: CONN_TIMEOUT * time.Second,
	}
	stdApi := StdApi{
		Servers: srvs,
		opt:     &opt,
	}
	_factory := func() (interface{}, error) {
		client := NewClient(stdApi.opt)
		server := stdApi.Acquire()
		err := client.Connect(server)
		if err != nil {
			stdApi.Release(client.server)
			return nil, err
		}
		err = stdApi.tdxHello1(client)
		if err != nil {
			_ = client.Close()
			stdApi.Release(client.server)
			return nil, err
		}
		err = stdApi.tdxHello2(client)
		if err != nil {
			_ = client.Close()
			stdApi.Release(client.server)
			return nil, err
		}
		return client, err
	}
	_close := func(v interface{}) error {
		client := v.(*TcpClient)
		defer stdApi.Release(client.server)
		return client.Close()
	}
	_ping := func(v interface{}) error {
		client := v.(*TcpClient)
		return stdApi.tdxPing(client)
	}
	maxCap := POOL_MAX
	bestIpCount := len(stdApi.Servers)
	if bestIpCount == 0 {
		panic("No available hosts")
	}
	if bestIpCount < maxCap {
		maxCap = bestIpCount
	}
	maxIdle := 1
	cp, err := NewConnPool(maxCap, maxIdle, _factory, _close, _ping)
	if err != nil {
		return nil, err
	}
	stdApi.connPool = cp
	return &stdApi, nil
}

func (this *StdApi) Len() int {
	return len(this.Servers)
}

func (this *StdApi) init() {
	this.ch = make(chan Server, this.Len())
	for _, v := range this.Servers {
		this.ch <- v
	}
}

// Acquire 获取一个地址
func (this *StdApi) Acquire() Server {
	this.Once.Do(this.init)
	// 非阻塞获取
	//srv, ok := <-this.ch
	// 阻塞获取一个地址
	return <-this.ch
}

// Release 返还一个地址
func (this *StdApi) Release(srv Server) {
	this.Once.Do(this.init)
	this.ch <- srv
}

// NumOfServers 增加返回服务器IP数量
func (this *StdApi) NumOfServers() int {
	return len(this.Servers)
}

// Close 关闭
func (this *StdApi) Close() {
	this.connPool.CloseAll()
}

// 通过池关闭连接
func (this *StdApi) poolClose(cli *TcpClient) error {
	return this.connPool.CloseConn(cli)
}
