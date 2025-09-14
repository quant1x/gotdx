package quotes

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"

	"gitee.com/quant1x/exchange"
	"gitee.com/quant1x/gotdx/internal"
	"gitee.com/quant1x/gotdx/proto"
	"gitee.com/quant1x/gox/api"
)

type HistoryMinuteTimePackage struct {
	reqHeader  *StdRequestHeader
	respHeader *StdResponseHeader
	request    *HistoryMinuteTimeRequest
	reply      *MinuteTimeReply

	contentHex string
}

type HistoryMinuteTimeRequest struct {
	Date   uint32
	Market uint8
	Code   [6]byte
}

//type HistoryMinuteTimeReply struct {
//	Count uint16
//	List  []HistoryMinuteTime
//}
//
//type HistoryMinuteTime struct {
//	Price float32
//	Vol   int
//}

func NewHistoryMinuteTimePackage() *HistoryMinuteTimePackage {
	obj := new(HistoryMinuteTimePackage)
	obj.reqHeader = new(StdRequestHeader)
	obj.respHeader = new(StdResponseHeader)
	obj.request = new(HistoryMinuteTimeRequest)
	obj.reply = new(MinuteTimeReply)

	obj.reqHeader.ZipFlag = proto.FlagNotZipped
	obj.reqHeader.SeqID = internal.SequenceId()
	obj.reqHeader.PacketType = 0x00
	//obj.reqHeader.PkgLen1  =
	//obj.reqHeader.PkgLen2  =
	obj.reqHeader.Method = proto.STD_MSG_HISTORY_MINUTETIME_DATA
	obj.contentHex = ""
	return obj
}

// SetParams 设置参数
func (obj *HistoryMinuteTimePackage) SetParams(req *HistoryMinuteTimeRequest) {
	obj.request = req
}

func (obj *HistoryMinuteTimePackage) Serialize() ([]byte, error) {
	obj.reqHeader.PkgLen1 = 0x0d
	obj.reqHeader.PkgLen2 = 0x0d

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

func (obj *HistoryMinuteTimePackage) UnSerialize(header interface{}, data []byte) error {
	obj.respHeader = header.(*StdResponseHeader)

	market := exchange.MarketType(obj.request.Market)
	code := api.Bytes2String(obj.request.Code[:])
	dataLen := len(data)
	if dataLen < 2 {
		return nil
	}

	pos := 0
	err := binary.Read(bytes.NewBuffer(data[pos:pos+2]), binary.LittleEndian, &obj.reply.Count)
	pos += 2
	if obj.reply.Count == 0 {
		return nil
	}
	// 跳过4个字节 功能未解析
	if dataLen < 6 {
		return nil
	}
	_, _, _, bType := data[pos], data[pos+1], data[pos+2], data[pos+3]
	pos += 4
	baseUnit := internal.BaseUnit(market, code)
	//var baseUnit float32
	//if bType > 0x40 {
	//	baseUnit = 100.0
	//} else {
	//	baseUnit = 1000.0
	//}
	lastPrice := 0
	for index := uint16(0); index < obj.reply.Count; index++ {
		rawPrice := internal.DecodeVarint(data, &pos)
		reversed1 := internal.DecodeVarint(data, &pos)
		_ = reversed1
		vol := internal.DecodeVarint(data, &pos)
		lastPrice += rawPrice

		p := float32(lastPrice) / float32(baseUnit)
		ele := MinuteTime{Price: p, Vol: vol}
		obj.reply.List = append(obj.reply.List, ele)
	}
	_ = bType
	return err
}

func (obj *HistoryMinuteTimePackage) Reply() interface{} {
	return obj.reply
}
