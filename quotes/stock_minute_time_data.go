package quotes

// todo API未有效解析

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"gitee.com/quant1x/gotdx/internal"
	"gitee.com/quant1x/gotdx/proto"
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

func NewMinuteTimePackage() *MinuteTimePackage {
	obj := new(MinuteTimePackage)
	obj.reqHeader = new(StdRequestHeader)
	obj.respHeader = new(StdResponseHeader)
	obj.request = new(MinuteTimeRequest)
	obj.reply = new(MinuteTimeReply)

	obj.reqHeader.ZipFlag = proto.FlagNotZipped
	obj.reqHeader.SeqID = internal.SeqID()
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

	pos := 0
	err := binary.Read(bytes.NewBuffer(data[pos:pos+2]), binary.LittleEndian, &obj.reply.Count)
	// 跳过4个字节
	pos += 6

	lastprice := 0
	for index := uint16(0); index < obj.reply.Count; index++ {
		priceraw := internal.DecodeVarint(data, &pos)
		internal.DecodeVarint(data, &pos)
		vol := internal.DecodeVarint(data, &pos)
		lastprice = lastprice + priceraw
		ele := MinuteTime{float32(lastprice) / 100.0, vol}
		obj.reply.List = append(obj.reply.List, ele)
	}
	return err
}

func (obj *MinuteTimePackage) Reply() interface{} {
	return obj.reply
}
