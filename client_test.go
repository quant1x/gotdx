package gotdx

import (
	"fmt"
	"gitee.com/quant1x/gotdx/proto/v1"
	"testing"
)

var opt = &Opt{
	Host: "119.147.212.81",
	Port: 7709,
}

func prepare() *TcpClient {
	api := NewClient(opt)
	reply, err := api.Connect()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(reply.Info)
	return api
}

func Test_tdx_Connect(t *testing.T) {
	api := NewClient(opt)
	reply, err := api.Connect()
	if err != nil {
		t.Errorf("error:%s", err)
	}
	fmt.Println(reply.Info)
	_ = api.Disconnect()

}

func Test_tdx_GetSecurityCount(t *testing.T) {
	api := prepare()
	reply, err := api.GetSecurityCount(v1.MARKET_SH)
	if err != nil {
		t.Errorf("error:%s", err)
	}
	fmt.Println(reply.Count)

	_ = api.Disconnect()
}

func Test_tdx_GetSecurityQuotes(t *testing.T) {
	api := prepare()
	reply, err := api.GetSecurityQuotes([]uint8{v1.MARKET_SH}, []string{"002062"})
	if err != nil {
		t.Errorf("error:%s", err)
	}

	for _, obj := range reply.List {
		fmt.Println(obj)
	}

	_ = api.Disconnect()

}

func Test_tdx_GetSecurityList(t *testing.T) {
	api := prepare()
	reply, err := api.GetSecurityList(v1.MARKET_SH, 0)
	if err != nil {
		t.Errorf("error:%s", err)
	}
	for _, obj := range reply.List {
		fmt.Println(obj)
	}

	_ = api.Disconnect()

}

func Test_tdx_GetSecurityBars(t *testing.T) {
	// SecurityBars 与 MarketIndexBars 使用同一个接口靠market区分
	api := prepare()
	reply, err := api.GetSecurityBars(v1.KLINE_TYPE_RI_K, 0, "000001", 0, 10)
	if err != nil {
		t.Errorf("error:%s", err)
	}

	for _, obj := range reply.List {
		fmt.Println(obj)
	}

	_ = api.Disconnect()

}

func Test_tdx_GetIndexBars(t *testing.T) {
	// SecurityBars 与 MarketIndexBars 使用同一个接口靠market区分
	api := prepare()
	reply, err := api.GetIndexBars(v1.KLINE_TYPE_RI_K, 1, "000001", 0, 10)
	if err != nil {
		t.Errorf("error:%s", err)
	}

	for _, obj := range reply.List {
		fmt.Println(obj)
	}

	_ = api.Disconnect()

}

func Test_tdx_GetMinuteTimeData(t *testing.T) {
	api := prepare()
	reply, err := api.GetMinuteTimeData(0, "159607")
	if err != nil {
		t.Errorf("error:%s", err)
	}

	for _, obj := range reply.List {
		fmt.Println(obj)
	}

	_ = api.Disconnect()

}

func Test_tdx_GetHistoryMinuteTimeData(t *testing.T) {
	api := prepare()
	//reply, err := api.HistoryMinuteTimeData(20220511, 0, "159607")
	reply, err := api.GetHistoryMinuteTimeData(20230113, v1.MARKET_SH, "600600")
	if err != nil {
		t.Errorf("error:%s", err)
	}

	for _, obj := range reply.List {
		fmt.Println(obj)
	}

	_ = api.Disconnect()

}

func Test_tdx_GetTransactionData(t *testing.T) {
	api := prepare()
	//reply, err := api.HistoryMinuteTimeData(20220511, 0, "159607")
	reply, err := api.GetTransactionData(0, "000629", 0, 3800)
	if err != nil {
		t.Errorf("error:%s", err)
	}

	for _, obj := range reply.List {
		fmt.Println(obj)
	}

	_ = api.Disconnect()

}

func Test_tdx_GetHistoryTransactionData(t *testing.T) {
	api := prepare()
	//reply, err := api.HistoryMinuteTimeData(20220511, 0, "159607")
	reply, err := api.GetHistoryTransactionData(20230111, 1, "600600", 99, 2)
	if err != nil {
		t.Errorf("error:%s", err)
	}

	for _, obj := range reply.List {
		fmt.Println(obj)
	}

	_ = api.Disconnect()

}
