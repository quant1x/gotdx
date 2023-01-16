package quotes

import (
	"errors"
	"gitee.com/quant1x/gotdx/proto"
)

type StdApi struct {
	connPool *ConnPool
}

func NewStdApi2() *StdApi {
	server := GetFastHost(TDX_HOST_HQ)
	return NewStdApi(*server)
}

func NewStdApi(srv Server) *StdApi {
	size := 1
	opt := Opt{
		Host: srv.Host,
		Port: srv.Port,
	}
	_factory := func(string) (interface{}, error) {
		client := NewClient(&opt)
		err := client.Connect()
		return client, err
	}
	_close := func(v interface{}) error {
		client := v.(*TcpClient)
		return client.Close()
	}
	_ping := func(v interface{}) error {
		return nil
	}
	cp := NewConnPool2(srv.Host, srv.Port, size, _factory, _close, _ping)
	api := StdApi{
		connPool: cp,
	}
	// TODO: 假定IP地址有效, IP地址的有效性依靠bestip模块
	_, _ = api.Hello1()
	_, _ = api.Hello2()
	return &api
}

// Close 关闭
func (this *StdApi) Close() {
	this.connPool.Close()
}

func (this *StdApi) command(msg Message) (interface{}, error) {

	// 2.1 获取TCP连接
	_conn := this.connPool.GetConn()
	cli := _conn.(*TcpClient)
	opt := cli.GetOpt()
	conn := cli.GetConn()
	err := process(conn, msg, opt)
	if err != nil {
		_ = cli.Close()
		return nil, err
	}

	this.connPool.ReturnConn(_conn)
	return msg.Reply(), nil
}

func (this *StdApi) Hello1() (*Hello1Reply, error) {
	// 创建一个hello1消息
	hello1 := NewHello1()
	reply, err := this.command(hello1)
	if err != nil {
		return nil, err
	}
	return reply.(*Hello1Reply), nil
}

func (this *StdApi) Hello2() (*Hello2Reply, error) {
	// 创建一个hello1消息
	hello2 := NewHello2()

	reply, err := this.command(hello2)
	if err != nil {
		return nil, err
	}
	return reply.(*Hello2Reply), nil
}

// GetFinanceInfo 基本面
func (this *StdApi) GetFinanceInfo(market proto.Market, code string) (*FinanceInfo, error) {
	msg := NewFinanceInfoPackage()
	_code := [6]byte{}
	_market := uint8(market)
	copy(_code[:], code)
	msg.SetParams(&FinanceInfoRequest{
		Market: _market,
		Code:   _code,
	})
	reply, err := this.command(msg)
	if err != nil {
		return nil, err
	}

	return reply.(*FinanceInfo), nil
}

// GetKLine K线
func (this *StdApi) GetKLine(market proto.Market, code string, category uint16, start uint16, count uint16) (*SecurityBarsReply, error) {
	msg := NewSecurityBarsPackage()
	_code := [6]byte{}
	_market := uint16(market)
	copy(_code[:], code)
	msg.SetParams(&SecurityBarsRequest{
		Market:   _market,
		Code:     _code,
		Category: category,
		Start:    start,
		Count:    count,
	})
	reply, err := this.command(msg)
	if err != nil {
		return nil, err
	}

	return reply.(*SecurityBarsReply), nil
}

// GetSecurityList 股票列表
func (this *StdApi) GetSecurityList(market proto.Market, start uint16) (*SecurityListReply, error) {
	msg := NewSecurityListPackage()
	_market := uint16(market)
	msg.SetParams(&SecurityListRequest{Market: _market, Start: start})
	reply, err := this.command(msg)
	if err != nil {
		return nil, err
	}

	return reply.(*SecurityListReply), nil
}

