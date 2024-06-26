package quotes

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"gitee.com/quant1x/exchange"
	"gitee.com/quant1x/gotdx/internal"
	"gitee.com/quant1x/gotdx/proto"
	"gitee.com/quant1x/gox/api"
)

type HistoryTransactionPackage struct {
	reqHeader  *StdRequestHeader
	respHeader *StdResponseHeader
	request    *HistoryTransactionRequest
	reply      *TransactionReply

	contentHex string
}

type HistoryTransactionRequest struct {
	Date   uint32
	Market uint16
	Code   [6]byte
	Start  uint16
	Count  uint16
}

//type HistoryTransactionReply struct {
//	Count uint16
//	List  []HistoryTransaction
//}
//
//type HistoryTransaction struct {
//	Time      string
//	Price     float64
//	Vol       int
//	Num       int
//	BuyOrSell int
//}

func NewHistoryTransactionPackage() *HistoryTransactionPackage {
	obj := new(HistoryTransactionPackage)
	obj.reqHeader = new(StdRequestHeader)
	obj.respHeader = new(StdResponseHeader)
	obj.request = new(HistoryTransactionRequest)
	obj.reply = new(TransactionReply)

	obj.reqHeader.ZipFlag = proto.FlagNotZipped
	obj.reqHeader.SeqID = internal.SequenceId()
	obj.reqHeader.PacketType = 0x00
	//obj.reqHeader.PkgLen1  =
	//obj.reqHeader.PkgLen2  =
	obj.reqHeader.Method = proto.STD_MSG_HISTORY_TRANSACTION_DATA
	obj.contentHex = ""
	return obj
}

// SetParams 设置参数
func (obj *HistoryTransactionPackage) SetParams(req *HistoryTransactionRequest) {
	obj.request = req
}

func (obj *HistoryTransactionPackage) Serialize() ([]byte, error) {
	obj.reqHeader.PkgLen1 = 0x12
	obj.reqHeader.PkgLen2 = 0x12

	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, obj.reqHeader)
	err = binary.Write(buf, binary.LittleEndian, obj.request)
	b, err := hex.DecodeString(obj.contentHex)
	buf.Write(b)

	//b, err := hex.DecodeString(obj.contentHex)
	//buf.Write(b)

	//err = binary.Write(buf, binary.LittleEndian, uint16(len(obj.stocks)))

	return buf.Bytes(), err
}

func (obj *HistoryTransactionPackage) UnSerialize(header interface{}, data []byte) error {
	obj.respHeader = header.(*StdResponseHeader)

	market := exchange.MarketType(obj.request.Market)
	code := api.Bytes2String(obj.request.Code[:])
	isIndex := exchange.AssertIndexByMarketAndCode(market, code)

	pos := 0
	err := binary.Read(bytes.NewBuffer(data[pos:pos+2]), binary.LittleEndian, &obj.reply.Count)
	pos += 2
	//var a uint16
	//var b uint32
	//binary.Read(bytes.NewBuffer(data[pos:pos+2]), binary.LittleEndian, &a)
	//binary.Read(bytes.NewBuffer(data[pos:pos+4]), binary.LittleEndian, &b)
	// 跳过4个字节
	pos += 4
	baseUnit := internal.BaseUnit(market, code)
	lastPrice := 0
	for index := uint16(0); index < obj.reply.Count; index++ {
		ele := TickTransaction{}
		h, m := internal.GetTime(data, &pos)
		ele.Time = fmt.Sprintf("%02d:%02d", h, m)
		rawPrice := internal.DecodeVarint(data, &pos)
		ele.Vol = internal.DecodeVarint(data, &pos)
		ele.BuyOrSell = internal.DecodeVarint(data, &pos)
		internal.DecodeVarint(data, &pos)

		lastPrice = lastPrice + rawPrice
		ele.Price = float64(lastPrice) / baseUnit

		if isIndex {
			amount := ele.Vol * 100
			ele.Amount = float64(amount)
			ele.Vol = int(float64(amount) / ele.Price)
		} else {
			ele.Vol *= 100
			ele.Amount = float64(ele.Vol) * ele.Price
		}

		obj.reply.List = append(obj.reply.List, ele)
	}
	return err
}

func (obj *HistoryTransactionPackage) Reply() interface{} {
	return obj.reply
}
