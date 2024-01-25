package securities

import (
	"fmt"
	"testing"
)

func TestMarginTradingList(t *testing.T) {
	v1 := MarginTradingList()
	fmt.Println(v1)
	v2 := MarginTradingList()
	fmt.Println(v2)
}
