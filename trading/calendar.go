package trading

import (
	"gitee.com/quant1x/gotdx/internal/cache"
	"gitee.com/quant1x/gotdx/internal/dfcf"
	"gitee.com/quant1x/gotdx/internal/js"
	"gitee.com/quant1x/gotdx/proto"
	"gitee.com/quant1x/gox/api"
	"gitee.com/quant1x/gox/http"
	"gitee.com/quant1x/gox/logger"
	"golang.org/x/exp/slices"
	"os"
	"sort"
	"strings"
	"time"
)

const (
	urlSinaRealstockCompanyKlcTdSh = "https://finance.sina.com.cn/realstock/company/klc_td_sh.txt"
	TradingDayDateFormat           = "2006-01-02" // 交易日历日期格式
	TradeDateFilename              = "calendar"
	calendarMissingDate            = "1992-05-04" // TODO:已知缺失的交易日期, 现在已经能自动甄别缺失的交易日期
)

var (
	//dataOnce         sync.Once
	__global_trade_dates []string // 交易日列表
	//__calendarFilename   = ""     // 日历文件路径

)

func getCalendarFilename() string {
	return cache.DefaultCachePath() + "/" + TradeDateFilename
}

func init() {
	//initCachePath()
	bUpdate := updateCalendar()
	if bUpdate {
		noDates, err := checkCalendar()
		if err == nil && len(noDates) > 0 {
			calendarFilename := getCalendarFilename()
			_ = os.Remove(calendarFilename)
			updateCalendar(noDates...)
		}
	}
	tradingDay := GetCurrentlyDay()
	proto.CorrectTradingDay(tradingDay)
}

type Calendar struct {
	Date   string `dataframe:"date"`
	Source string `dataframe:"source"`
}

// 加载交易日历
func loadCalendar() {
	list := []Calendar{}
	calendarFilename := getCalendarFilename()
	err := api.CsvToSlices(calendarFilename, &list)
	if err != nil && len(list) == 0 {
		return
	}
	for _, v := range list {
		__global_trade_dates = append(__global_trade_dates, v.Date)
	}
}

func getAllDates() []string {
	return slices.Clone(__global_trade_dates)
}

// IsHoliday 是否节假日
func IsHoliday(date string) bool {
	dates := getAllDates()
	iRet, found := sort.Find(len(dates), func(i int) int {
		return strings.Compare(date, dates[i])
	})
	_ = iRet
	return !found
}

// 尝试更新日历
func updateCalendar(noDates ...string) (bUpdate bool) {
	bUpdate = false
	calendarFilename := getCalendarFilename()
	if !api.FileExist(calendarFilename) {
		err := api.CheckFilepath(calendarFilename, true)
		if err != nil {
			panic("文件路径创建失败: " + calendarFilename)
		}
		bUpdate = true
	} else {
		loadCalendar()
	}
	finfo, err := os.Stat(calendarFilename)
	var fileModTime time.Time
	//var fileCreateTime time.Time
	if err == nil {
		fileModTime = finfo.ModTime()
		//fileCreateTime = cache.GetCreateTime(finfo)
	}
	now := time.Now()
	today := now.Format(TradingDayDateFormat)
	toTm := now.Format(CN_SERVERTIME_FORMAT)
	fileDate := fileModTime.Format(TradingDayDateFormat)
	fileTm := fileModTime.Format(CN_SERVERTIME_FORMAT)
	if !bUpdate && DateIsTradingDay(today) {
		if fileDate < today && toTm >= CN_MarketInitTime {
			bUpdate = true
		} else if fileDate >= today && toTm >= CN_MarketInitTime && fileTm < CN_MarketInitTime {
			bUpdate = true
		}
	}

	// 如果不需要更新则直接加载缓存
	if !bUpdate {
		loadCalendar()
		return
	}

	header := map[string]any{
		http.IfModifiedSince: fileModTime,
	}
	data, lastModified, err := http.Request(urlSinaRealstockCompanyKlcTdSh, "get", header)
	if err != nil {
		panic("获取交易日历失败: " + urlSinaRealstockCompanyKlcTdSh)
	}
	if len(data) == 0 {
		loadCalendar()
		err = os.Chtimes(calendarFilename, now, now)
		if err != nil {
			logger.Error(err)
		}
		return
	}
	ret, err := js.SinaJsDecode(api.Bytes2String(data))
	if err != nil {
		panic("js解码失败: " + urlSinaRealstockCompanyKlcTdSh)
	}
	dates := []Calendar{}
	for _, v := range ret.([]any) {
		ts := v.(time.Time)
		date := ts.Format(TradingDayDateFormat)
		e := Calendar{
			Date:   date,
			Source: "sina",
		}
		dates = append(dates, e)
	}
	for _, v := range noDates {
		ts, _ := api.ParseTime(v)
		date := ts.Format(TradingDayDateFormat)
		e := Calendar{
			Date:   date,
			Source: "tdx",
		}
		dates = append(dates, e)
	}
	api.SliceUnique(&dates, func(i, j int) bool {
		var a, b = dates[i], dates[j]
		return a.Date < b.Date
	})

	err = api.SlicesToCsv(calendarFilename, dates)
	if err != nil {
		return
	}
	now = time.Now()
	err = os.Chtimes(calendarFilename, lastModified, now)
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
	dest := TradeRange(start, end)
	noDates = []string{}
	for _, v := range dateList {
		found := slices.Contains(dest, v)
		if !found {
			noDates = append(noDates, v)
		}
	}
	return noDates, nil
}

//// 获取上证指数的交易日期, 目的是校验日期
//func getShangHaiTradeDates() (dates []string) {
//	securityCode := "sh000001"
//	tdxApi, err := quotes.NewStdApi()
//	if err != nil {
//		return
//	}
//	history := make([]quotes.SecurityBar, 0)
//	step := uint16(quotes.TDX_SECURITY_BARS_MAX)
//	start := uint16(0)
//	hs := make([]quotes.SecurityBarsReply, 0)
//	for {
//		count := step
//		var data *quotes.SecurityBarsReply
//		var err error
//		data, err = tdxApi.GetIndexBars(securityCode, proto.KLINE_TYPE_RI_K, start, count)
//		if err != nil {
//			return
//		}
//		hs = append(hs, *data)
//		if data.Count < count {
//			// 已经是最早的记录
//			// 需要排序
//			break
//		}
//		start += count
//	}
//	hs = api.Reverse(hs)
//	for _, v := range hs {
//		history = append(history, v.ConstituentStocks...)
//	}
//	dates = []string{}
//	for _, v := range history {
//		date1 := v.DateTime
//		dt, _ := api.ParseTime(date1)
//		date1 = dt.Format(TradingDayDateFormat)
//		dates = append(dates, date1)
//	}
//
//	return dates
//}

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
