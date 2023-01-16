package quotes

import (
	"fmt"
	"gitee.com/quant1x/gotdx/proto/market"
	"testing"
)

func TestStdApi_ALL(t *testing.T) {
	quotesSrv := Server{Host: "119.147.212.81", Port: 7709}
	stdApi := NewStdApi(quotesSrv)
	defer stdApi.Close()

	// 1. hello1
	hello1, err := stdApi.Hello1()
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", hello1)

	// 2. hello2
	hello2, err := stdApi.Hello2()
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", hello2)

	// 3. finance_info
	fi, err := stdApi.GetFinanceInfo(market.MarketShangHai, "600600")
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", fi)

	// 4. kline
	kl, err := stdApi.GetKLine(market.MarketShenZhen, "000002", market.KLINE_TYPE_RI_K, 0, 1)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", kl)

	// 5. stock list
	sl, err := stdApi.GetSecurityList(market.MarketShenZhen, 1)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", sl)

	// 6 index kline

	ikl, err := stdApi.GetIndexBars(market.MarketShangHai, "000001", market.KLINE_TYPE_RI_K, 0, 1)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", ikl)

	// 7. 获取指定市场内的证券数目
	sc, err := stdApi.GetSecurityCount(market.MarketShangHai)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", sc)

	// 8.
	sq, err := stdApi.GetSecurityQuotes([]uint8{uint8(market.MarketShangHai)}, []string{"600600"})
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", sq)
}
