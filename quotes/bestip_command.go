package quotes

import (
	"fmt"
	"gitee.com/quant1x/gotdx/proto/std"
)

func CommandWithConn(cli *LabClient, callback std.Factory) std.Unmarshaler {
	req, resp, err := callback()
	if err != nil {
		fmt.Println(err)
	}
	err = cli.Do(req, resp)
	if err != nil {
		fmt.Println(err)
		_ = cli.Close()
		return nil
	}
	return resp
}
