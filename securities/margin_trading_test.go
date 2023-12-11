package securities

import (
	"fmt"
	"testing"
)

func TestMarginTradingList(t *testing.T) {
	v := MarginTradingList()
	fmt.Println(v)
}
