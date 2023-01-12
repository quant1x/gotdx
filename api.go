package gotdx

import (
	"errors"
	"gitee.com/quant1x/gotdx/proto"
)

type TdxApi struct {
	TcpClient
}

// GetSecurityCount 获取指定市场内的证券数目
func (client *TdxApi) GetSecurityCount(market uint16) (*proto.SecurityCountReply, error) {
	obj := proto.NewGetSecurityCount()
	obj.SetParams(&proto.SecurityCountRequest{
		Market: market,
	})
	err := client.do(obj)
	if err != nil {
		return nil, err
	}
	return obj.Reply(), err
}

// GetSecurityQuotes 获取盘口五档报价
func (client *TdxApi) GetSecurityQuotes(markets []uint8, codes []string) (*proto.SecurityQuotesReply, error) {
	if len(markets) != len(codes) {
		return nil, errors.New("market code count error")
	}
	obj := proto.NewGetSecurityQuotes()
	var params []proto.Stock
	for i, market := range markets {
		params = append(params, proto.Stock{
			Market: market,
			Code:   codes[i],
		})
	}
	obj.SetParams(&proto.SecurityQuotesRequest{StockList: params})
	err := client.do(obj)
	if err != nil {
		return nil, err
	}
	return obj.Reply(), err
}

// GetSecurityList 获取市场内指定范围内的所有证券代码
func (client *TdxApi) GetSecurityList(market uint8, start uint16) (*proto.SecurityListReply, error) {
	obj := proto.NewGetSecurityList()
	_market := uint16(market)
	obj.SetParams(&proto.SecurityListRequest{Market: _market, Start: start})
	err := client.do(obj)
	if err != nil {
		return nil, err
	}
	return obj.Reply(), err
}

// GetSecurityBars 获取股票K线
func (client *TdxApi) GetSecurityBars(category uint16, market uint8, code string, start uint16, count uint16) (*proto.SecurityBarsReply, error) {
	obj := proto.NewGetSecurityBars()
	_code := [6]byte{}
	_market := uint16(market)
	copy(_code[:], code)
	obj.SetParams(&proto.SecurityBarsRequest{
		Market:   _market,
		Code:     _code,
		Category: category,
		Start:    start,
		Count:    count,
	})
	err := client.do(obj)
	if err != nil {
		return nil, err
	}
	return obj.Reply(), err
}

// GetIndexBars 获取指数K线
func (client *TdxApi) GetIndexBars(category uint16, market uint8, code string, start uint16, count uint16) (*proto.GetIndexBarsReply, error) {
	obj := proto.NewGetIndexBars()
	_code := [6]byte{}
	_market := uint16(market)
	copy(_code[:], code)
	obj.SetParams(&proto.IndexBarsRequest{
		Market:   _market,
		Code:     _code,
		Category: category,
		Start:    start,
		Count:    count,
	})
	err := client.do(obj)
	if err != nil {
		return nil, err
	}
	return obj.Reply(), err
}

// GetMinuteTimeData 获取分时图数据
func (client *TdxApi) GetMinuteTimeData(market uint8, code string) (*proto.MinuteTimeReply, error) {
	obj := proto.NewGetMinuteTimeData()
	_code := [6]byte{}
	_market := uint16(market)
	copy(_code[:], code)
	obj.SetParams(&proto.MinuteTimeRequest{
		Market: _market,
		Code:   _code,
	})
	err := client.do(obj)
	if err != nil {
		return nil, err
	}
	return obj.Reply(), err
}

// GetHistoryMinuteTimeData 获取历史分时图数据
func (client *TdxApi) GetHistoryMinuteTimeData(date uint32, market uint8, code string) (*proto.HistoryMinuteTimeReply, error) {
	obj := proto.NewGetHistoryMinuteTimeData()
	_code := [6]byte{}
	copy(_code[:], code)
	obj.SetParams(&proto.HistoryMinuteTimeRequest{
		Date:   date,
		Market: market,
		Code:   _code,
	})
	err := client.do(obj)
	if err != nil {
		return nil, err
	}
	return obj.Reply(), err
}

// GetTransactionData 获取分时成交
func (client *TdxApi) GetTransactionData(market uint8, code string, start uint16, count uint16) (*proto.TransactionReply, error) {
	obj := proto.NewGetTransactionData()
	_code := [6]byte{}
	_market := uint16(market)
	copy(_code[:], code)
	obj.SetParams(&proto.TransactionRequest{
		Market: _market,
		Code:   _code,
		Start:  start,
		Count:  count,
	})
	err := client.do(obj)
	if err != nil {
		return nil, err
	}
	return obj.Reply(), err
}

// GetHistoryTransactionData 获取历史分时成交
func (client *TdxApi) GetHistoryTransactionData(date uint32, market uint8, code string, start uint16, count uint16) (*proto.HistoryTransactionReply, error) {
	obj := proto.NewGetHistoryTransactionData()
	_code := [6]byte{}
	_market := uint16(market)
	copy(_code[:], code)
	obj.SetParams(&proto.HistoryTransactionRequest{
		Date:   date,
		Market: _market,
		Code:   _code,
		Start:  start,
		Count:  count,
	})
	err := client.do(obj)
	if err != nil {
		return nil, err
	}
	return obj.Reply(), err
}
