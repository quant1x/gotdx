package proto

import (
	"gitee.com/quant1x/gotdx/util"
	"strings"
)

type Market = uint8

const (
	MarketShenZhen Market = iota // 深圳
	MarketShangHai Market = 1    // 上海
	MarketBeiJing  Market = 2    // 北京
)

// GetMarket 判断股票ID对应的证券市场匹配规则
//
//	['50', '51', '60', '90', '110'] 为 sh
//	['00', '12'，'13', '18', '15', '16', '18', '20', '30', '39', '115'] 为 sz
//	['5', '6', '9'] 开头的为 sh， 其余为 sz
func GetMarket(symbol string) string {
	market := "sh"
	if util.StartsWith(symbol, []string{"sh", "sz", "SH", "SZ"}) {
		market = strings.ToLower(symbol[0:2])
	} else if util.StartsWith(symbol, []string{"50", "51", "60", "68", "90", "110", "113", "132", "204"}) {
		market = "sh"
	} else if util.StartsWith(symbol, []string{"00", "12", "13", "18", "15", "16", "18", "20", "30", "39", "115", "1318"}) {
		market = "sz"
	} else if util.StartsWith(symbol, []string{"5", "6", "9", "7"}) {
		market = "sh"
	} else if util.StartsWith(symbol, []string{"4", "8"}) {
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
