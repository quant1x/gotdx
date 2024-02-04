package securities

import (
	"fmt"
	"gitee.com/quant1x/exchange"
	"gitee.com/quant1x/gotdx/internal"
	"testing"
)

func TestGetStockName(t *testing.T) {
	code := "sh880635"
	v := GetStockName(code)
	fmt.Println(v)
}

func TestAllCodeList(t *testing.T) {
	v := AllCodeList()
	fmt.Println(v)
}

func TestBaseUnit(t *testing.T) {
	marketId := exchange.MarketIdShangHai
	code := "000001"
	v := internal.BaseUnit(marketId, code)
	fmt.Println(v)
}
