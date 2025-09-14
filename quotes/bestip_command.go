package quotes

import (
	"fmt"

	"github.com/quant1x/gotdx/proto/std"
)

func CommandWithConn(cli *LabClient, callback std.Factory) (std.Unmarshaler, error) {
	req, resp, err := callback()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	err = cli.Do(req, resp)
	if err != nil {
		fmt.Println(err)
		_ = cli.Close()
		return nil, err
	}
	return resp, nil
}
