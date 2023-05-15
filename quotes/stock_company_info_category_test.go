package quotes

import (
	"encoding/json"
	"fmt"
	"gitee.com/quant1x/gotdx/proto"
	"gitee.com/quant1x/gox/api"
	"testing"
)

func TestCompanyInfoCategoryPackage(t *testing.T) {
	stdApi, err := NewStdApi()
	if err != nil {
		panic(err)
	}
	defer stdApi.Close()
	reply, err := stdApi.GetCompanyInfoCategory(proto.MarketIdShangHai, "600977")
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", reply)
	data, _ := json.Marshal(reply)
	text := api.Bytes2String(data)
	fmt.Println(text)
}
