package quotes

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"gitee.com/quant1x/gotdx/proto"
)

type HistoryTransactionPackage struct {
	reqHeader  *StdRequestHeader
	respHeader *StdResponseHeader
	request    *HistoryTransactionRequest
	reply      *HistoryTransactionReply

	contentHex string
}

type HistoryTransactionRequest struct {
	Date   uint32
	Market uint16
	Code   [6]byte
	Start  uint16
	Count  uint16
}

type HistoryTransactionReply struct {
	Count uint16
	List  []HistoryTransaction
}

type HistoryTransaction struct {
	Time      string
	Price     float64
	Vol       int
	Num       int
	BuyOrSell int
}

func NewHistoryTransactionPackage() *HistoryTransactionPackage {
	obj := new(HistoryTransactionPackage)
	obj.reqHeader = new(StdRequestHeader)
	obj.respHeader = new(StdResponseHeader)
	obj.request = new(HistoryTransactionRequest)
	obj.reply = new(HistoryTransactionReply)

	obj.reqHeader.Zip = 0x0c
	obj.reqHeader.SeqID = seqID()
	obj.reqHeader.PacketType = 0x00
	//obj.reqHeader.PkgLen1  =
	//obj.reqHeader.PkgLen2  =
	obj.reqHeader.Method = proto.KMSG_HISTORYTRANSACTIONDATA
	//obj.reqHeader.Method = KMSG_MINUTETIMEDATA
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

	pos := 0
	err := binary.Read(bytes.NewBuffer(data[pos:pos+2]), binary.LittleEndian, &obj.reply.Count)
	// 跳过4个字节
	pos += 6

	lastPrice := 0
	for index := uint16(0); index < obj.reply.Count; index++ {
		ele := HistoryTransaction{}
		h, m := gettime(data, &pos)
		ele.Time = fmt.Sprintf("%02d:%02d", h, m)
		rawPrice := getprice(data, &pos)
		ele.Vol = getprice(data, &pos)
		ele.BuyOrSell = getprice(data, &pos)
		getprice(data, &pos)

		lastPrice = lastPrice + rawPrice
		ele.Price = float64(lastPrice) / baseUnit(string(obj.request.Code[:]))
		obj.reply.List = append(obj.reply.List, ele)
	}
	return err
}

func (obj *HistoryTransactionPackage) Reply() interface{} {
	return obj.reply
}
