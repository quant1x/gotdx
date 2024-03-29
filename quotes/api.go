package quotes

import (
	"errors"
	"gitee.com/quant1x/gotdx/proto"
	"github.com/mymmsc/gox/api"
	"github.com/mymmsc/gox/logger"
	"io"
	"strconv"
	"strings"
)

// Server 主机信息
type Server struct {
	Name      string `json:"name"`
	Host      string `json:"host"`
	Port      int    `json:"port"`
	CrossTime int64  `json:"crossTime"`
}

func (s Server) Addr() string {
	return strings.Join([]string{s.Host, strconv.Itoa(s.Port)}, ":")
}

type StdApi struct {
	connPool *ConnPool
}

// NewStdApi 创建一个标准接口
func NewStdApi() (*StdApi, error) {
	server := GetFastHost(TDX_HOST_HQ)
	return NewStdApiWithServers(server)
}

// NewStdApiWithServers 通过服务器组创建一个标准接口
func NewStdApiWithServers(srvs []Server) (*StdApi, error) {
	size := 1
	opt := Opt{
		Servers: srvs,
	}
	stdApi := StdApi{}
	_factory := func() (interface{}, error) {
		client := NewClient(&opt)
		err := client.Connect()
		if err != nil {
			return nil, err
		}
		_ = stdApi.tdx_hello1(client)
		_ = stdApi.tdx_hello2(client)
		return client, err
	}
	_close := func(v interface{}) error {
		client := v.(*TcpClient)
		return client.Close()
	}
	_ping := func(v interface{}) error {
		client := v.(*TcpClient)
		return stdApi.tdx_ping(client)
	}
	cp, err := NewConnPool(opt, size, _factory, _close, _ping)
	if err != nil {
		return nil, err
	}
	stdApi.connPool = cp
	return &stdApi, nil
}

// Close 关闭
func (this *StdApi) Close() {
	this.connPool.Close()
}

// 通过池关闭连接
func (this *StdApi) poolClose(cli *TcpClient) error {
	return this.connPool.CloseConn(cli)
}

func (this *StdApi) tdx_hello1(client *TcpClient) error {
	opt := client.GetOpt()
	conn := client.GetConn()
	// 创建一个hello1消息
	hello1 := NewHello1()
	err := process(conn, hello1, opt)
	if err != nil {
		_ = this.poolClose(client)
		return err
	}
	reply := hello1.Reply().(*Hello1Reply)
	logger.Warnf(reply.Info)
	return nil
}

func (this *StdApi) tdx_hello2(client *TcpClient) error {
	opt := client.GetOpt()
	conn := client.GetConn()
	// 创建一个hello1消息
	hello2 := NewHello2()
	err := process(conn, hello2, opt)
	if err != nil {
		_ = this.poolClose(client)
		return err
	}
	reply := hello2.Reply().(*Hello2Reply)
	logger.Warnf(reply.Info)
	return nil
}

func (this *StdApi) tdx_ping(client *TcpClient) error {
	opt := client.GetOpt()
	conn := client.GetConn()
	msg := NewSecurityCountPackage()
	msg.SetParams(&SecurityCountRequest{
		Market: uint16(1),
	})
	err := process(conn, msg, opt)
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
	cli := _conn.(*TcpClient)
	opt := cli.GetOpt()
	conn := cli.GetConn()
	err := process(conn, msg, opt)
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
func (this *StdApi) GetFinanceInfo(market proto.Market, code string, num uint16) (*FinanceInfo, error) {
	msg := NewFinanceInfoPackage()
	_code := [6]byte{}
	_market := uint8(market)
	copy(_code[:], code)
	msg.SetParams(&FinanceInfoRequest{
		Market: _market,
		Code:   _code,
		Count:  num,
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
func (this *StdApi) GetIndexBars(market proto.Market, code string, category uint16, start uint16, count uint16) (*SecurityBarsReply, error) {
	msg := NewIndexBarsPackage()
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
	return reply.(*SecurityBarsReply), err
}

// GetSecurityQuotes 获取盘口五档报价
func (this *StdApi) GetSecurityQuotes(markets []proto.Market, codes []string) (*SecurityQuotesReply, error) {
	if len(markets) != len(codes) {
		return nil, errors.New("market code count error")
	}
	obj := NewSecurityQuotesPackage()
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

func (this *StdApi) V2GetSecurityQuotes(markets []proto.Market, codes []string) (*V2SecurityQuotesReply, error) {
	if len(markets) != len(codes) {
		return nil, errors.New("market code count error")
	}
	obj := NewV2SecurityQuotesPackage()
	var params []V2Stock
	for i, market := range markets {
		params = append(params, V2Stock{
			Market: market,
			Code:   codes[i],
		})
	}
	obj.SetParams(&V2SecurityQuotesRequest{StockList: params})
	reply, err := this.command(obj)
	if err != nil {
		return nil, err
	}
	return reply.(*V2SecurityQuotesReply), err
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
func (this *StdApi) GetHistoryTransactionData(market proto.Market, code string, date uint32, start uint16, count uint16) (*TransactionReply, error) {
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
	return reply.(*TransactionReply), err
}

// GetXdxrInfo 获取除权除息信息
func (this *StdApi) GetXdxrInfo(market proto.Market, code string) ([]XdxrInfo, error) {
	obj := NewXdxrInfoPackage()
	_code := [6]byte{}
	_market := market
	copy(_code[:], code)
	obj.SetParams(&XdxrInfoRequest{
		Market: _market,
		Code:   _code,
		//Count:  1,
	})
	reply, err := this.command(obj)
	if err != nil {
		return nil, err
	}
	return reply.([]XdxrInfo), err
}

func (this *StdApi) GetBlockMeta(block_file string) (*BlockMeta, error) {
	obj := NewBlockMetaPackage()
	_code := [40]byte{}
	bf := api.String2Bytes(block_file)
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

func (this *StdApi) v1GetBlockInfo(block_file string) (*BlockInfoResponse, error) {
	obj := NewBlockInfoPackage()
	_code := [100]byte{}
	bf := api.String2Bytes(block_file)
	copy(_code[:], bf)
	obj.SetParams(&BlockInfoRequest{
		BlockFile: _code,
		Start:     0,
		Size:      BLOCK_CHUNKS_SIZE,
	})
	reply, err := this.command(obj)
	if err != nil {
		return nil, err
	}
	return reply.(*BlockInfoResponse), err
}

func (this *StdApi) GetBlockInfo(block_file string) (*BlockInfoResponse, error) {
	var resp BlockInfoResponse
	_code := [100]byte{}
	bf := api.String2Bytes(block_file)
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
