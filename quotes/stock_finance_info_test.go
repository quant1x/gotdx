package quotes

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/quant1x/x/api"
)

func TestNewFinanceInfoPackage(t *testing.T) {
	stdApi, err := NewStdApi()
	if err != nil {
		panic(err)
	}
	defer stdApi.Close()
	sq1, err := stdApi.GetFinanceInfo("bj920116")
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", sq1)
	data, _ := json.Marshal(sq1)
	text := api.Bytes2String(data)
	fmt.Println(text)
}
