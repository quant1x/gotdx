package cache

import (
	"gitee.com/quant1x/gox/util/homedir"
	"sync"
)

var (
	__global_cache_once   sync.Once      // 懒加载锁
	__global_cache_path   = "~/.quant1x" // 数据根路径
	__once_temporary_path = "~/.quant1x" // 临时路径
)

func initPath(path string) {
	finalPath, err := homedir.Expand(path)
	if err != nil {
		panic(err)
	}
	__once_temporary_path = path
	__global_cache_path = finalPath
}

// InitCachePath 公开给外部调用的初始化路径的函数
//
//	lazyInit和InitCachePath两者只能真正被调用一次
func InitCachePath(path string) {
	__global_cache_once.Do(func() {
		__once_temporary_path = path
		initPath(path)
	})
}

// 默认的初始化路径
func lazyInit() {
	initPath(__once_temporary_path)
}

// DefaultCachePath 数据缓存的根路径
func DefaultCachePath() string {
	__global_cache_once.Do(lazyInit)
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
