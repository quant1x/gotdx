package dfcf

import (
	"fmt"
	"gitee.com/quant1x/gotdx/proto"
	"gitee.com/quant1x/gox/api"
	json "gitee.com/quant1x/gox/fastjson"
	"gitee.com/quant1x/gox/http"
	"gitee.com/quant1x/gox/logger"
	urlpkg "net/url"
	"strings"
	"time"
)

const (
	kUrlEastMonryZhKLine = "http://push2his.eastmoney.com/api/qt/stock/kline/get"
)

var (
	// 复权
	adjust_dict = map[string]string{
		"qfq": "1",
		"hfq": "2",
		"nil": "0",
	}
	// 周期
	period_dict = map[string]string{
		"daily":   "101",
		"weekly":  "102",
		"monthly": "103",
	}
)

// K线历史
func stock_hist(marketId int, symbol string, args ...string) ([]byte, error) {
	period := "daily"
	start_date := "19700101"
	end_date := "20500101"
	adjust := "qfq"
	argc := len(args)
	if argc >= 1 {
		start_date = args[0]
	}
	if argc >= 2 {
		end_date = args[1]
	}
	if argc >= 3 {
		adjust = args[2]
	}

	timestamp := time.Now().UnixMilli()
	params := urlpkg.Values{
		"fields1": {"f1,f2,f3,f4,f5,f6"},
		"fields2": {"f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61,f116"},
		"ut":      {"7eea3edcaed734bea9cbfc24409ed989"},
		"klt":     {period_dict[period]},
		"fqt":     {adjust_dict[adjust]},
		"secid":   {fmt.Sprintf("%d.%s", marketId, symbol)},
		"beg":     {start_date},
		"end":     {end_date},
		"_":       {fmt.Sprint(timestamp)},
	}
	url := kUrlEastMonryZhKLine + "?" + params.Encode()
	data, err := http.Get(url)
	return data, err
}

// A 下载A股数据
func A(code string) ([]KLine, error) {
	marketId, _, symbol := proto.DetectMarket(code)
	data, err := stock_hist(int(marketId), symbol)
	var kl = []KLine{}
	obj, err := json.ParseBytes(data)
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
		if item.Type() != json.TypeString {
			continue
		}
		sb, err := item.StringBytes()
		if err != nil {
			logger.Fatalf("cannot obtain string: %s", err)
		}

		tmp := string(sb)
		hd := strings.Split(tmp, ",")
		var kl0 KLine
		err = api.Convert(hd, &kl0)
		if err == nil {
			kl = append(kl, kl0)
		}
	}
	return kl, nil
}
