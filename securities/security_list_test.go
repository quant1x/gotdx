package securities

import (
	"fmt"
	"testing"
)

func TestGetStockName(t *testing.T) {
	code := "sh510050"
	v := GetStockName(code)
	fmt.Println(v)
}

func TestAllCodeList(t *testing.T) {
	v := AllCodeList()
	fmt.Println(v)
}
