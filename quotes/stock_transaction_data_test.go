package quotes

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/quant1x/x/api"
)

func TestTransaction(t *testing.T) {
	stdApi, err := NewStdApi()
	if err != nil {
		panic(err)
	}
	defer stdApi.Close()
	reply, err := stdApi.GetTransactionData("sh600010", 0, 2)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", reply)
	data, _ := json.Marshal(reply)
	text := api.Bytes2String(data)
	fmt.Println(text)
}

func TestHistoryTransaction(t *testing.T) {
	stdApi, err := NewStdApi()
	if err != nil {
		panic(err)
	}
	defer stdApi.Close()
	code := "sh880534"
	date := 20240110
	code = "sh600010"
	date = 20250417
	reply, err := stdApi.GetHistoryTransactionData(code, uint32(date), 0, 2)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", reply)
	data, _ := json.Marshal(reply)
	text := api.Bytes2String(data)
	fmt.Println(text)
}
