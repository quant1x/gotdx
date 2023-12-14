package securities

import (
	"gitee.com/quant1x/gotdx"
	"gitee.com/quant1x/gotdx/internal/cache"
	"gitee.com/quant1x/gotdx/proto"
	"gitee.com/quant1x/gotdx/quotes"
	"gitee.com/quant1x/gotdx/trading"
	"gitee.com/quant1x/gox/api"
	"gitee.com/quant1x/gox/coroutine"
	"golang.org/x/exp/maps"
	"os"
	"slices"
)

var (
	cacheSecurityCodeList = cache.GetMetaPath() + "/securities.csv"
)

var (
	__mapStockList  = map[string]quotes.Security{} // 股票列表缓存
	__onceStockList coroutine.PeriodicOnce
	__stock_list    = []string{}
)

var (
	// A股指数列表
	aShareIndexList = []string{
		"sh000001", // 上证综合指数
		"sh000002", // 上证A股指数
		"sh000905", // 中证500指数
		"sz399001", // 深证成份指数
		"sz399006", // 创业板指
		"sz399107", // 深证A指
		"sh880005", // 通达信板块-涨跌家数
		"sh510050", // 上证50ETF
		"sh510300", // 沪深300ETF
		"sh510900", // H股ETF
	}
)

// IndexList 指数列表
func IndexList() []string {
	return aShareIndexList
}

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
	list := maps.Keys(__mapStockList)
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
		bUpdated = trading.CanInitialize(finfo.ModTime())
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
	list = maps.Values(__mapStockList)
	// 更新代码列表
	reloadCodeList()
	writeCacheSecurityList(list)
}

// CheckoutSecurityInfo 获取证券名称
func CheckoutSecurityInfo(securityCode string) (*quotes.Security, bool) {
	__onceStockList.Do(lazyLoadStockList)
	securityCode = proto.CorrectSecurityCode(securityCode)
	security, ok := __mapStockList[securityCode]
	if ok {
		return &security, true
	}
	return nil, false
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

// 检查指数和个股
func checkIndexAndStock(security quotes.Security) bool {
	if proto.AssertIndexBySecurityCode(security.Code) {
		return true
	}
	if proto.AssertStockBySecurityCode(security.Code) {
		return true
	}
	return false
}

// getSecurityList 证券列表
func getSecurityList() (allList []quotes.Security) {
	stdApi := gotdx.GetTdxApi()
	offset := uint16(quotes.TDX_SECURITY_LIST_MAX)
	start := uint16(0)
	for {
		reply, err := stdApi.GetSecurityList(proto.MarketIdShangHai, start)
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
		reply, err := stdApi.GetSecurityList(proto.MarketIdShenZhen, start)
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
