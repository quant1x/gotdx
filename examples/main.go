package main

import (
	"fmt"
	"gitee.com/quant1x/gotdx/internal"
	"gitee.com/quant1x/gotdx/proto/v2"
	"log"
	"strconv"
	"strings"
	"unsafe"
)

type Server struct {
	Name string `json:"name"`
	IP   string `json:"ip"`
	Port int    `json:"port"`
}

func (srv *Server) Addr() string {
	return strings.Join([]string{srv.IP, strconv.Itoa(srv.Port)}, ":")
}

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate)
}

func main() {
	fmt.Println(unsafe.Sizeof(v2.FinanceInfo{}))
	//quotesSrv := config.GetBestStockQuotesServer()
	quotesSrv := Server{IP: "119.147.212.81", Port: 7709}
	//quotesSrvAddr := "106.120.74.86:7711" // quotesSrv.Addr()
	log.Println("正在连接到最优行情服务器: ", quotesSrv.Addr())
	T(quotesSrv.IP, quotesSrv.Port)
	//T("106.120.74.86", 7709)
}
func T(ip string, port int) {
	quotesSrv := Server{IP: ip, Port: port}
	cp := internal.NewConnPool(quotesSrv.Addr(), 1)

	// CMD信令 1
	internal.Command(cp, func() (req v2.Marshaler, resp v2.Unmarshaler, err error) {
		req, resp, err = v2.NewSetupCmd1()
		return
	})
	// CMD信令 2
	internal.Command(cp, func() (req v2.Marshaler, resp v2.Unmarshaler, err error) {
		req, resp, err = v2.NewSetupCmd2()
		return
	})
	// CMD信令 3
	internal.Command(cp, func() (req v2.Marshaler, resp v2.Unmarshaler, err error) {
		req, resp, err = v2.NewSetupCmd3()
		return
	})
	// 查询股票数量
	internal.Command(cp, func() (req v2.Marshaler, resp v2.Unmarshaler, err error) {
		req, resp, err = v2.NewSecurityCount(v2.MarketShangHai)
		return
	})
	// 查询股票列表
	internal.Command(cp, func() (req v2.Marshaler, resp v2.Unmarshaler, err error) {
		req, resp, err = v2.NewGetSecurityList(v2.MarketShangHai, 255)
		return
	})
	// 查询个股基本面
	resp := internal.Command(cp, func() (req v2.Marshaler, resp v2.Unmarshaler, err error) {
		req, resp, err = v2.NewFinanceInfo(v2.MarketShangHai, "600600")
		return
	})
	fmt.Println(resp)
}
