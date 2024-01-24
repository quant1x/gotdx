package quotes

import (
	"errors"
	"gitee.com/quant1x/exchange"
	"gitee.com/quant1x/gox/api"
	"gitee.com/quant1x/gox/cache"
	"gitee.com/quant1x/gox/logger"
	"gitee.com/quant1x/gox/num"
	"io"
)

func (this *StdApi) tdxHello1(client *TcpClient) error {
	// 创建一个hello1消息
	hello1 := NewHello1()
	err := client.Command(hello1)
	if err != nil {
		//_ = this.poolClose(client)
		return err
	}
	reply := hello1.Reply().(*Hello1Reply)
	logger.Warnf(reply.Info)
	return nil
}

func (this *StdApi) tdxHello2(client *TcpClient) error {
	// 创建一个hello1消息
	hello2 := NewHello2()
	err := client.Command(hello2)
	if err != nil {
		//_ = this.poolClose(client)
		return err
	}
	reply := hello2.Reply().(*Hello2Reply)
	logger.Warnf(reply.Info)
	return nil
}

func (this *StdApi) v2Tdx_ping(client *TcpClient) error {
	message := NewHeartBeat()
	err := client.Command(message)
	if err != nil {
		//_ = this.poolClose(client)
		return err
	}
	reply := message.Reply().(*HeartBeatReply)
	if reply == nil {
		return io.EOF
	}
	return nil
}

func (this *StdApi) tdxPing(client *TcpClient) error {
	msg := NewSecurityCountPackage()
	msg.SetParams(&SecurityCountRequest{
		Market: uint16(1),
	})
	err := client.Command(msg)
	if err != nil {
		_ = this.poolClose(client)
		return err
	}
	reply := msg.Reply().(*SecurityCountReply)
	if reply.Count == 0 {
		return io.EOF
	}
	return nil
}

