package gotdx

import (
	"fmt"
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
	fmt.Println(api.NumOfServers())
}
