package cache

import "fmt"

// CalendarFilename 交易日历文件路径
func CalendarFilename() string {
	filename := GetMetaPath() + "/calendar"
	return filename
}

// BlockFilename 板块缓存路径
func BlockFilename(ns ...string) string {
	// 默认取板块列表
	name := "blocks"
	if len(ns) > 0 {
		name = ns[0]
	}
	filename := fmt.Sprintf("%s/%s.csv", GetMetaPath(), name)
	return filename
}
