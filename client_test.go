package gotdx

import (
	"fmt"
	"gotdx/proto"
	"testing"
)

var opt = &Opt{
	Host: "119.147.212.81",
	//Host: "58.63.254.191",
	//Host: "218.16.117.138",
	//Host: "222.85.139.177",
	Port: 7709,
}

func prepare() *Client {
	api := NewClient(opt)
	r, err := api.Connect()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)
	return api
}

func Test_tdx_Connect(t *testing.T) {
	api := NewClient(opt)
	reply, err := api.Connect()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(reply)

	_ = api.Disconnect()

}

func Test_tdx_GetSecurityCount(t *testing.T) {
	api := prepare()
	reply, err := api.GetSecurityCount(MarketSh)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(reply)

	reply, err = api.GetSecurityCount(MarketSz)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(reply)

	_ = api.Disconnect()

}

func Test_tdx_GetSecurityQuotes(t *testing.T) {
	api := prepare()
	params := []proto.Stock{}
	params = append(params, proto.Stock{Market: MarketSh, Code: "002062"})
	params = append(params, proto.Stock{Market: MarketSh, Code: "000001"})
	reply, err := api.GetSecurityQuotes(params)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(reply)

	_ = api.Disconnect()

}
