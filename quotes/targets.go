package quotes

import (
	"math"
	"os"
	"path/filepath"

	"github.com/quant1x/exchange"
	"github.com/quant1x/exchange/cache"
	"github.com/quant1x/gotdx/internal"
	"github.com/quant1x/x/api"
	"github.com/quant1x/x/coroutine"
)

var (
	cacheSecurityCodeList = filepath.Join(cache.GetMetaPath(), "targets.csv")
)

type SecurityTarget struct {
	Code         string
	VolUnit      uint16
	DecimalPoint int8
	Name         string
}

var (
	__mapStockList  = map[string]SecurityTarget{} // 股票列表缓存
	__onceStockList coroutine.PeriodicOnce
)

// 读取股票列表缓存
func readCacheSecurityList() {
	filename := cacheSecurityCodeList
	if !api.FileExist(filename) {
		return
	}
	list := []SecurityTarget{}
	err := api.CsvToSlices(filename, &list)
	if err != nil || len(list) == 0 {
		return
	}
	for _, v := range list {
		code := v.Code
		__mapStockList[code] = v
	}
}

func writeCacheSecurityList(list []SecurityTarget) {
	filename := cacheSecurityCodeList
	api.SliceSort(list, func(a, b SecurityTarget) bool {
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
		__mapStockList[code] = SecurityTarget{Code: v.Code, VolUnit: v.VolUnit, DecimalPoint: v.DecimalPoint, Name: v.Name}
	}
	// 更新缓存
	newList := api.Values(__mapStockList)
	// 更新代码列表
	writeCacheSecurityList(newList)
}

// getSecurityList 证券列表
func getSecurityList() (allList []Security) {
	stdApi, err := NewStdApi()
	if err != nil {
		return
	}
	defer stdApi.Close()
	offset := uint16(TDX_SECURITY_LIST_MAX)
	start := uint16(0)
	for {
		reply, err := stdApi.GetSecurityList(exchange.MarketIdShangHai, start)
		if err != nil {
			return
		}
		for i := 0; i < int(reply.Count); i++ {
			security := &reply.List[i]
			security.Code = "sh" + security.Code
			//if exchange.AssertBlockBySecurityCode(&(security.Code)) {
			//	blk := GetBlockInfo(security.Code)
			//	if blk != nil {
			//		security.Name = blk.Name
			//	}
			//}
		}
		//list := api.Filter(reply.List, checkIndexAndStock)
		if len(reply.List) > 0 {
			allList = append(allList, reply.List...)
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
		//list := api.Filter(reply.List, checkIndexAndStock)
		if len(reply.List) > 0 {
			allList = append(allList, reply.List...)
		}
		if reply.Count < offset {
			break
		}
		start += reply.Count
	}

	return
}

// CheckoutSecurityInfo 获取证券信息
func CheckoutSecurityInfo(securityCode string) (*SecurityTarget, bool) {
	__onceStockList.Do(lazyLoadStockList)
	securityCode = exchange.CorrectSecurityCode(securityCode)
	security, ok := __mapStockList[securityCode]
	if ok {
		return &security, true
	}
	return nil, false
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

func init() {
	internal.RegisterBaseUnitFunction(SecurityBaseUnit)
}
