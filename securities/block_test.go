package securities

import (
	"fmt"
	"testing"
)

func TestBlockList(t *testing.T) {
	v := BlockList()
	fmt.Println(v)
}

func TestGetBlockInfo(t *testing.T) {
	code := "880818"
	v := GetBlockInfo(code)
	fmt.Println(v)
}
