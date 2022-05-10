package gotdx

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"gotdx/proto"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
)

type Client struct {
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

func NewClient(opt *Opt) *Client {
	client := &Client{}
	if opt.MaxRetryTimes <= 0 {
		opt.MaxRetryTimes = DefaultRetryTimes
	}

	client.opt = opt
	client.sending = make(chan bool, 1)
	client.complete = make(chan bool, 1)

	return client
}

func (client *Client) connect() error {
	addr := strings.Join([]string{client.opt.Host, strconv.Itoa(client.opt.Port)}, ":")
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	client.conn = conn
	return err
}

func (client *Client) do(msg proto.Msg) error {
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
	var header proto.RespHeader
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
func (client *Client) Connect() (*proto.Hello1Reply, error) {
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
func (client *Client) Disconnect() error {
	return client.conn.Close()
}

// GetSecurityCount 获取指定市场内的证券数目
func (client *Client) GetSecurityCount(market uint16) (*proto.GetSecurityCountReply, error) {
	obj := proto.NewGetSecurityCount()
	obj.SetParams(&proto.GetSecurityCountRequest{
		Market: market,
	})
	err := client.do(obj)
	if err != nil {
		return nil, err
	}
	return obj.Reply(), err
}

// GetSecurityQuotes 获取盘口五档报价
func (client *Client) GetSecurityQuotes(params []proto.Stock) (*proto.GetSecurityQuotesReply, error) {
	obj := proto.NewGetSecurityQuotes()
	obj.SetParams(&proto.GetSecurityQuotesRequest{StockList: params})
	err := client.do(obj)
	if err != nil {
		return nil, err
	}
	return obj.Reply(), err
}

// GetSecurityList 获取市场内指定范围内的所有证券代码
func (client *Client) GetSecurityList(market uint16, start uint16) (*proto.GetSecurityListReply, error) {
	obj := proto.NewGetSecurityList()
	obj.SetParams(&proto.GetSecurityListRequest{Market: market, Start: start})
	err := client.do(obj)
	if err != nil {
		return nil, err
	}
	return obj.Reply(), err
}

// GetSecurityBars 获取股票K线
// e Category, byte Market, char* Zqdm, short Start, short& Count, char* Result, char* ErrInfo
func (client *Client) GetSecurityBars(category uint16, market uint16, code string, start uint16, count uint16) (*proto.GetSecurityBarsReply, error) {
	obj := proto.NewGetSecurityBars()
	_code := [6]byte{}
	copy(_code[:], code)
	obj.SetParams(&proto.GetSecurityBarsRequest{
		Market:   market,
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
func (client *Client) GetIndexBars(category uint16, market uint16, code string, start uint16, count uint16) (*proto.GetIndexBarsReply, error) {
	obj := proto.NewGetIndexBars()
	_code := [6]byte{}
	copy(_code[:], code)
	obj.SetParams(&proto.GetIndexBarsRequest{
		Market:   market,
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
