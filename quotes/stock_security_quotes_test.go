package quotes

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/quant1x/exchange"
	"github.com/quant1x/x/api"
)

func TestSecurityQuotesPackage(t *testing.T) {
	stdApi, err := NewStdApi()
	if err != nil {
		panic(err)
	}
	defer stdApi.Close()
	//sq1, err := stdApi.GetSecurityQuotes([]uint8{proto.MarketIdShangHai, proto.MarketIdShangHai, proto.MarketIdShangHai, proto.MarketIdShenZhen}, []string{"600275", "600455", "600086", "300742"})
	sq1, err := stdApi.GetSecurityQuotes(
		[]uint8{exchange.MarketIdShangHai, exchange.MarketIdShangHai, exchange.MarketIdShangHai, exchange.MarketIdShangHai, exchange.MarketIdShenZhen},
		[]string{"000001", "000002", "880005", "880656", "399107"})
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", sq1)
	data, _ := json.Marshal(sq1)
	text := api.Bytes2String(data)
	fmt.Println(text)
}
