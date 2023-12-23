package gotdx

import (
	"fmt"
	"gitee.com/quant1x/gotdx/quotes"
	"testing"
	"time"
)

func TestReOpen(t *testing.T) {
	api := GetTdxApi()
	v, _ := api.GetXdxrInfo("sh600072")
	fmt.Println(v)
	time.Sleep(2 * time.Second)
	ReOpen()
	v, _ = api.GetXdxrInfo("sh600072")
	fmt.Println(v)
	fmt.Println(quotes.NumberOfServers)
}
