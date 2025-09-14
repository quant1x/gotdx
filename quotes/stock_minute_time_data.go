package quotes

// todo API未有效解析

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"

	"github.com/quant1x/exchange"
	"github.com/quant1x/gotdx/internal"
	"github.com/quant1x/gotdx/proto"
	"github.com/quant1x/x/api"
)

type MinuteTimePackage struct {
	reqHeader  *StdRequestHeader
	respHeader *StdResponseHeader
	request    *MinuteTimeRequest
	reply      *MinuteTimeReply

	contentHex string
}

type MinuteTimeRequest struct {
	Market uint16
	Code   [6]byte
	Date   uint32
}

type MinuteTimeReply struct {
	Count uint16
	List  []MinuteTime
}

type MinuteTime struct {
	Price float32
	Vol   int
}

// NewMinuteTimePackage 获取分时数据
//
// Deprecated: 废弃, ETF的数据不对需要进一步处理, 推荐 NewHistoryMinuteTimePackage
func NewMinuteTimePackage() *MinuteTimePackage {
	obj := new(MinuteTimePackage)
	obj.reqHeader = new(StdRequestHeader)
	obj.respHeader = new(StdResponseHeader)
	obj.request = new(MinuteTimeRequest)
	obj.reply = new(MinuteTimeReply)

	obj.reqHeader.ZipFlag = proto.FlagNotZipped
	obj.reqHeader.SeqID = internal.SequenceId()
	obj.reqHeader.PacketType = 0x00
	//obj.reqHeader.PkgLen1  =
	//obj.reqHeader.PkgLen2  =
	//obj.reqHeader.Method = 0x051d
	obj.reqHeader.Method = proto.STD_MSG_MINUTETIME_DATA
	obj.contentHex = ""
	return obj
}
func (obj *MinuteTimePackage) SetParams(req *MinuteTimeRequest) {
	obj.request = req
}

func (obj *MinuteTimePackage) Serialize() ([]byte, error) {
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

func (obj *MinuteTimePackage) UnSerialize(header interface{}, data []byte) error {
	obj.respHeader = header.(*StdResponseHeader)

	market := exchange.MarketType(obj.request.Market)
	code := api.Bytes2String(obj.request.Code[:])

	pos := 0
	err := binary.Read(bytes.NewBuffer(data[pos:pos+2]), binary.LittleEndian, &obj.reply.Count)
	pos += 2
	// 跳过4个字节
	pos += 6

	pos += 3

	baseUnit := internal.BaseUnit(market, code)
	lastPrice := 0
	//TODO: ETF的数据不对需要进一步处理
	for index := uint16(0); index < obj.reply.Count; index++ {
		rawPrice := internal.DecodeVarint(data, &pos)
		reversed1 := internal.DecodeVarint(data, &pos)
		_ = reversed1
		vol := internal.DecodeVarint(data, &pos)
		lastPrice += rawPrice

		p := float32(lastPrice) / float32(baseUnit)

		ele := MinuteTime{p, vol}
		obj.reply.List = append(obj.reply.List, ele)
	}
	return err
}

func (obj *MinuteTimePackage) Reply() interface{} {
	return obj.reply
}
