package gotdx

import (
	"fmt"
	"testing"
)

func TestReOpen(t *testing.T) {
	api := GetTdxApi()
	v, _ := api.GetXdxrInfo("sh600072")
	fmt.Println(v)
	ReOpen()
	v, _ = api.GetXdxrInfo("sh600072")
	fmt.Println(v)
}