func (this *StdApi) command(msg Message) (interface{}, error) {
	// 2.1 获取TCP连接
	_conn := this.connPool.GetConn()
	if _conn == nil {
		return nil, io.ErrClosedPipe
	}
	cli := _conn.(*TcpClient)
	err := cli.Command(msg)
	if err != nil {
		_ = this.poolClose(cli)
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

func (this *StdApi) HeartBeat() (*HeartBeatReply, error) {
	heartBeat := NewHeartBeat()
	reply, err := this.command(heartBeat)
	if err != nil {
		return nil, err
	}
	return reply.(*HeartBeatReply), nil
}

// GetFinanceInfo 基本面
func (this *StdApi) GetFinanceInfo(code string) (*FinanceInfo, error) {
	msg := NewFinanceInfoPackage()
	mId, _, symbol := exchange.DetectMarket(code)
	_code := [6]byte{}
	_market := mId
	copy(_code[:], symbol)
	msg.SetParams(&FinanceInfoRequest{
		Market: _market,
		Code:   _code,
		Count:  1,
	})
	reply, err := this.command(msg)
	if err != nil {
		return nil, err
	}

	return reply.(*FinanceInfo), nil
}

// GetKLine K线
func (this *StdApi) GetKLine(code string, category uint16, start uint16, count uint16) (*SecurityBarsReply, error) {
	msg := NewSecurityBarsPackage()
	mId, _, symbol := exchange.DetectMarket(code)
	_code := [6]byte{}
	_market := uint16(mId)
	copy(_code[:], symbol)
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

// GetIndexBars 指数K线
func (this *StdApi) GetIndexBars(code string, category uint16, start uint16, count uint16) (*SecurityBarsReply, error) {
	msg := NewIndexBarsPackage()
	mId, _, symbol := exchange.DetectMarket(code)
	_code := [6]byte{}
	_market := uint16(mId)
	copy(_code[:], symbol)
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
	return reply.(*SecurityBarsReply), err
}

// GetSecurityCount 获取指定市场内的证券数目
func (this *StdApi) GetSecurityCount(market exchange.MarketType) (*SecurityCountReply, error) {
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

// GetSecurityList 股票列表
func (this *StdApi) GetSecurityList(market exchange.MarketType, start uint16) (*SecurityListReply, error) {
	msg := NewSecurityListPackage()
	_market := uint16(market)
	msg.SetParams(&SecurityListRequest{Market: _market, Start: start})
	reply, err := this.command(msg)
	if err != nil {
		return nil, err
	}

	return reply.(*SecurityListReply), nil
}

// GetSecurityQuotes 获取盘口五档报价
//
// Deprecated: 废弃, 推荐 GetSnapshot.
func (this *StdApi) GetSecurityQuotes(markets []exchange.MarketType, symbols []string) (*SecurityQuotesReply, error) {
	if len(markets) != len(symbols) {
		return nil, errors.New("market code count error")
	}
	obj := NewSecurityQuotesPackage()
	var params []Stock
	for i, market := range markets {
		params = append(params, Stock{
			Market: market,
			Code:   symbols[i],
		})
	}
	obj.SetParams(&SecurityQuotesRequest{StockList: params})
	reply, err := this.command(obj)
	if err != nil {
		return nil, err
	}
	return reply.(*SecurityQuotesReply), err
}

// V2GetSecurityQuotes 测试版本快照
//
// Deprecated: 不推荐
func (this *StdApi) V2GetSecurityQuotes(markets []exchange.MarketType, symbols []string) (*V2SecurityQuotesReply, error) {
	if len(markets) != len(symbols) {
		return nil, errors.New("market code count error")
	}
	obj := NewV2SecurityQuotesPackage()
	var params []V2Stock
	for i, market := range markets {
		params = append(params, V2Stock{
			Market: market,
			Code:   symbols[i],
		})
	}
	obj.SetParams(&V2SecurityQuotesRequest{StockList: params})
	reply, err := this.command(obj)
	if err != nil {
		return nil, err
	}
	return reply.(*V2SecurityQuotesReply), err
}

var (
	poolSnapshot cache.Pool[Snapshot]
)

// GetSnapshot 获取快照数据
func (this *StdApi) GetSnapshot(codes []string) (list []Snapshot, err error) {
	var marketIds []exchange.MarketType
	var symbols []string
	for _, code := range codes {
		id, _, symbol := exchange.DetectMarket(code)
		if len(symbol) == 6 {
			marketIds = append(marketIds, id)
			symbols = append(symbols, symbol)
		}
	}
	if len(symbols) == 0 {
		err = errors.New("code count error")
		return
	}
	if len(symbols) > TDX_SECURITY_QUOTES_MAX {
		err = errors.New("code count error")
		return
	}
	obj := NewSecurityQuotesPackage()
	var params []Stock
	for i, market := range marketIds {
		params = append(params, Stock{
			Market: market,
			Code:   symbols[i],
		})
	}
	obj.SetParams(&SecurityQuotesRequest{StockList: params})
	reply, err := this.command(obj)
	if err != nil {
		return list, err
	}

	upDateInRealTime, status := exchange.CanUpdateInRealtime()
	quoteReply := reply.(*SecurityQuotesReply)
	list = make([]Snapshot, 0, len(quoteReply.List))
	currentTransactionDate := exchange.GetCurrentlyDay()
	for _, v := range quoteReply.List {
		//var snapshot Snapshot
		snapshot := poolSnapshot.Acquire()
		err := api.Copy(snapshot, &v)
		if err == nil {
			snapshot.Date = currentTransactionDate
			snapshot.SecurityCode = exchange.GetSecurityCode(v.Market, v.Code)
			snapshot.Active = v.Active1
			snapshot.ExchangeState = TDX_EXCHANGE_STATE_CLOSING
			if snapshot.State == TDX_SECURITY_TRADE_STATE_DELISTING {
				// 终止上市
				snapshot.ExchangeState = TDX_EXCHANGE_STATE_DELISTING
			}
			if upDateInRealTime {
				// 交易时段
				snapshot.ExchangeState = TDX_EXCHANGE_STATE_NORMAL
			}
			if status == exchange.ExchangeSuspend {
				// 交易暂停
				snapshot.ExchangeState = TDX_EXCHANGE_STATE_PAUSE
			}
			//if snapshot.ExchangeState == TDX_EXCHANGE_STATE_CLOSING {
			//	// 收盘
			//	snapshot.CloseVolume = v.CurVol * 100
			//}
			amount := num.Decimal(snapshot.Amount, 4)
			if amount <= 0.0000 {
				snapshot.Amount = 0.00
			}
			list = append(list, *snapshot)
		}
		poolSnapshot.Release(snapshot)
	}
	return list, nil
}

// GetMinuteTimeData 获取分时图数据
func (this *StdApi) GetMinuteTimeData(code string) (*MinuteTimeReply, error) {
	obj := NewMinuteTimePackage()
	mId, _, symbol := exchange.DetectMarket(code)
	_code := [6]byte{}
	copy(_code[:], symbol)
	obj.SetParams(&MinuteTimeRequest{
		Market: uint16(mId),
		Code:   _code,
	})
	reply, err := this.command(obj)
	if err != nil {
		return nil, err
	}
	return reply.(*MinuteTimeReply), err
}

// GetHistoryMinuteTimeData 获取历史分时图数据
func (this *StdApi) GetHistoryMinuteTimeData(code string, date uint32) (*MinuteTimeReply, error) {
	obj := NewHistoryMinuteTimePackage()
	mId, _, symbol := exchange.DetectMarket(code)
	_code := [6]byte{}
	copy(_code[:], symbol)
	obj.SetParams(&HistoryMinuteTimeRequest{
		Date:   date,
		Market: mId,
		Code:   _code,
	})
	reply, err := this.command(obj)
	if err != nil {
		return nil, err
	}
	return reply.(*MinuteTimeReply), err
}

// GetTransactionData 获取分时成交
func (this *StdApi) GetTransactionData(code string, start uint16, count uint16) (*TransactionReply, error) {
	obj := NewTransactionPackage()
	mId, _, symbol := exchange.DetectMarket(code)
	_code := [6]byte{}
	copy(_code[:], symbol)
	obj.SetParams(&TransactionRequest{
		Market: uint16(mId),
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
func (this *StdApi) GetHistoryTransactionData(code string, date uint32, start uint16, count uint16) (*TransactionReply, error) {
	obj := NewHistoryTransactionPackage()
	mId, _, symbol := exchange.DetectMarket(code)
	_code := [6]byte{}
	copy(_code[:], symbol)
	obj.SetParams(&HistoryTransactionRequest{
		Date:   date,
		Market: uint16(mId),
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

// GetXdxrInfo 获取除权除息信息
func (this *StdApi) GetXdxrInfo(code string) ([]XdxrInfo, error) {
	obj := NewXdxrInfoPackage()
	mId, _, symbol := exchange.DetectMarket(code)
	_code := [6]byte{}
	copy(_code[:], symbol)
	obj.SetParams(&XdxrInfoRequest{
		Market: mId,
		Code:   _code,
	})
	reply, err := this.command(obj)
	if err != nil {
		return nil, err
	}
	return reply.([]XdxrInfo), err
}

func (this *StdApi) GetBlockMeta(blockFile string) (*BlockMeta, error) {
	obj := NewBlockMetaPackage()
	_code := [40]byte{}
	bf := api.String2Bytes(blockFile)
	copy(_code[:], bf)
	obj.SetParams(&BlockMetaRequest{
		BlockFile: _code,
	})
	reply, err := this.command(obj)
	if err != nil {
		return nil, err
	}
	return reply.(*BlockMeta), err
}

func (this *StdApi) GetBlockInfo(blockFile string) (*BlockInfoResponse, error) {
	var resp BlockInfoResponse
	_code := [100]byte{}
	bf := api.String2Bytes(blockFile)
	copy(_code[:], bf)
	start := uint32(0)
	for {
		obj := NewBlockInfoPackage()
		obj.SetParams(&BlockInfoRequest{
			BlockFile: _code,
			Start:     start,
			Size:      BLOCK_CHUNKS_SIZE,
		})
		reply, err := this.command(obj)
		if err != nil {
			return nil, err
		}
		tmp := reply.(*BlockInfoResponse)
		resp.Size += tmp.Size
		resp.Data = append(resp.Data, tmp.Data...)
		if tmp.Size == 0 {
			return nil, io.EOF
		} else if tmp.Size < BLOCK_CHUNKS_SIZE {
			break
		}
		start += tmp.Size
	}
	return &resp, nil
}

func (this *StdApi) GetCompanyInfoCategory(code string) ([]CompanyInfoCategory, error) {
	obj := NewCompanyInfoCategoryPackage()
	mId, _, symbol := exchange.DetectMarket(code)
	_code := [6]byte{}
	_market := uint16(mId)
	copy(_code[:], symbol)
	obj.SetParams(&CompanyInfoCategoryRequest{
		Market: _market,
		Code:   _code,
	})
	reply, err := this.command(obj)
	if err != nil {
		return nil, err
	}
	return reply.([]CompanyInfoCategory), err
}

func (this *StdApi) GetCompanyInfoContent(code string, name string) (*CompanyInfoContent, error) {
	categories, err := this.GetCompanyInfoCategory(code)
	if err != nil {
		return nil, err
	}
	var category *CompanyInfoCategory
	for _, v := range categories {
		if v.Name == name {
			category = &v
			break
		}
	}

	if category == nil {
		return nil, errors.New("not found")
	}
	obj := NewCompanyInfoContentPackage()
	mId, _, symbol := exchange.DetectMarket(code)
	reqest := CompanyInfoContentRequest{
		Market: uint16(mId),
		Offset: category.Offset,
		Length: category.Length,
	}
	copy(reqest.Code[:], symbol)
	copy(reqest.Filename[:], category.Filename)

	obj.SetParams(&reqest)
	reply, err := this.command(obj)
	if err != nil {
		return nil, err
	}
	return reply.(*CompanyInfoContent), err
}
