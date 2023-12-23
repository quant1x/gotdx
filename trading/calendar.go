package trading

import (
	"gitee.com/quant1x/gotdx/internal/cache"
	"gitee.com/quant1x/gotdx/internal/dfcf"
	"gitee.com/quant1x/gotdx/internal/js"
	"gitee.com/quant1x/gox/api"
	"gitee.com/quant1x/gox/coroutine"
	"gitee.com/quant1x/gox/http"
	"gitee.com/quant1x/gox/logger"
	"gitee.com/quant1x/pkg/gocsv"
	"os"
	"slices"
	"strings"
	"time"
)

const (
	urlSinaRealstockCompanyKlcTdSh = "https://finance.sina.com.cn/realstock/company/klc_td_sh.txt"
	urlSinaRealstockCompanyKlcTdSz = "https://finance.sina.com.cn/realstock/company/klc_td_sz.txt"
	TradingDayDateFormat           = "2006-01-02" // 交易日历日期格式
	calendarMissingDate            = "1992-05-04" // TODO:已知缺失的交易日期, 现在已经能自动甄别缺失的交易日期
)

var (
	__global_calendar_once coroutine.PeriodicOnce // 1D滚动锁
	__global_trade_dates   []string               // 交易日列表
)

// 非安全方式交易日历
func unsafeDates() []string {
	return __global_trade_dates
}

// 只读, 直接返回日历全部
func readOnlyDates() []string {
	__global_calendar_once.Do(resetCalendar)
	return __global_trade_dates
}

// 有修改需求的, 克隆一个副本
func cloneAllDates() []string {
	__global_calendar_once.Do(resetCalendar)
	return slices.Clone(__global_trade_dates)
}

// resetCalendar 重置日历
func resetCalendar() {
	bUpdate := updateCalendar()
	if bUpdate {
		noDates, err := checkCalendar()
		if err == nil && len(noDates) > 0 {
			calendarFilename := cache.CalendarFilename()
			_ = os.Remove(calendarFilename)
			updateCalendar(noDates...)
		}
	}
}

type calendar struct {
	Date   string `dataframe:"date"`
	Source string `dataframe:"source"`
}

// 加载交易日历, 数据源sina
func loadCalendarFromSina() {
	var list []calendar
	calendarFilename := cache.CalendarFilename()
	err := api.CsvToSlices(calendarFilename, &list)
	if err != nil && len(list) == 0 {
		return
	}
	for _, v := range list {
		__global_trade_dates = append(__global_trade_dates, v.Date)
	}
}

// 加载交易日历, 数据源内置
func loadCalendarFromCache() {
	var list []calendar
	reader := strings.NewReader(calendar2024)
	err := gocsv.Unmarshal(reader, &list)
	if err != nil && len(list) == 0 {
		return
	}
	for _, v := range list {
		__global_trade_dates = append(__global_trade_dates, v.Date)
	}
}

func loadCalendar() {
	loadCalendarFromCache()
}

// 尝试更新日历
func updateCalendar(noDates ...string) (bUpdate bool) {
	bUpdate = false
	bLoaded := false
	// 1. 检查日历文件是否存在
	calendarFilename := cache.CalendarFilename()
	if !api.FileExist(calendarFilename) {
		err := api.CheckFilepath(calendarFilename, true)
		if err != nil {
			panic("文件路径创建失败: " + calendarFilename)
		}
		bUpdate = true
	} else {
		loadCalendar()
		bLoaded = true
	}
	fileStat, err := api.GetFileStat(calendarFilename)
	var fileCreateTime time.Time
	if err == nil {
		fileCreateTime = fileStat.CreationTime
	}
	now := time.Now()
	today := now.Format(TradingDayDateFormat)
	toTm := now.Format(CN_SERVERTIME_FORMAT)
	fileDate := fileCreateTime.Format(TradingDayDateFormat)
	fileTm := fileCreateTime.Format(CN_SERVERTIME_FORMAT)
	if !bUpdate && unsafeDateIsTradingDay(today) {
		if fileDate < today && toTm >= CN_MarketInitTime {
			bUpdate = true
		} else if fileDate >= today && toTm >= CN_MarketInitTime && fileTm < CN_MarketInitTime {
			bUpdate = true
		}
	}

	// 如果不需要更新则直接加载缓存
	if !bUpdate {
		//if !bLoaded {
		//	loadCalendar()
		//}
		//return
	}

	header := map[string]any{
		//http.IfModifiedSince: fileCreateTime,
	}
	data, lastModified, err := http.Request(urlSinaRealstockCompanyKlcTdSh, http.MethodGet, "", header)
	if err != nil {
		panic("获取交易日历失败: " + urlSinaRealstockCompanyKlcTdSh)
	}
	if len(data) == 0 {
		// 如果没有数据, 则直接用缓存
		if !bLoaded {
			loadCalendar()
		}
		if !lastModified.IsZero() {
			err = os.Chtimes(calendarFilename, lastModified, lastModified)
			if err != nil {
				logger.Error(err)
			}
		}
		return
	}
	ret, err := js.SinaJsDecode(api.Bytes2String(data))
	if err != nil {
		panic("js解码失败: " + urlSinaRealstockCompanyKlcTdSh)
	}
	var dates []calendar
	for _, v := range ret.([]any) {
		ts := v.(time.Time)
		date := ts.Format(TradingDayDateFormat)
		e := calendar{
			Date:   date,
			Source: "sina",
		}
		dates = append(dates, e)
	}
	for _, v := range noDates {
		ts, _ := api.ParseTime(v)
		date := ts.Format(TradingDayDateFormat)
		e := calendar{
			Date:   date,
			Source: "tdx",
		}
		dates = append(dates, e)
	}
	dates = api.SliceUnique(dates, func(a, b calendar) int {
		if a.Date == b.Date {
			return 0
		}
		if a.Date < b.Date {
			return -1
		}
		return 1
	})

	err = api.SlicesToCsv(calendarFilename, dates)
	if err != nil {
		return
	}
	err = os.Chtimes(calendarFilename, now, lastModified)
	if err != nil {
		logger.Error(err)
	}
	for _, v := range dates {
		__global_trade_dates = append(__global_trade_dates, v.Date)
	}
	return
}

// 校验缺失的日期, 返回没有的日期列表
func checkCalendar() (noDates []string, err error) {
	dateList := getShangHaiTradeDates()
	// 校验日期的缺失
	if len(dateList) == 0 {
		return
	}
	start := dateList[0]
	end := dateList[len(dateList)-1]
	dest := TradeRange(start, end, false)
	noDates = []string{}
	for _, v := range dateList {
		found := slices.Contains(dest, v)
		if !found {
			noDates = append(noDates, v)
		}
	}
	return noDates, nil
}

// 获取上证指数的交易日期, 目的是校验日期
func getShangHaiTradeDates() (dates []string) {
	securityCode := "sh000001"
	klines, err := dfcf.A(securityCode)
	if err != nil {
		return
	}
	for _, v := range klines {
		dates = append(dates, v.Date)
	}
	return
}
