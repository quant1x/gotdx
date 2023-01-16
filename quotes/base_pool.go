package quotes

import (
	"fmt"
	"github.com/mymmsc/gox/pool"
	"strconv"
	"strings"
	"time"
)

const (
	// 连接池初始化
	POOL_INITED = 1
	// 连接池最大 2
	POOL_MAX = 2
	// 链接超时 30 s
	CONN_TIMEOUT = 30
)

// ConnPool 连接池
type ConnPool struct {
	addr string
	pool pool.Pool
}

// NewConnPool2 创新一个新连接池
func NewConnPool2(host string, port int, size int, factory func(string) (interface{}, error), close func(interface{}) error, ping func(interface{}) error) *ConnPool {
	addr := strings.Join([]string{host, strconv.Itoa(port)}, ":")
	return NewConnPool(addr, size, factory, close, ping)
}

// NewConnPool 创新一个新连接池
func NewConnPool(addr string, size int, factory func(string) (interface{}, error), close func(interface{}) error, ping func(interface{}) error) *ConnPool {
	_factory := func() (interface{}, error) {
		return factory(addr)
	}

	if size < POOL_INITED {
		size = POOL_INITED
	}

	//创建一个连接池: 初始化5,最大连接30
	poolConfig := &pool.Config{
		InitialCap: POOL_INITED,
		MaxCap:     POOL_MAX,
		MaxIdle:    size,
		Factory:    _factory,
		Close:      close,
		Ping:       ping,
		//连接最大空闲时间，超过该时间的连接 将会关闭，可避免空闲时连接EOF，自动失效的问题
		IdleTimeout: CONN_TIMEOUT * time.Second,
	}
	_pool, err := pool.NewChannelPool(poolConfig)
	if err != nil {
		fmt.Println("err=", err)
		return nil
	}
	cp := &ConnPool{
		addr: addr,
		pool: _pool,
	}
	return cp
}

func (p *ConnPool) GetConn() interface{} {
	conn, err := p.pool.Get()
	if err != nil {
		return nil
	}
	return conn
}

func (p *ConnPool) ReturnConn(conn interface{}) {
	_ = p.pool.Put(conn)
}

func (p *ConnPool) Close() {
	p.pool.Release()
}
