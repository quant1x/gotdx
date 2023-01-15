package v2

import (
	"gitee.com/quant1x/gotdx/internal"
	"log"
)

// Command 命令字
func Command(pool *internal.ConnPool, factory Factory) Unmarshaler {
	conn := pool.GetConn()

	cli := conn.(*Client)
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

func CommandWithConn(cli *Client, callback Factory) Unmarshaler {
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
