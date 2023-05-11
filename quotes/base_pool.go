package quotes

import (
	"fmt"
	"github.com/mymmsc/gox/logger"
	"github.com/mymmsc/gox/pool"
	"time"
)

const (
	// POOL_INITED 连接池初始化
	POOL_INITED = 1
	// POOL_MAX 连接池最大 2
	POOL_MAX = 5
	// CONN_TIMEOUT 链接超时 10 s
	CONN_TIMEOUT = 10
	// RECV_TIMEOUT 接收数据超时
	RECV_TIMEOUT = 2
)

// ConnPool 连接池
type ConnPool struct {
	addr string
	pool pool.Pool
}

// NewConnPool 创新一个新连接池
func NewConnPool(opt Opt, size int, factory func() (interface{}, error), close func(interface{}) error, ping func(interface{}) error) (*ConnPool, error) {
	if size < POOL_INITED {
		size = POOL_INITED
	}
	poolMax := POOL_MAX
	bestIpCount := len(opt.Servers)
	if bestIpCount == 0 {
		panic("No available hosts")
	}
	if bestIpCount < poolMax {
		poolMax = bestIpCount
	}
	// 创建一个连接池: 初始化5,最大连接30
	poolConfig := &pool.Config{
		InitialCap: POOL_INITED,
		MaxCap:     poolMax,
		MaxIdle:    size,
		Factory:    factory,
		Close:      close,
		Ping:       ping,
		//连接最大空闲时间，超过该时间的连接 将会关闭，可避免空闲时连接EOF，自动失效的问题
		IdleTimeout: CONN_TIMEOUT * time.Second,
	}
	_pool, err := pool.NewChannelPool(poolConfig)
	if err != nil {
		fmt.Println("err=", err)
		return nil, err
	}
	cp := &ConnPool{
		//addr: addr,
		pool: _pool,
	}
	return cp, nil
}

func (p *ConnPool) GetConn() interface{} {
	conn, err := p.pool.Get()
	if err != nil {
		logger.Errorf("获取连接失败", err)
		return nil
	}
	return conn
}

func (p *ConnPool) CloseConn(conn interface{}) error {
	return p.pool.Close(conn)
}

func (p *ConnPool) ReturnConn(conn interface{}) {
	_ = p.pool.Put(conn)
}

func (p *ConnPool) Close() {
	p.pool.Release()
}
