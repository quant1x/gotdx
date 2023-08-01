package cache

import (
	"gitee.com/quant1x/gox/util/homedir"
	"runtime"
)

const (
	cacheRootPathOfWin  = "c:/.quant1x"
	cacheRootPathOfUinx = "~/.quant1x"
)

var (
	__default_cache_path = "~/.quant1x" // 数据根路径
)

func init() {
	// 初始化缓存路径
	switch runtime.GOOS {
	case "windows":
		__default_cache_path = cacheRootPathOfWin
	default:
		__default_cache_path = cacheRootPathOfUinx
	}
	rootPath, err := homedir.Expand(__default_cache_path)
	if err != nil {
		panic(err)
	}
	__default_cache_path = rootPath
}

// DefaultCachePath 数据缓存的根路径
func DefaultCachePath() string {
	return __default_cache_path
}

// GetMetaPath 元数据缓存路径
func GetMetaPath() string {
	return DefaultCachePath() + "/meta"
}

// GetBlockPath 板块路径
func GetBlockPath() string {
	return GetMetaPath()
}
