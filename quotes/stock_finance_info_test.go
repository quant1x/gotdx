package quotes

import (
	"encoding/json"
	"fmt"
	"gitee.com/quant1x/gotdx/proto"
	"github.com/mymmsc/gox/api"
	"testing"
)

func TestNewFinanceInfoPackage(t *testing.T) {
	stdApi, err := NewStdApi()
	if err != nil {
		panic(err)
	}
	defer stdApi.Close()
	sq1, err := stdApi.GetFinanceInfo(proto.MarketIdShangHai, "600005", 1)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", sq1)
	data, _ := json.Marshal(sq1)
	text := api.Bytes2String(data)
	fmt.Println(text)
}
