//go:build windows

package cache

const (
	// 数据工具如果用作于windows服务，获取到的用户信息是SYSTEM，无法定位到windows系统的用户宿主目录
	cache_default_path = "c:/.quant1x"
)