// GetIndexBars 指数K线
func (this *StdApi) GetIndexBars(market proto.Market, code string, category uint16, start uint16, count uint16) (*IndexBarsReply, error) {
	msg := NewIndexBarsPackage()
	_code := [6]byte{}
	_market := uint16(market)
	copy(_code[:], code)
	msg.SetParams(&IndexBarsRequest{
		Market:   _market,
		Code:     _code,
		Category: category,
		Start:    start,
		Count:    count,
	})
	reply, err := this.command(msg)
	if err != nil {
		return nil, err
	}
	return reply.(*IndexBarsReply), err
}

// GetSecurityCount 获取指定市场内的证券数目
func (this *StdApi) GetSecurityCount(market proto.Market) (*SecurityCountReply, error) {
	obj := NewSecurityCountPackage()
	obj.SetParams(&SecurityCountRequest{
		Market: uint16(market),
	})
	reply, err := this.command(obj)
	if err != nil {
		return nil, err
	}
	return reply.(*SecurityCountReply), err
}

// GetSecurityQuotes 获取盘口五档报价
func (this *StdApi) GetSecurityQuotes(markets []proto.Market, codes []string) (*SecurityQuotesReply, error) {
	if len(markets) != len(codes) {
		return nil, errors.New("market code count error")
	}
	obj := NewGetSecurityQuotesPackage()
	var params []Stock
	for i, market := range markets {
		params = append(params, Stock{
			Market: market,
			Code:   codes[i],
		})
	}
	obj.SetParams(&SecurityQuotesRequest{StockList: params})
	reply, err := this.command(obj)
	if err != nil {
		return nil, err
	}
	return reply.(*SecurityQuotesReply), err
}

// GetMinuteTimeData 获取分时图数据
func (this *StdApi) GetMinuteTimeData(market proto.Market, code string) (*MinuteTimeReply, error) {
	obj := NewMinuteTimePackage()
	_code := [6]byte{}
	_market := uint16(market)
	copy(_code[:], code)
	obj.SetParams(&MinuteTimeRequest{
		Market: _market,
		Code:   _code,
	})
	reply, err := this.command(obj)
	if err != nil {
		return nil, err
	}
	return reply.(*MinuteTimeReply), err
}

// GetHistoryMinuteTimeData 获取历史分时图数据
func (this *StdApi) GetHistoryMinuteTimeData(market proto.Market, code string, date uint32) (*HistoryMinuteTimeReply, error) {
	obj := NewHistoryMinuteTimePackage()
	_code := [6]byte{}
	copy(_code[:], code)
	obj.SetParams(&HistoryMinuteTimeRequest{
		Date:   date,
		Market: market,
		Code:   _code,
	})
	reply, err := this.command(obj)
	if err != nil {
		return nil, err
	}
	return reply.(*HistoryMinuteTimeReply), err
}

// GetTransactionData 获取分时成交
func (this *StdApi) GetTransactionData(market proto.Market, code string, start uint16, count uint16) (*TransactionReply, error) {
	obj := NewTransactionPackage()
	_code := [6]byte{}
	_market := uint16(market)
	copy(_code[:], code)
	obj.SetParams(&TransactionRequest{
		Market: _market,
		Code:   _code,
		Start:  start,
		Count:  count,
	})
	reply, err := this.command(obj)
	if err != nil {
		return nil, err
	}
	return reply.(*TransactionReply), err
}

// GetHistoryTransactionData 获取历史分时成交
func (this *StdApi) GetHistoryTransactionData(market proto.Market, code string, date uint32, start uint16, count uint16) (*HistoryTransactionReply, error) {
	obj := NewHistoryTransactionPackage()
	_code := [6]byte{}
	_market := uint16(market)
	copy(_code[:], code)
	obj.SetParams(&HistoryTransactionRequest{
		Date:   date,
		Market: _market,
		Code:   _code,
		Start:  start,
		Count:  count,
	})
	reply, err := this.command(obj)
	if err != nil {
		return nil, err
	}
	return reply.(*HistoryTransactionReply), err
}
