package internal

import "gitee.com/quant1x/exchange"

// BaseUnit 交易单位
//
//	A股、债券交易和债券买断式回购交易的申报价格最小变动单位为0.01元人民币
//	基金、权证交易为0.001元人民币
//	B股交易为0.001美元
//	债券质押式回购交易为0.005元
func defaultBaseUnit(marketId exchange.MarketType, code string) float64 {
	unit := 100.00
	if marketId == exchange.MarketIdShangHai {
		c := code[:2]
		switch c {
		case "51":
			unit = 1000.00
		}
	} else if marketId == exchange.MarketIdShenZhen {
		c := code[:3]
		switch c {
		case "159":
			unit = 1000.0
		}
	}

	return unit
}

type unitHandler func(marketId exchange.MarketType, code string) float64

var (
	BaseUnit unitHandler = defaultBaseUnit
)

func RegisterBaseUnitFunction(f unitHandler) {
	BaseUnit = f
}
