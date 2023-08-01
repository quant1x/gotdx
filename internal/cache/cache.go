package cache

import (
	"gitee.com/quant1x/gox/util/homedir"
	"runtime"
)

const (
	cacheRootPathOfWin = "c:/.quant1x"
)

var (
	default_cache_path = "~/.quant1x" // 数据根路径
)

func init() {
	// 初始化缓存路径
	switch runtime.GOOS {
	case "windows":
		default_cache_path = cacheRootPathOfWin
	}
	rootPath, err := homedir.Expand(default_cache_path)
	if err != nil {
		panic(err)
	}
	default_cache_path = rootPath
}

// DefaultCachePath 数据缓存的根路径
func DefaultCachePath() string {
	return default_cache_path
}

// GetMetaPath 元数据缓存路径
func GetMetaPath() string {
	return DefaultCachePath() + "/meta"
}

// GetBlockPath 板块路径
func GetBlockPath() string {
	return GetMetaPath()
}
