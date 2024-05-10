package quotes

import (
	"errors"
	"gitee.com/quant1x/gox/logger"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	ErrInvalidServerAddress = errors.New("invalid server address")
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
	ConnectionTimeout time.Duration        // 连接超时
	ReadTimeout       time.Duration        // 读超时
	WriteTimeout      time.Duration        // 写超时
	MaxRetryTimes     int                  // 最大重试次数
	RetryDuration     time.Duration        // 重试时间
	releaseAddress    func(server *Server) // 归还服务器地址回调函数
}

// StdApi 标准行情API接口
type StdApi struct {
	connPool *ConnPool   // 连接池
	opt      *Options    // 选项
	once     sync.Once   // 滑动窗口式Once
	servers  []Server    // 服务器组
	ch       chan Server // 服务器地址channel
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
		servers: srvs,
		opt:     &opt,
	}
	stdApi.ch = make(chan Server, POOL_MAX)
	//stdApi.once.SetOffsetTime(serverResetOffsetHours, serverResetOffsetMinutes)
	_factory := func() (any, error) {
		client := NewClient(stdApi.opt)
		server := stdApi.AcquireAddress()
		if server == nil {
			return nil, ErrInvalidServerAddress
		}
		err := client.Connect(server)
		if err != nil {
			stdApi.ReleaseAddress(server)
			return nil, err
		}
		stdApi.opt.releaseAddress = stdApi.ReleaseAddress
		err = stdApi.tdxHello1(client)
		if err != nil {
			_ = client.Close()
			return nil, err
		}
		err = stdApi.tdxHello2(client)
		if err != nil {
			_ = client.Close()
			return nil, err
		}
		return client, err
	}
	_close := func(v any) error {
		client := v.(*TcpClient)
		return client.Close()
	}
	_ping := func(v any) error {
		client := v.(*TcpClient)
		return stdApi.tdxPing(client)
	}
	maxCap := POOL_MAX
	bestIpCount := len(stdApi.servers)
	if bestIpCount == 0 {
		logger.Fatalf("no available hosts")
	}
	if maxCap > bestIpCount {
		maxCap = bestIpCount
	}

	maxIdle := maxCap

	halfCpuCount := runtime.NumCPU() / 2
	if maxIdle > halfCpuCount {
		maxIdle = halfCpuCount
	}

	cp, err := NewConnPool(maxCap, maxIdle, _factory, _close, _ping)
	if err != nil {
		return nil, err
	}
	stdApi.connPool = cp
	return &stdApi, nil
}

func (this *StdApi) Len() int {
	return len(this.servers)
}

func (this *StdApi) init() {
	//if this.inited.Load() == 1 {
	//	servs := GetFastHost(TDX_HOST_HQ)
	//	if len(servs) > 0 {
	//		this.servers = servs
	//	}
	//	// 关闭channel
	//	close(this.ch)
	//	// 读取剩余的服务地址
	//	for v := range this.ch {
	//		_ = v
	//	}
	//	this.ch = make(chan Server, this.Len())
	//}
	for _, v := range this.servers {
		this.ch <- v
	}
	//this.inited.Store(1)
}

// AcquireAddress 获取一个地址
func (this *StdApi) AcquireAddress() *Server {
	this.once.Do(this.init)
	// 非阻塞获取
	//srv, ok := <-this.ch
	logger.Warnf("获取一个服务器地址...begin")
	// 阻塞获取一个地址
	server := <-this.ch
	logger.Warnf("获取一个服务器地址...end")
	if len(server.Host) == 0 || server.Port == 0 {
		logger.Warnf("获取一个服务器地址...failed: nil")
		return nil
	}
	logger.Warnf("获取一个服务器地址...server=%s", server)
	return &server
}

// ReleaseAddress 返还一个地址
func (this *StdApi) ReleaseAddress(srv *Server) {
	logger.Warnf("返回一个服务器地址...")
	if srv == nil || len(srv.Host) == 0 || srv.Port == 0 {
		logger.Warnf("返回一个服务器地址...failed: nil")
		return
	}
	this.once.Do(this.init)
	logger.Warnf("返回一个服务器地址...server=%s, begin", *srv)
	// 阻塞返还一个地址
	this.ch <- *srv
	logger.Warnf("返回一个服务器地址...server=%s, end", *srv)
}

// NumOfServers 增加返回服务器IP数量
func (this *StdApi) NumOfServers() int {
	return len(this.servers)
}

func (this *StdApi) GetMaxIdleCount() int {
	return this.connPool.GetMaxIdleCount()
}

// Close 关闭
func (this *StdApi) Close() {
	this.connPool.CloseAll()
}

// 通过池关闭连接
func (this *StdApi) poolClose(cli *TcpClient) error {
	return this.connPool.CloseConn(cli)
}
