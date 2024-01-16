package securities

import (
	"fmt"
	"testing"
)

func TestBlockList(t *testing.T) {
	v := BlockList()
	fmt.Println(v)
}

func TestParseAndGenerateBlockFile(t *testing.T) {
	parseAndGenerateBlockFile()
}

func TestGetBlockInfo(t *testing.T) {
	code := "880818"
	code = "881432"
	v := GetBlockInfo(code)
	fmt.Println(v)
}
