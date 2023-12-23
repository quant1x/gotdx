package trading

import (
	"fmt"
	"gitee.com/quant1x/gox/api"
	"gitee.com/quant1x/gox/http"
	"testing"
	"time"
)

func Test_resetCalendar(t *testing.T) {
	resetCalendar()
}

func TestDowndata(t *testing.T) {
	header := map[string]any{
		//http.IfModifiedSince: fileModTime,
	}
	data, lastModified, err := http.Request(urlSinaRealstockCompanyKlcTdSh, http.MethodGet, "", header)
	if err != nil {
		panic("获取交易日历失败: " + urlSinaRealstockCompanyKlcTdSh)
	}
	fmt.Println(data)
	fmt.Println(lastModified, err)
}

func Test_updateHoliday(t *testing.T) {
	updateCalendar()
}

func TestIsHoliday(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "周末",
			args: args{date: "2023-02-18"},
			want: true,
		},
		{
			name: "周末",
			args: args{date: "2023-02-19"},
			want: true,
		},
		{
			name: "春节",
			args: args{date: "2023-01-23"},
			want: true,
		},
		{
			name: "工作日",
			args: args{date: "2023-02-20"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsHoliday(tt.args.date); got != tt.want {
				t.Errorf("IsHoliday() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTradeRange1(t *testing.T) {
	ds := TradeRange("2023-04-29", "2023-05-03")
	fmt.Println(len(ds))
	for _, v := range ds {
		fmt.Println(v)
	}
}

func TestUnique(t *testing.T) {
	a := []int{4, 1, 2, 1, 2, 3, 3, 3}
	a = api.Unique(a)
	fmt.Println(a)
}

func Test_checkCalendar(t *testing.T) {
	dates, err := checkCalendar()
	fmt.Println(dates, err)
}

func TestGetShangHaiTradeDates(t *testing.T) {
	dates := getShangHaiTradeDates()
	fmt.Println(dates)
}

func TestOnce(t *testing.T) {
	count := 1000
	for i := 0; i < count; i++ {
		lastDate := LastTradeDate()
		fmt.Println(lastDate)
		time.Sleep(time.Second * 1)
	}
}
