package quotes

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"gitee.com/quant1x/gotdx/proto"
)

const (
	TDX_TRANSACTION_MAX = 1800 // 单次最多获取多少条分笔成交记录
)

// TransactionPackage 当日分笔成交信息
type TransactionPackage struct {
	reqHeader  *StdRequestHeader
	respHeader *StdResponseHeader
	request    *TransactionRequest
	reply      *TransactionReply

	contentHex string
}

type TransactionRequest struct {
	Market uint16
	Code   [6]byte
	Start  uint16
	Count  uint16
}

type TransactionReply struct {
	Count uint16
	List  []TickTransaction
}

type TickTransaction struct {
	Time      string
	Price     float64
	Vol       int
	Num       int
	BuyOrSell int
}

func NewTransactionPackage() *TransactionPackage {
	obj := new(TransactionPackage)
	obj.reqHeader = new(StdRequestHeader)
	obj.respHeader = new(StdResponseHeader)
	obj.request = new(TransactionRequest)
	obj.reply = new(TransactionReply)

	obj.reqHeader.Zip = 0x0c
	obj.reqHeader.SeqID = seqID()
	obj.reqHeader.PacketType = 0x00
	//obj.reqHeader.PkgLen1  =
	//obj.reqHeader.PkgLen2  =
	obj.reqHeader.Method = proto.STD_MSG_TRANSACTION_DATA
	obj.contentHex = ""
	return obj
}
func (obj *TransactionPackage) SetParams(req *TransactionRequest) {
	obj.request = req
}

func (obj *TransactionPackage) Serialize() ([]byte, error) {
	obj.reqHeader.PkgLen1 = 0x0e
	obj.reqHeader.PkgLen2 = 0x0e

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

func (obj *TransactionPackage) UnSerialize(header interface{}, data []byte) error {
	obj.respHeader = header.(*StdResponseHeader)

	pos := 0
	err := binary.Read(bytes.NewBuffer(data[pos:pos+2]), binary.LittleEndian, &obj.reply.Count)
	pos += 2

	lastprice := 0
	for index := uint16(0); index < obj.reply.Count; index++ {
		ele := TickTransaction{}
		hour, minute := gettime(data, &pos)
		ele.Time = fmt.Sprintf("%02d:%02d", hour, minute)
		priceraw := getprice(data, &pos)
		ele.Vol = getprice(data, &pos)
		ele.Num = getprice(data, &pos)
		ele.BuyOrSell = getprice(data, &pos)
		lastprice += priceraw
		ele.Price = float64(lastprice) / baseUnit(string(obj.request.Code[:]))
		_ = getprice(data, &pos)
		obj.reply.List = append(obj.reply.List, ele)
	}
	return err
}

func (obj *TransactionPackage) Reply() interface{} {
	return obj.reply
}
