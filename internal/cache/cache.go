package cache

import (
	"gitee.com/quant1x/gox/util/homedir"
)

var (
	//__global_cache_path = cache_default_path // 数据根路径
	__global_cache_path = "~/.quant1x"
)

func init() {
	// 初始化缓存路径
	rootPath, err := homedir.Expand(__global_cache_path)
	if err != nil {
		panic(err)
	}
	__global_cache_path = rootPath
}

// DefaultCachePath 数据缓存的根路径
func DefaultCachePath() string {
	return __global_cache_path
}

// GetMetaPath 元数据缓存路径
func GetMetaPath() string {
	return DefaultCachePath() + "/meta"
}

// GetBlockPath 板块路径
func GetBlockPath() string {
	return GetMetaPath()
}
