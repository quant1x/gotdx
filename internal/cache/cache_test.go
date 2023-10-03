package cache

import (
	"fmt"
	"testing"
)

func TestCacheInit(t *testing.T) {
	InitCachePath("~/.quant2x")
	path := GetMetaPath()
	fmt.Println(path)
	InitCachePath("~/.quant3x")
	path = GetMetaPath()
	fmt.Println(path)
}
