package securities

import (
	"gitee.com/quant1x/exchange"
	"gitee.com/quant1x/exchange/cache"
	"gitee.com/quant1x/gotdx"
	"gitee.com/quant1x/gotdx/internal"
	"gitee.com/quant1x/gotdx/quotes"
	"gitee.com/quant1x/gox/api"
	"gitee.com/quant1x/gox/coroutine"
	"math"
	"os"
	"path/filepath"
	"slices"
)

var (
	cacheSecurityCodeList = filepath.Join(cache.GetMetaPath(), "securities.csv")
)

var (
	__mapStockList  = map[string]quotes.Security{} // 股票列表缓存
	__onceStockList coroutine.PeriodicOnce
	__stock_list    = []string{}
)

// 读取股票列表缓存
func readCacheSecurityList() {
	filename := cacheSecurityCodeList
	if !api.FileExist(filename) {
		return
	}
	list := []quotes.Security{}
	err := api.CsvToSlices(filename, &list)
	if err != nil || len(list) == 0 {
		return
	}
	for _, v := range list {
		code := v.Code
		__mapStockList[code] = v
	}
	reloadCodeList()
}

func reloadCodeList() {
	list := api.Keys(__mapStockList)
	list = api.Unique(list)
	__stock_list = slices.Clone(list)
}

func writeCacheSecurityList(list []quotes.Security) {
	filename := cacheSecurityCodeList
	api.SliceSort(list, func(a, b quotes.Security) bool {
		return a.Code < b.Code
	})
	_ = api.SlicesToCsv(filename, list)
}

func lazyLoadStockList() {
	filename := cacheSecurityCodeList
	bUpdated := false
	if !api.FileExist(filename) {
		// 不存在需要创建
		bUpdated = true
	} else {
		// 存在, 先加载一次
		readCacheSecurityList()
		// 获取文件创建时间
		finfo, _ := os.Stat(filename)
		bUpdated = exchange.CanInitialize(finfo.ModTime())
	}
	if !bUpdated {
		return
	}
	list := getSecurityList()
	if len(list) == 0 {
		return
	}
	// 覆盖当前缓存
	for _, v := range list {
		code := v.Code
		__mapStockList[code] = v
	}
	// 更新缓存
	list = api.Values(__mapStockList)
	// 更新代码列表
	reloadCodeList()
	writeCacheSecurityList(list)
}

// getSecurityList 证券列表
func getSecurityList() (allList []quotes.Security) {
	stdApi := gotdx.GetTdxApi()
	offset := uint16(quotes.TDX_SECURITY_LIST_MAX)
	start := uint16(0)
	for {
		reply, err := stdApi.GetSecurityList(exchange.MarketIdShangHai, start)
		if err != nil {
			return
		}
		for i := 0; i < int(reply.Count); i++ {
			reply.List[i].Code = "sh" + reply.List[i].Code
		}
		list := api.Filter(reply.List, checkIndexAndStock)
		if len(list) > 0 {
			allList = append(allList, list...)
		}
		if reply.Count < offset {
			break
		}
		start += reply.Count
	}
	start = uint16(0)
	for {
		reply, err := stdApi.GetSecurityList(exchange.MarketIdShenZhen, start)
		if err != nil {
			return
		}
		for i := 0; i < int(reply.Count); i++ {
			reply.List[i].Code = "sz" + reply.List[i].Code
		}
		list := api.Filter(reply.List, checkIndexAndStock)
		if len(list) > 0 {
			allList = append(allList, list...)
		}
		if reply.Count < offset {
			break
		}
		start += reply.Count
	}

	return
}

// CheckoutSecurityInfo 获取证券信息
func CheckoutSecurityInfo(securityCode string) (*quotes.Security, bool) {
	__onceStockList.Do(lazyLoadStockList)
	securityCode = exchange.CorrectSecurityCode(securityCode)
	security, ok := __mapStockList[securityCode]
	if ok {
		return &security, true
	}
	return nil, false
}

// 检查指数和个股
func checkIndexAndStock(security quotes.Security) bool {
	if exchange.AssertIndexBySecurityCode(security.Code) {
		return true
	}
	if exchange.AssertStockBySecurityCode(security.Code) {
		return true
	}
	return false
}

// GetStockName 获取证券名称
func GetStockName(securityCode string) string {
	security, ok := CheckoutSecurityInfo(securityCode)
	if ok {
		return security.Name
	}
	return "Unknown"
}

// AllCodeList 获取全部证券代码
func AllCodeList() []string {
	__onceStockList.Do(lazyLoadStockList)
	return __stock_list
}

// SecurityBaseUnit 获取证券标价格的最小变动单位, 0.01返回100, 0.001返回1000
func SecurityBaseUnit(marketId exchange.MarketType, code string) float64 {
	securityCode := exchange.GetSecurityCode(marketId, code)
	securityInfo, ok := CheckoutSecurityInfo(securityCode)
	if !ok {
		return 100.00
	}
	return math.Pow10(int(securityInfo.DecimalPoint))
}

// SecurityPriceDigits 获取证券标的价格保留小数点后几位
//
//	默认范围2, 即小数点后2位
func SecurityPriceDigits(marketId exchange.MarketType, code string) int {
	securityCode := exchange.GetSecurityCode(marketId, code)
	securityInfo, ok := CheckoutSecurityInfo(securityCode)
	if !ok {
		return 2
	}
	return int(securityInfo.DecimalPoint)
}

func init() {
	internal.RegisterBaseUnitFunction(SecurityBaseUnit)
}
