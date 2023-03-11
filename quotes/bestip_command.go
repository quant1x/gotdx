package quotes

import (
	"gitee.com/quant1x/gotdx/proto/std"
	"log"
)

// Command 命令字
func Command(pool *ConnPool, factory std.Factory) std.Unmarshaler {
	conn := pool.GetConn()

	cli := conn.(*std.Client)
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

func CommandWithConn(cli *std.Client, callback std.Factory) std.Unmarshaler {
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
