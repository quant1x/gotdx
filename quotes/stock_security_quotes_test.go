package quotes

import (
	"encoding/json"
	"fmt"
	"gitee.com/quant1x/gotdx/proto"
	"github.com/mymmsc/gox/api"
	"testing"
)

func TestSecurityQuotesPackage_ResponseFields(t *testing.T) {
	stdApi, err := NewStdApi()
	if err != nil {
		panic(err)
	}
	defer stdApi.Close()
	//sq1, err := stdApi.GetSecurityQuotes([]uint8{proto.MarketShangHai, proto.MarketShangHai, proto.MarketShangHai, proto.MarketShenZhen}, []string{"600275", "600455", "600086", "300742"})
	sq1, err := stdApi.GetSecurityQuotes([]uint8{proto.MarketShangHai, proto.MarketShangHai, proto.MarketShangHai, proto.MarketShenZhen}, []string{"600600", "600455", "600086", "300011"})
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", sq1)
	data, _ := json.Marshal(sq1)
	text := api.Bytes2String(data)
	fmt.Println(text)
}
