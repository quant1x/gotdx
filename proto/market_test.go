package proto

import (
	"fmt"
	"gitee.com/quant1x/pkg/testify/assert"
	"testing"
)

func TestGetSecurityCode(t *testing.T) {
	fmt.Println(GetSecurityCode(1, "600600"))
	fmt.Println(GetSecurityCode(0, "399001"))
	fmt.Println(GetSecurityCode(2, "399001"))
}

func TestGetMarket(t *testing.T) {
	code := "sh600600"
	fmt.Println(GetMarket(code))
	code = "sh.600600"
	fmt.Println(GetMarket(code))
	code = "600600.sh"
	fmt.Println(GetMarket(code))

	code = "880818"
	v := AssertBlockBySecurityCode(&code)
	fmt.Println(v)
}

func TestCorrectSecurityCode(t *testing.T) {
	correctedCode := CorrectSecurityCode("")
	assert.Equal(t, 0, len(correctedCode))
}
