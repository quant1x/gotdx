package proto

import (
	"github.com/mymmsc/gox/api"
	"strings"
)

type Market = uint8

const (
	MarketShenZhen Market = iota // 深圳
	MarketShangHai Market = 1    // 上海
	MarketBeiJing  Market = 2    // 北京
	MarketHongKong Market = 21   // 香港
	MarketUSA      Market = 22   // 美国

	StockDelisting = "DELISTING" // 退市
)

const (
	MARKET_SH string = "sh" // 上海
	MARKET_SZ string = "sz" // 深圳
	MARKET_BJ string = "bj" // 北京
	MARKET_HK string = "hk" // 香港
	MARKET_US string = "us" // 美国
)

// GetMarket 判断股票ID对应的证券市场匹配规则
//
//	['50', '51', '60', '90', '110'] 为 sh
//	['00', '12'，'13', '18', '15', '16', '18', '20', '30', '39', '115'] 为 sz
//	['5', '6', '9'] 开头的为 sh， 其余为 sz
func GetMarket(symbol string) string {
	market := "sh"
	if api.StartsWith(symbol, []string{"sh", "sz", "SH", "SZ"}) {
		market = strings.ToLower(symbol[0:2])
	} else if api.StartsWith(symbol, []string{"50", "51", "60", "68", "90", "110", "113", "132", "204"}) {
		market = "sh"
	} else if api.StartsWith(symbol, []string{"00", "12", "13", "18", "15", "16", "18", "20", "30", "39", "115", "1318"}) {
		market = "sz"
	} else if api.StartsWith(symbol, []string{"5", "6", "9", "7"}) {
		market = "sh"
	} else if api.StartsWith(symbol, []string{"88"}) {
		market = "sh"
	} else if api.StartsWith(symbol, []string{"4", "8"}) {
		market = "bj"
	}
	return market
}

// GetMarketId 获得市场ID
func GetMarketId(symbol string) uint8 {
	market := GetMarket(symbol)
	marketId := MarketShangHai
	if market == "sh" {
		marketId = MarketShangHai
	} else if market == "sz" {
		marketId = MarketShenZhen
	} else if market == "bj" {
		marketId = MarketBeiJing
	}
	return marketId
}

func GetMarketFlag(marketId Market) string {
	switch marketId {
	case MarketShenZhen:
		return MARKET_SZ
	case MarketBeiJing:
		return MARKET_BJ
	case MarketHongKong:
		return MARKET_HK
	case MarketUSA:
		return MARKET_US
	default:
		return MARKET_SH
	}
}

// AssertIndexByMarketAndCode 通过市场id和短码判断是否指数
func AssertIndexByMarketAndCode(marketId Market, code string) (isIndex bool) {
	if marketId == MarketShangHai && api.StartsWith(code, []string{"000", "880", "881"}) {
		return true
	} else if marketId == MarketShenZhen && api.StartsWith(code, []string{"399"}) {
		return true
	}
	return false
}
