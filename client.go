package gotdx

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"errors"
	"gitee.com/quant1x/gotdx/proto"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
)

type TcpClient struct {
	conn     net.Conn
	opt      *Opt
	complete chan bool
	sending  chan bool
}

type Opt struct {
	Host          string
	Port          int
	MaxRetryTimes int
}

func NewClient(opt *Opt) *TcpClient {
	client := &TcpClient{}
	if opt.MaxRetryTimes <= 0 {
		opt.MaxRetryTimes = DefaultRetryTimes
	}

	client.opt = opt
	client.sending = make(chan bool, 1)
	client.complete = make(chan bool, 1)

	return client
}

func (client *TcpClient) connect() error {
	addr := strings.Join([]string{client.opt.Host, strconv.Itoa(client.opt.Port)}, ":")
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	client.conn = conn
	return err
}

func (client *TcpClient) do(msg proto.Msg) error {
	sendData, err := msg.Serialize()
	if err != nil {
		return err
	}

	retryTimes := 0

	for {
		n, err := client.conn.Write(sendData)
		if n < len(sendData) {
			retryTimes++
			if retryTimes <= client.opt.MaxRetryTimes {
				log.Printf("第%d次重试\n", retryTimes)
			} else {
				return err
			}
		} else {
			if err != nil {
				return err
			}
			break
		}
	}

	headerBytes := make([]byte, proto.MessageHeaderBytes)
	_, err = io.ReadFull(client.conn, headerBytes)
	if err != nil {
		return err
	}

	headerBuf := bytes.NewReader(headerBytes)
	var header proto.ResponseHeader
	if err := binary.Read(headerBuf, binary.LittleEndian, &header); err != nil {
		return err
	}

	if header.ZipSize > proto.MessageMaxBytes {
		log.Printf("msgData has bytes(%d) beyond max %d\n", header.ZipSize, proto.MessageMaxBytes)
		return ErrBadData
	}

	msgData := make([]byte, header.ZipSize)
	_, err = io.ReadFull(client.conn, msgData)
	if err != nil {
		return err
	}

	var out bytes.Buffer
	if header.ZipSize != header.UnZipSize {
		b := bytes.NewReader(msgData)
		r, _ := zlib.NewReader(b)
		io.Copy(&out, r)
		err = msg.UnSerialize(&header, out.Bytes())
	} else {
		err = msg.UnSerialize(&header, msgData)
	}

	return err
}

// Connect 连接券商行情服务器
func (client *TcpClient) Connect() (*proto.Hello1Reply, error) {
	err := client.connect()
	if err != nil {
		return nil, err
	}
	obj := proto.NewHello1()
	err = client.do(obj)
	if err != nil {
		return nil, err
	}
	return obj.Reply(), err
}

// Disconnect 断开服务器
func (client *TcpClient) Disconnect() error {
	return client.conn.Close()
}

// GetSecurityCount 获取指定市场内的证券数目
func (client *TcpClient) GetSecurityCount(market uint16) (*proto.SecurityCountReply, error) {
	obj := proto.NewSecurityCounts()
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
func (client *TcpClient) GetSecurityQuotes(markets []uint8, codes []string) (*proto.SecurityQuotesReply, error) {
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
func (client *TcpClient) GetSecurityList(market uint8, start uint16) (*proto.SecurityListReply, error) {
	obj := proto.GetSecurityList()
	_market := uint16(market)
	obj.SetParams(&proto.SecurityListRequest{Market: _market, Start: start})
	err := client.do(obj)
	if err != nil {
		return nil, err
	}
	return obj.Reply(), err
}

// GetSecurityBars 获取股票K线
func (client *TcpClient) GetSecurityBars(category uint16, market uint8, code string, start uint16, count uint16) (*proto.SecurityBarsReply, error) {
	obj := proto.NewSecurityBars()
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
func (client *TcpClient) GetIndexBars(category uint16, market uint8, code string, start uint16, count uint16) (*proto.IndexBarsReply, error) {
	obj := proto.NewIndexBars()
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
func (client *TcpClient) GetMinuteTimeData(market uint8, code string) (*proto.MinuteTimeReply, error) {
	obj := proto.NewMinuteTimeData()
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
func (client *TcpClient) GetHistoryMinuteTimeData(date uint32, market uint8, code string) (*proto.HistoryMinuteTimeReply, error) {
	obj := proto.NewHistoryMinuteTimeData()
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
func (client *TcpClient) GetTransactionData(market uint8, code string, start uint16, count uint16) (*proto.TransactionReply, error) {
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
func (client *TcpClient) GetHistoryTransactionData(date uint32, market uint8, code string, start uint16, count uint16) (*proto.HistoryTransactionReply, error) {
	obj := proto.NewHistoryTransactionData()
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
