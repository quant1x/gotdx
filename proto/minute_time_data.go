package proto

// todo API未有效解析

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
)

type MinuteTimeData struct {
	reqHeader  *RequestHeader
	respHeader *ResponseHeader
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

func NewMinuteTimeData() *MinuteTimeData {
	obj := new(MinuteTimeData)
	obj.reqHeader = new(RequestHeader)
	obj.respHeader = new(ResponseHeader)
	obj.request = new(MinuteTimeRequest)
	obj.reply = new(MinuteTimeReply)

	obj.reqHeader.Zip = 0x0c
	obj.reqHeader.SeqID = seqID()
	obj.reqHeader.PacketType = 0x00
	//obj.reqHeader.PkgLen1  =
	//obj.reqHeader.PkgLen2  =
	obj.reqHeader.Method = 0x051d
	//obj.reqHeader.Method = KMSG_MINUTETIMEDATA
	obj.contentHex = ""
	return obj
}
func (obj *MinuteTimeData) SetParams(req *MinuteTimeRequest) {
	obj.request = req
}

func (obj *MinuteTimeData) Serialize() ([]byte, error) {
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

// 结果数据都是\n,\t分隔的中文字符串，比如查询K线数据，返回的结果字符串就形如
// /“时间\t开盘价\t收盘价\t最高价\t最低价\t成交量\t成交额\n
// /20150519\t4.644000\t4.732000\t4.747000\t4.576000\t146667487\t683638848.000000\n
// /20150520\t4.756000\t4.850000\t4.960000\t4.756000\t353161092\t1722953216.000000”
func (obj *MinuteTimeData) UnSerialize(header interface{}, data []byte) error {
	obj.respHeader = header.(*ResponseHeader)

	pos := 0
	err := binary.Read(bytes.NewBuffer(data[pos:pos+2]), binary.LittleEndian, &obj.reply.Count)
	// 跳过4个字节
	pos += 6

	lastprice := 0
	for index := uint16(0); index < obj.reply.Count; index++ {
		priceraw := getprice(data, &pos)
		getprice(data, &pos)
		vol := getprice(data, &pos)
		lastprice = lastprice + priceraw
		ele := MinuteTime{float32(lastprice) / 100.0, vol}
		obj.reply.List = append(obj.reply.List, ele)
	}
	return err
}

func (obj *MinuteTimeData) Reply() *MinuteTimeReply {
	return obj.reply
}
