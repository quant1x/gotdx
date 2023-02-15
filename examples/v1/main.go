package main

import (
	"fmt"
	"gitee.com/quant1x/gotdx/proto"
	"gitee.com/quant1x/gotdx/proto/v1"
	"gitee.com/quant1x/gotdx/quotes"
	"log"
	"strconv"
	"strings"
	"unsafe"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate)
}

func main() {
	fmt.Println(unsafe.Sizeof(v1.FinanceInfo{}))
	//quotesSrv := config.GetBestStockQuotesServer()
	quotesSrv := quotes.Server{Host: "119.147.212.81", Port: 7709}
	//quotesSrvAddr := "106.120.74.86:7711" // quotesSrv.Addr()
	log.Println("正在连接到最优行情服务器: ", quotesSrv.Addr())
	T(quotesSrv.Host, quotesSrv.Port)
	//T("106.120.74.86", 7709)
}

func T(ip string, port int) {
	quotesSrv := quotes.Server{Host: ip, Port: port}
	//cp := internal.NewConnPool(quotesSrv.Addr(), 1, v1.ConnCreate, v1.ConnClose, nil)
	opt := quotes.Opt{Servers: []quotes.Server{quotesSrv}}
	cp, _ := quotes.NewConnPool(opt, 1, func() (interface{}, error) {
		s := strings.Join([]string{quotesSrv.Host, strconv.Itoa(quotesSrv.Port)}, ":")
		return v1.NewClient2(s)
	}, func(v interface{}) error {
		client := v.(*v1.Client)
		return client.Close()
	}, nil)
	// CMD信令 1
	quotes.Command(cp, func() (req v1.Marshaler, resp v1.Unmarshaler, err error) {
		req, resp, err = v1.NewSetupCmd1()
		return
	})
	// CMD信令 2
	quotes.Command(cp, func() (req v1.Marshaler, resp v1.Unmarshaler, err error) {
		req, resp, err = v1.NewSetupCmd2()
		return
	})
	// CMD信令 3
	quotes.Command(cp, func() (req v1.Marshaler, resp v1.Unmarshaler, err error) {
		req, resp, err = v1.NewSetupCmd3()
		return
	})
	// 查询股票数量
	quotes.Command(cp, func() (req v1.Marshaler, resp v1.Unmarshaler, err error) {
		req, resp, err = v1.NewSecurityCount(proto.MarketShangHai)
		return
	})
	// 查询股票列表
	quotes.Command(cp, func() (req v1.Marshaler, resp v1.Unmarshaler, err error) {
		req, resp, err = v1.NewGetSecurityList(proto.MarketShangHai, 255)
		return
	})
	// 查询个股基本面
	resp := quotes.Command(cp, func() (req v1.Marshaler, resp v1.Unmarshaler, err error) {
		req, resp, err = v1.NewFinanceInfo(proto.MarketShangHai, "600600")
		return
	})
	fmt.Println(resp)

	cp.Close()
}
