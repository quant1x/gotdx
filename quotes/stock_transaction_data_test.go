package quotes

import (
	"encoding/json"
	"fmt"
	"gitee.com/quant1x/gotdx/proto"
	"github.com/mymmsc/gox/api"
	"testing"
)

func TestNewTransactionPackage(t *testing.T) {
	stdApi, err := NewStdApi()
	if err != nil {
		panic(err)
	}
	defer stdApi.Close()
	reply, err := stdApi.GetTransactionData(proto.MarketIdShangHai, "600010", 0, 2)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", reply)
	data, _ := json.Marshal(reply)
	text := api.Bytes2String(data)
	fmt.Println(text)
}
