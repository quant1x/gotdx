package quotes

import (
	"gitee.com/quant1x/gotdx/proto/v1"
	"log"
)

// Command 命令字
func Command(pool *ConnPool, factory v1.Factory) v1.Unmarshaler {
	conn := pool.GetConn()

	cli := conn.(*v1.Client)
	req, resp, err := factory()
	if err != nil {
		log.Fatal(err)
	}
	err = cli.Do(req, resp)
	if err != nil {
		log.Println(err)
		_ = cli.Close()
		return nil
	}
	//log.Println(resp)
	pool.ReturnConn(conn)
	return resp
}

func CommandWithConn(cli *v1.Client, callback v1.Factory) v1.Unmarshaler {
	req, resp, err := callback()
	if err != nil {
		log.Fatal(err)
	}
	err = cli.Do(req, resp)
	if err != nil {
		log.Println(err)
		_ = cli.Close()
		return nil
	}
	return resp
}
