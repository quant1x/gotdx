package quotes

import (
	"fmt"
	"gitee.com/quant1x/gotdx/proto"
	"gitee.com/quant1x/gox/logger"
	"testing"
)

func TestStdApi_ALL(t *testing.T) {
	logger.SetLevel(logger.DEBUG)
	//quotesSrv := Server{Host: "119.147.212.81", Port: 7709}
	//stdApi := NewStdApi(quotesSrv)
	stdApi, err := NewStdApi()
	if err != nil {
		panic(err)
	}
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

	// 2.1 heartbeat
	heartBeat, err := stdApi.HeartBeat()
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	fmt.Printf("%+v\n", heartBeat)

	// 3. finance_info
	fi, err := stdApi.GetFinanceInfo("sh600600")
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", fi)

	// 4. kline
	kl, err := stdApi.GetKLine("sz002528", proto.KLINE_TYPE_RI_K, 0, 800)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("GetKLine: %+v\n", kl)

	// 5. stock list
	sl, err := stdApi.GetSecurityList(proto.MarketIdShangHai, 0)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("GetSecurityList: %+v\n", sl)

	// 6 index kline
	ikl, err := stdApi.GetIndexBars("sh000001", proto.KLINE_TYPE_RI_K, 0, 800)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("GetIndexBars: %+v\n", ikl)
	// 7. 获取指定市场内的证券数目
	sc, err := stdApi.GetSecurityCount(proto.MarketIdShangHai)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", sc)

	//time.Sleep(time.Second * 15)
	// 8.1 获取5档行情
	//sq, err := stdApi.GetSecurityQuotes([]uint8{proto.MarketIdShangHai, proto.MarketIdShangHai}, []string{"600030", "600600"})
	//sq1, err := stdApi.GetSecurityQuotes([]uint8{proto.MarketIdShangHai}, []string{"600600"})
	sq1, err := stdApi.GetSecurityQuotes([]uint8{proto.MarketIdShangHai}, []string{"688981"})
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", sq1)

	// 8.2 获取5档行情
	//sq2, err := stdApi.V2GetSecurityQuotes([]uint8{proto.MarketIdShangHai}, []string{"600600"})
	//sq2, err := stdApi.V2GetSecurityQuotes([]uint8{proto.MarketIdShangHai}, []string{"880082"})
	sq2, err := stdApi.V2GetSecurityQuotes([]uint8{proto.MarketIdShenZhen}, []string{"002423"})
	//sq2, err := stdApi.V2GetSecurityQuotes([]uint8{proto.MarketIdShangHai, proto.MarketIdShangHai, proto.MarketIdShangHai}, []string{"600030", "600600", "880082"})
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", sq2)

	// 9. 分时数据
	//mt, err := stdApi.GetMinuteTimeData("sz159607") // 数据异常
	mt, err := stdApi.GetMinuteTimeData("sh510050")
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", mt)
	// 10. 历史分时
	hmt, err := stdApi.GetHistoryMinuteTimeData("sh600600", 20230113)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", hmt)
	// 11. 分笔成交
	td, err := stdApi.GetTransactionData("sz000629", 0, 3800)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", td)
	// 12. 历史分笔成交
	htd, err := stdApi.GetHistoryTransactionData("sh600105", 20230421, 0, 1800)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", htd)
	// 13. 除权除息
	xdxr, err := stdApi.GetXdxrInfo("sz002528")
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", xdxr)
	// 14. 板块meta信息
	blkMeta, err := stdApi.GetBlockMeta(BLOCK_DEFAULT)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", blkMeta)

	// 15. 板块信息
	blkInfo, err := stdApi.GetBlockInfo(BLOCK_DEFAULT)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", blkInfo)
}
