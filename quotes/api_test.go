package quotes

import (
	"fmt"
	"gitee.com/quant1x/gotdx/proto"
	"testing"
)

func TestStdApi_ALL(t *testing.T) {
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

	// 3. finance_info
	fi, err := stdApi.GetFinanceInfo(proto.MarketShangHai, "600600", 1)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", fi)

	// 4. kline
	kl, err := stdApi.GetKLine(proto.MarketShenZhen, "002528", proto.KLINE_TYPE_RI_K, 0, 800)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("GetKLine: %+v\n", kl)

	// 5. stock list
	sl, err := stdApi.GetSecurityList(proto.MarketShangHai, 0)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("GetSecurityList: %+v\n", sl)

	// 6 index kline
	ikl, err := stdApi.GetIndexBars(proto.MarketShangHai, "000001", proto.KLINE_TYPE_RI_K, 0, 800)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("GetIndexBars: %+v\n", ikl)
	// 休眠20秒触发超时流程
	//time.Sleep(time.Second * 20)
	// 7. 获取指定市场内的证券数目
	sc, err := stdApi.GetSecurityCount(proto.MarketShangHai)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", sc)

	// 8.1 获取5档行情
	//sq, err := stdApi.GetSecurityQuotes([]uint8{proto.MarketShangHai, proto.MarketShangHai}, []string{"600030", "600600"})
	//sq1, err := stdApi.GetSecurityQuotes([]uint8{proto.MarketShangHai}, []string{"600600"})
	sq1, err := stdApi.GetSecurityQuotes([]uint8{proto.MarketShangHai}, []string{"880082"})
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", sq1)

	// 8.2 获取5档行情
	//sq2, err := stdApi.V2GetSecurityQuotes([]uint8{proto.MarketShangHai}, []string{"600600"})
	//sq2, err := stdApi.V2GetSecurityQuotes([]uint8{proto.MarketShangHai}, []string{"880082"})
	//sq2, err := stdApi.V2GetSecurityQuotes([]uint8{proto.MarketShangHai}, []string{"600600"})
	sq2, err := stdApi.V2GetSecurityQuotes([]uint8{proto.MarketShangHai, proto.MarketShangHai, proto.MarketShangHai}, []string{"600030", "600600", "880082"})
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", sq2)

	// 9. 分时数据
	mt, err := stdApi.GetMinuteTimeData(0, "159607")
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", mt)
	// 10. 历史分时
	hmt, err := stdApi.GetHistoryMinuteTimeData(proto.MarketShangHai, "600600", 20230113)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", hmt)
	// 11. 分笔成交
	td, err := stdApi.GetTransactionData(0, "000629", 0, 3800)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", td)
	// 12. 历史分笔成交
	htd, err := stdApi.GetHistoryTransactionData(proto.MarketShangHai, "600600", 20230210, 0, 10)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", htd)
	// 13. 除权除息
	xdxr, err := stdApi.GetXdxrInfo(proto.MarketShenZhen, "002528")
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
