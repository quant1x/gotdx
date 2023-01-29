# gotdx
golang实现的一个通达信数据协议库

## 1. 概要
- 整合了[gotdx](https://github.com/bensema/gotdx.git)和[TdxPy](https://github.com/rainx/pytdx)
- 增加了连接池的功能
- 自动探测主机网络速度
- 调用简单

## 2. 编写测试代码
第一次运行时, 连接池会探测服务器网络速度会慢一些, 网络测速后会缓存到本地。
#### 2.1. 创建 tdx.go
```go
package tdx

import (
	"gitee.com/quant1x/gotdx/proto"
	"gitee.com/quant1x/gotdx/quotes"
	"strings"
)

var (
	stdApi *quotes.StdApi = nil
)

func prepare() *quotes.StdApi {
	if stdApi == nil {
		//srv := quotes.Server{
		//	Name:      "临时主机",
		//	Host:      "119.147.212.81",
		//	Port:      7709,
		//	CrossTime: 0,
		//}
		//stdApi = quotes.NewStdApi(srv)
		stdApi = quotes.NewStdApi2()
	}
	return stdApi
}

func startsWith(str string, prefixs []string) bool {
	if len(str) == 0 || len(prefixs) == 0 {
		return false
	}
	for _, prefix := range prefixs {
		if strings.HasPrefix(str, prefix) {
			return true
		}
	}
	return false
}

// 判断股票ID对应的证券市场匹配规则
//
// ['50', '51', '60', '90', '110'] 为 sh
// ['00', '12'，'13', '18', '15', '16', '18', '20', '30', '39', '115'] 为 sz
// ['5', '6', '9'] 开头的为 sh， 其余为 sz
func getStockMarket(symbol string) string {
	//:param string: False 返回市场ID，否则市场缩写名称
	//:param symbol: 股票ID, 若以 'sz', 'sh' 开头直接返回对应类型，否则使用内置规则判断
	//:return 'sh' or 'sz'

	market := "sh"
	if startsWith(symbol, []string{"sh", "sz", "SH", "SZ"}) {
		market = strings.ToLower(symbol[0:2])
	} else if startsWith(symbol, []string{"50", "51", "60", "68", "90", "110", "113", "132", "204"}) {
		market = "sh"
	} else if startsWith(symbol, []string{"00", "12", "13", "18", "15", "16", "18", "20", "30", "39", "115", "1318"}) {
		market = "sz"
	} else if startsWith(symbol, []string{"5", "6", "9", "7"}) {
		market = "sh"
	} else if startsWith(symbol, []string{"4", "8"}) {
		market = "bj"
	}
	return market
}

func getStockMarketId(symbol string) uint8 {
	market := getStockMarket(symbol)
	marketId := proto.MarketShangHai
	if market == "sh" {
		marketId = proto.MarketShangHai
	} else if market == "sz" {
		marketId = proto.MarketShenZhen
	} else if market == "bj" {
		marketId = proto.MarketBeiJing
	}
	//# logger.debug(f"market => {market}")

	return uint8(marketId)
}

// GetKLine 获取日K线
func GetKLine(code string, start uint16, count uint16) *quotes.SecurityBarsReply {
	api := prepare()

	marketId := getStockMarketId(code)
	data, _ := api.GetKLine(marketId, code, proto.KLINE_TYPE_RI_K, start, count)
	return data
}

```

#### 2.2. 测试 tdx_test.go
```go
package tdx

import (
	"fmt"
	"testing"
)

func TestGetKLine(t *testing.T) {
	data := GetKLine("000002", 0, 1)
	fmt.Println(data)
}

```
