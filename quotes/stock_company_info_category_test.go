package quotes

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/quant1x/x/api"
)

func TestCompanyInfoCategoryPackage(t *testing.T) {
	stdApi, err := NewStdApi()
	if err != nil {
		panic(err)
	}
	defer stdApi.Close()
	reply, err := stdApi.GetCompanyInfoCategory("sh600977")
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", reply)
	data, _ := json.Marshal(reply)
	text := api.Bytes2String(data)
	fmt.Println(text)
}
