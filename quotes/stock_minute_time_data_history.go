package quotes

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"gitee.com/quant1x/gotdx/proto"
)

type HistoryMinuteTimePackage struct {
	reqHeader  *StdRequestHeader
	respHeader *StdResponseHeader
	request    *HistoryMinuteTimeRequest
	reply      *HistoryMinuteTimeReply

	contentHex string
}

type HistoryMinuteTimeRequest struct {
	Date   uint32
	Market uint8
	Code   [6]byte
}

type HistoryMinuteTimeReply struct {
	Count uint16
	List  []HistoryMinuteTime
}

type HistoryMinuteTime struct {
	Price float32
	Vol   int
}

func NewHistoryMinuteTimePackage() *HistoryMinuteTimePackage {
	obj := new(HistoryMinuteTimePackage)
	obj.reqHeader = new(StdRequestHeader)
	obj.respHeader = new(StdResponseHeader)
	obj.request = new(HistoryMinuteTimeRequest)
	obj.reply = new(HistoryMinuteTimeReply)

	obj.reqHeader.Zip = 0x0c
	obj.reqHeader.SeqID = seqID()
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

// 结果数据都是\n,\t分隔的中文字符串，比如查询K线数据，返回的结果字符串就形如
// /“时间\t开盘价\t收盘价\t最高价\t最低价\t成交量\t成交额\n
// /20150519\t4.644000\t4.732000\t4.747000\t4.576000\t146667487\t683638848.000000\n
// /20150520\t4.756000\t4.850000\t4.960000\t4.756000\t353161092\t1722953216.000000”
func (obj *HistoryMinuteTimePackage) UnSerialize(header interface{}, data []byte) error {
	obj.respHeader = header.(*StdResponseHeader)

	pos := 0
	err := binary.Read(bytes.NewBuffer(data[pos:pos+2]), binary.LittleEndian, &obj.reply.Count)
	pos += 2
	// 跳过4个字节 功能未解析
	_, _, _, bType := data[pos], data[pos+1], data[pos+2], data[pos+3]
	pos += 4

	lastprice := 0
	for index := uint16(0); index < obj.reply.Count; index++ {
		priceraw := getPrice(data, &pos)
		_ = getPrice(data, &pos)
		vol := getPrice(data, &pos)
		lastprice += priceraw

		var p float32
		if bType > 0x40 {
			p = float32(lastprice) / 100.0
		} else {
			p = float32(lastprice) / 1000.0
		}

		ele := HistoryMinuteTime{Price: p,
			Vol: vol}
		obj.reply.List = append(obj.reply.List, ele)
	}
	return err
}

func (obj *HistoryMinuteTimePackage) Reply() interface{} {
	return obj.reply
}
