package v1

import (
	"fmt"
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

func TestV1Connect(t *testing.T) {
	api := NewClient(opt)
	reply, err := api.Connect()
	if err != nil {
		t.Errorf("error:%s", err)
	}
	fmt.Println(reply.Info)
	_ = api.Disconnect()

}

func TestV1GetSecurityCount(t *testing.T) {
	api := prepare()
	reply, err := api.GetSecurityCount(MARKET_SH)
	if err != nil {
		t.Errorf("error:%s", err)
	}
	fmt.Println(reply.Count)

	_ = api.Disconnect()
}

func TestV1GetSecurityQuotes(t *testing.T) {
	api := prepare()
	reply, err := api.GetSecurityQuotes([]uint8{MARKET_SH}, []string{"002062"})
	if err != nil {
		t.Errorf("error:%s", err)
	}

	for _, obj := range reply.List {
		fmt.Println(obj)
	}

	_ = api.Disconnect()

}

func TestV1GetSecurityList(t *testing.T) {
	api := prepare()
	reply, err := api.GetSecurityList(MARKET_SH, 0)
	if err != nil {
		t.Errorf("error:%s", err)
	}
	for _, obj := range reply.List {
		fmt.Println(obj)
	}

	_ = api.Disconnect()

}

func TestV1GetSecurityBars(t *testing.T) {
	// SecurityBars 与 MarketIndexBars 使用同一个接口靠market区分
	api := prepare()
	reply, err := api.GetSecurityBars(KLINE_TYPE_RI_K, 0, "000001", 0, 10)
	if err != nil {
		t.Errorf("error:%s", err)
	}

	for _, obj := range reply.List {
		fmt.Println(obj)
	}

	_ = api.Disconnect()

}

func TestV1GetIndexBars(t *testing.T) {
	// SecurityBars 与 MarketIndexBars 使用同一个接口靠market区分
	api := prepare()
	reply, err := api.GetIndexBars(KLINE_TYPE_RI_K, 1, "000001", 0, 10)
	if err != nil {
		t.Errorf("error:%s", err)
	}

	for _, obj := range reply.List {
		fmt.Println(obj)
	}

	_ = api.Disconnect()

}

func TestV1GetMinuteTimeData(t *testing.T) {
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

func TestV1GetHistoryMinuteTimeData(t *testing.T) {
	api := prepare()
	//reply, err := api.HistoryMinuteTimeData(20220511, 0, "159607")
	reply, err := api.GetHistoryMinuteTimeData(20230113, MARKET_SH, "600600")
	if err != nil {
		t.Errorf("error:%s", err)
	}

	for _, obj := range reply.List {
		fmt.Println(obj)
	}

	_ = api.Disconnect()

}

func TestV1GetTransactionData(t *testing.T) {
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

func TestV1GetHistoryTransactionData(t *testing.T) {
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

func TestV1GetFinanceInfo(t *testing.T) {
	api := prepare()
	reply, err := api.GetFinanceInfo(1, "600600")
	if err != nil {
		t.Errorf("error:%s", err)
	}

	fmt.Printf("%+v", reply)

	_ = api.Disconnect()

}
