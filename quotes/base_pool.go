package quotes

import (
	"fmt"
	"gitee.com/quant1x/gox/logger"
	"gitee.com/quant1x/gox/pool"
	"time"
)

const (
	// POOL_INITED 连接池初始化
	POOL_INITED = 1
	// POOL_MAX 连接池最大 2
	POOL_MAX = 10
	// CONN_TIMEOUT 链接超时 10 s
	CONN_TIMEOUT = 10
	// RECV_TIMEOUT 接收数据超时
	RECV_TIMEOUT = 5
)

// ConnPool 连接池
type ConnPool struct {
	addr    string
	pool    pool.Pool
	maxIdle int
}

// NewConnPool 创新一个新连接池
func NewConnPool(maxCap, maxIdle int, factory func() (any, error), close func(any) error, ping func(any) error) (*ConnPool, error) {
	if maxIdle < POOL_INITED {
		maxIdle = POOL_INITED
	}
	// 创建一个连接池: 初始化5,最大连接30
	poolConfig := &pool.Config{
		InitialCap: POOL_INITED,
		MaxCap:     maxCap,
		MaxIdle:    maxIdle,
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
		pool:    _pool,
		maxIdle: maxIdle,
	}
	return cp, nil
}

func (p *ConnPool) GetMaxIdleCount() int {
	return p.maxIdle
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

func (p *ConnPool) CloseAll() {
	p.pool.CloseAll()
}

func (p *ConnPool) Close() {
	p.pool.Release()
}
