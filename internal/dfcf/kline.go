// 东方财富 数据接口

package dfcf

import (
	"gitee.com/quant1x/gox/api"
	"gitee.com/quant1x/gox/http"
	"gitee.com/quant1x/gox/logger"
	"gitee.com/quant1x/pkg/fastjson"
	urlpkg "net/url"
	"strings"
)

type EastmoneyApi struct{}

const (
	em_url = "http://33.push2his.eastmoney.com/api/qt/stock/kline/get"
)

// GetHistory sina获取历史数据的唯一方法
func GetHistory(fullCode string, datalen int) ([]DfcfHistory, error) {
	period := "daily"
	adjust := "qfq"
	params := urlpkg.Values{"secid": {"116." + fullCode[2:]},
		"ut":      {"fa5fd1943c7b386f172d6893dbfba10b"},
		"fields1": {"f1,f2,f3,f4,f5,f6"},
		"fields2": {"f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61"},
		"klt":     {period_dict[period]},
		"fqt":     {adjust_dict[adjust]},
		"end":     {"20500000"},
		"lmt":     {"1000000"},
		"_":       {"1623766962675"},
	}
	url := em_url + "?" + params.Encode()
	data, err := http.Get(url)
	if err != nil {
		logger.Errorf("%+v\n", err)
		return nil, err
	}
	// m := make(map[string]map[string]interface{})
	// err = fastjson.Unmarshal(data, &m)
	var kl []DfcfHistory
	obj, err := fastjson.ParseBytes(data)
	if err != nil {
		logger.Errorf("%+v\n", err)
		return kl, nil
	}
	errCode := obj.GetInt("rc")
	if errCode != 0 {
		logger.Errorf("%d: %s\n", err, obj.GetString("msg"))
		return kl, nil
	}
	_ = data
	biz := obj.Get("data")
	if biz == nil {
		logger.Errorf("数据非法\n")
		return kl, nil
	}

	history := biz.GetArray("klines")
	if history == nil {
		logger.Errorf("数据非法\n")
		return kl, nil
	}
	for _, item := range history {
		if item.Type() != fastjson.TypeString {
			continue
		}
		sb, err := item.StringBytes()
		if err != nil {
			logger.Fatalf("cannot obtain string: %s", err)
		}

		tmp := string(sb)
		hd := strings.Split(tmp, ",")
		var kl0 DfcfHistory
		err = api.Convert(hd, &kl0)
		if err == nil {
			kl = append(kl, kl0)
		}
	}
	//fmt.Printf("1. %+v\n", kl)
	return kl, nil
}

func (this *EastmoneyApi) Name() string {
	return "eastmoney"
}

//func (this *EastmoneyApi) CompleteKLine(code string) (kline []stock.DayKLine, err error) {
//	staticInfo, err := security.GetBasicInfo(code)
//	if err != nil {
//		return nil, err
//	}
//
//	//now := time.Now()
//	//now = utils.DateZero(now)
//	//now := utils.CanUpdateTime()
//	listTime := time.Unix(int64(staticInfo.ListTimestamp), 0)
//
//	// 计算需要补充多少年和多少天的数据
//	startTime := listTime
//	klines, _, err := this.DailyFromDate(code, startTime)
//	return klines, err
//}
//
//func (this *EastmoneyApi) DailyFromDate(code string, startTime time.Time) ([]stock.DayKLine, time.Time, error) {
//	staticInfo, err := security.GetBasicInfo(code)
//	if err != nil {
//		return nil, time.Time{}, err
//	}
//
//	now := utils.CanUpdateTime()
//	//now = utils.DateZero(now)
//	listTime := time.Unix(int64(staticInfo.ListTimestamp), 0)
//
//	// 计算需要补充多少年和多少天的数据
//	if listTime.After(startTime) {
//		startTime = listTime
//	}
//	days := utils.KLineRequireDays(now, startTime)
//	var kLines []stock.DayKLine
//	// 需要补充数据的最后一天
//	nextTradingDay := utils.DateZero(startTime)
//	// 测试时间比对
//	//nextTradingDay = time.Date(2021, 1, 4, 0, 0, 0, 0, time.Local)
//	history, err := GetHistory(code, days)
//	dls, lastDay := extract(nextTradingDay, history)
//	//nextTradingDay = lastDay
//	_ = lastDay
//	kLines = append(kLines, dls...)
//	listDay := listTime.Format(util.DateFormat)
//	startDay := startTime.Format(util.DateFormat)
//	endDay := now.Format(util.DateFormat)
//	logger.Infof("%s[%s]: %s -> %s", code, listDay, startDay, endDay)
//	return kLines, nextTradingDay, nil
//}
//
//// 转换行情数据为标准的K线数据
//func extract(nextTradingDay time.Time, history []DfcfHistory) ([]stock.DayKLine, time.Time) {
//	var kLines []stock.DayKLine
//	if len(history) > 0 {
//		for _, item := range history {
//			_lastDay, _ := utils.ParseTime(item.Date)
//			_lastDay = utils.DateZero(_lastDay)
//			if _lastDay.Before(nextTradingDay) {
//				continue
//			}
//			nextTradingDay = _lastDay.AddDate(0, 0, 1)
//			var dl stock.DayKLine
//			/*dl.Date = item.Date
//			dl.Open, _ = strconv.ParseFloat(item.Open, 64)
//			dl.High, _ = strconv.ParseFloat(item.High, 64)
//			dl.Low, _ = strconv.ParseFloat(item.Low, 64)
//			dl.Close, _ = strconv.ParseFloat(item.Close, 64)
//			dl.Volume, _ = strconv.ParseInt(item.Volume, 10, 64)
//			*/
//			api.Copy(&dl, &item)
//			kLines = append(kLines, dl)
//		}
//	}
//	return kLines, nextTradingDay
//}
