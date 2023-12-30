package trading

import (
	"fmt"
	"gitee.com/quant1x/gox/api"
	"gitee.com/quant1x/gox/http"
	"gitee.com/quant1x/gox/logger"
	"gitee.com/quant1x/gox/text/encoding"
	"regexp"
	"strings"
)

const (
	urlHoliday = "https://www.tdx.com.cn/url/holiday/"
)

// FinancialHoliday 金融假日
type FinancialHoliday struct {
	Date     string `name:"日期" array:"0"`
	Holiday  string `name:"节日" array:"1"`
	Country  string `name:"地区" array:"2"`
	Exchange string `name:"交易所" array:"3"`
}

var (
	holidayPattern = `<textarea id="data" style="display:none;">([\s\w\W]+)</textarea>`
	holidayRegexp  = regexp.MustCompile(holidayPattern)
)

var (
	mapFinancialHolidays = map[string]FinancialHoliday{}
)

func getHolidayFromTdx() {
	data, err := http.Get(urlHoliday)
	if err != nil {
		logger.Error(err)
		return
	}
	text := api.Bytes2String(data)
	encoder := encoding.NewDecoder("gbk")
	text = encoder.ConvertString(text)

	arr := holidayRegexp.FindStringSubmatch(text)
	if len(arr) != 2 {
		return
	}
	textHolidays := arr[1]
	arr = strings.Split(textHolidays, "\n")
	for _, v := range arr {
		v := strings.TrimSpace(v)
		if len(v) == 0 {
			continue
		}
		line := strings.Split(v, `|`)
		if len(line) < 4 {
			continue
		}
		var fh FinancialHoliday
		err = api.Convert(line, &fh)
		fmt.Println(err)
		fh.Date = FixTradeDate(fh.Date)
		fmt.Println(fh)
		break
	}
}
