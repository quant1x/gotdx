package securities

import (
	"fmt"
	"gitee.com/quant1x/gotdx/internal/cache"
	"gitee.com/quant1x/gotdx/proto"
	"gitee.com/quant1x/gox/api"
	"gitee.com/quant1x/gox/coroutine"
	"slices"
)

const (
	// 两融配置文件
	marginTradingFilename = "margin-trading.csv"
	// https://data.eastmoney.com/rzrq/detail/all.html
)

type FinancingAndSecuritiesLendingTarget struct {
	Code string `name:"证券代码" dataframe:"code"`
}

var (
	onceMarginTrading      coroutine.PeriodicOnce
	cacheMarginTradingList []string
)

func lazyLoadMarginTrading() {
	target := cache.GetMetaPath() + "/" + marginTradingFilename
	if !api.FileExist(target) {
		filename := fmt.Sprintf("%s/%s", ResourcesPath, marginTradingFilename)
		_ = api.Export(resources, filename, target)
	}
	var tempList []FinancingAndSecuritiesLendingTarget
	if api.FileExist(target) {
		_ = api.CsvToSlices(target, &tempList)
	}
	var codes []string
	for _, v := range tempList {
		code := v.Code
		securityCode := proto.CorrectSecurityCode(code)
		codes = append(codes, securityCode)
	}
	if len(codes) > 0 {
		codes = api.SliceUnique(codes, func(a string, b string) int {
			if a < b {
				return -1
			} else if a > b {
				return 1
			} else {
				return 0
			}
		})
		cacheMarginTradingList = slices.Clone(codes)
	}
}

// MarginTradingList 获取两融标的列表
func MarginTradingList() []string {
	onceMarginTrading.Do(lazyLoadMarginTrading)
	return cacheMarginTradingList
}
