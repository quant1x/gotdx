package quotes

import (
	"encoding/json"
	"fmt"
	"gitee.com/quant1x/gox/api"
	"testing"
)

func TestSnapshotPackage(t *testing.T) {
	stdApi, err := NewStdApi()
	if err != nil {
		panic(err)
	}
	defer stdApi.Close()
	//sq1, err := stdApi.GetSecurityQuotes([]uint8{proto.MarketIdShangHai, proto.MarketIdShangHai, proto.MarketIdShangHai, proto.MarketIdShenZhen}, []string{"600275", "600455", "600086", "300742"})
	sq1, err := stdApi.GetSnapshot([]string{"sh000001", "600105", "880656", "880367"})
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", sq1)
	data, _ := json.Marshal(sq1)
	text := api.Bytes2String(data)
	fmt.Println(text)
}
