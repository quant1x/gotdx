package internal

import (
	"fmt"
	"github.com/mymmsc/asio"
	"github.com/mymmsc/gox/pool"
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

type ConnPool struct {
	addr string
	pool pool.Pool
}

func NewConnPool_old(addr string, size int) *ConnPool {
	//ping 检测连接的方法
	//ping := redisPing
	//factory 创建连接的方法
	_factory := func() (interface{}, error) {
		fd, err := asio.Socket()
		if err == nil {
			err = asio.Connect(fd, addr)
		}
		if err == nil {
			_ = asio.Setsockopt(fd)
		}
		return fd, err
	}

	//close 关闭连接的方法
	_close := func(v interface{}) error {
		fd := v.(int)
		return asio.Close(fd)
	}

	//创建一个连接池： 初始化5，最大连接30
	poolConfig := &pool.Config{
		InitialCap: 5,
		MaxCap:     size,
		Factory:    _factory,
		Close:      _close,
		//Ping:       ping,
		//连接最大空闲时间，超过该时间的连接 将会关闭，可避免空闲时连接EOF，自动失效的问题
		IdleTimeout: CONN_TIMEOUT * time.Second,
	}
	_pool, err := pool.NewChannelPool(poolConfig)
	if err != nil {
		fmt.Println("err=", err)
	}
	cp := &ConnPool{
		addr: addr,
		pool: _pool,
	}
	return cp
}

func NewConnPool(addr string, size int) *ConnPool {
	//ping 检测连接的方法
	//ping := redisPing
	//factory 创建连接的方法
	_factory := func() (interface{}, error) {
		conn := conn_create(addr)
		return conn, nil
	}

	//close 关闭连接的方法
	_close := func(v interface{}) error {
		fd := v.(*Client)
		return conn_close(fd)
	}

	if size < POOL_INITED {
		size = POOL_INITED
	}

	//创建一个连接池： 初始化5，最大连接30
	poolConfig := &pool.Config{
		InitialCap: POOL_INITED,
		MaxCap:     size,
		MaxIdle:    size,
		Factory:    _factory,
		Close:      _close,
		//Ping:       ping,
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
	p.pool.Put(conn)
}

/*
func (p *ConnPool) MarkUnusable(context net.Conn) {
	if pc, ok := context.(*group.PoolConn); ok {
		pc.MarkUnusable()
	}
}*/

func conn_create(addr string) *Client {
	return NewClient2(addr)
}

func conn_close(client *Client) error {
	if client == nil {
		return nil
	}
	return client.Close()
}
