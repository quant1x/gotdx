package util

import (
	"fmt"
	"testing"
	"time"
)

func Test__format_time(t *testing.T) {
	fmt.Println(timeFromStr("073382"))
	fmt.Println(timeFromStr("14989631"))

	fmt.Println(TimeFromInt(73382))
	fmt.Println(TimeFromInt(14989631))

	fmt.Println(TimeFromInt(9804942))

	// 小时~毫秒
	timeStamp := 9804942
	//返回time对象
	tm := time.Unix(int64(timeStamp/1000), int64(timeStamp%1000))
	//返回string
	dateStr := tm.Format("2006/01/02 15:04:05.000")
	fmt.Printf("%-10s %-10T %s\n", "dateStr", dateStr, dateStr)
}
