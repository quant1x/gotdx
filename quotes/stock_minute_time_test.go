package quotes

import (
	"encoding/json"
	"fmt"
	"gitee.com/quant1x/gox/api"
	"testing"
)

func TestStockMinuteTime(t *testing.T) {
	stdApi, err := NewStdApi()
	if err != nil {
		panic(err)
	}
	defer stdApi.Close()
	code := "sh600105"
	reply, err := stdApi.GetMinuteTimeData(code)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", reply)
	data, _ := json.Marshal(reply)
	text := api.Bytes2String(data)
	fmt.Println(text)
}

func TestStockMinuteTimeHistory(t *testing.T) {
	stdApi, err := NewStdApi()
	if err != nil {
		panic(err)
	}
	defer stdApi.Close()
	code := "sz000666"
	//code = "sh000001"
	//code = "sh510050"
	var date uint32 = 20230103
	reply, err := stdApi.GetHistoryMinuteTimeData(code, date)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", reply)
	data, _ := json.Marshal(reply)
	text := api.Bytes2String(data)
	fmt.Println(text)
}
