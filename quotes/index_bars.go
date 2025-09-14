package quotes

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"

	"github.com/quant1x/gotdx/internal"
	"github.com/quant1x/gotdx/proto"
)

// IndexBarsPackage 指数K线
type IndexBarsPackage struct {
	reqHeader  *StdRequestHeader
	respHeader *StdResponseHeader
	//request    *IndexBarsRequest
	//reply      *IndexBarsReply
	request *SecurityBarsRequest
	reply   *SecurityBarsReply

	contentHex string
}

//type IndexBarsRequest struct {
//	MarketType   uint16
//	Code     [6]byte
//	Category uint16 // 种类 5分钟  10分钟
//	I        uint16 // 未知 填充
//	Start    uint16
//	Count    uint16
//}
//
//type IndexBarsReply struct {
//	Count uint16
//	List  []IndexBar
//}
//
//type IndexBar struct {
//	Open      float64
//	Close     float64
//	High      float64
//	Low       float64
//	Vol       float64
//	Amount    float64
//	Year      int
//	Month     int
//	Day       int
//	Hour      int
//	Minute    int
//	DateTime  string
//	UpCount   uint16
//	DownCount uint16
//}

func NewIndexBarsPackage() *IndexBarsPackage {
	obj := new(IndexBarsPackage)
	obj.reqHeader = new(StdRequestHeader)
	obj.respHeader = new(StdResponseHeader)
	obj.request = new(SecurityBarsRequest)
	obj.reply = new(SecurityBarsReply)

	obj.reqHeader.ZipFlag = proto.FlagNotZipped
	obj.reqHeader.SeqID = internal.SequenceId()
	obj.reqHeader.PacketType = 0x00
	//obj.reqHeader.PkgLen1  =
	//obj.reqHeader.PkgLen2  =
	obj.reqHeader.Method = proto.STD_MSG_INDEXBARS
	obj.contentHex = "00000000000000000000"
	return obj
}
func (obj *IndexBarsPackage) SetParams(req *SecurityBarsRequest) {
	obj.request = req
	obj.request.I = 1
}

func (obj *IndexBarsPackage) Serialize() ([]byte, error) {
	obj.reqHeader.PkgLen1 = 0x1c
	obj.reqHeader.PkgLen2 = 0x1c

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

func (obj *IndexBarsPackage) UnSerialize(header interface{}, data []byte) error {
	obj.respHeader = header.(*StdResponseHeader)

	pos := 0
	err := binary.Read(bytes.NewBuffer(data[pos:pos+2]), binary.LittleEndian, &obj.reply.Count)
	pos += 2

	pre_diff_base := 0
	for index := uint16(0); index < obj.reply.Count; index++ {
		ele := SecurityBar{}

		ele.Year, ele.Month, ele.Day, ele.Hour, ele.Minute = internal.GetDatetime(int(obj.request.Category), data, &pos)

		//if index == 0 {
		//	ele.Year, ele.Month, ele.Day, ele.Hour, ele.Minute = getDatetime(int(obj.request.Category), data, &pos)
		//} else {
		//	ele.Year, ele.Month, ele.Day, ele.Hour, ele.Minute = getDatetimeNow(int(obj.request.Category), lasttime)
		//}
		ele.DateTime = fmt.Sprintf("%d-%02d-%02d %02d:%02d:00.000", ele.Year, ele.Month, ele.Day, ele.Hour, ele.Minute)

		price_open_diff := internal.DecodeVarint(data, &pos)
		price_close_diff := internal.DecodeVarint(data, &pos)

		price_high_diff := internal.DecodeVarint(data, &pos)
		price_low_diff := internal.DecodeVarint(data, &pos)

		var ivol uint32
		_ = binary.Read(bytes.NewBuffer(data[pos:pos+4]), binary.LittleEndian, &ivol)
		ele.Vol = internal.IntToFloat64(int(ivol))
		pos += 4

		var dbvol uint32
		_ = binary.Read(bytes.NewBuffer(data[pos:pos+4]), binary.LittleEndian, &dbvol)
		ele.Amount = internal.IntToFloat64(int(dbvol))
		pos += 4

		_ = binary.Read(bytes.NewBuffer(data[pos:pos+2]), binary.LittleEndian, &ele.UpCount)
		pos += 2
		_ = binary.Read(bytes.NewBuffer(data[pos:pos+2]), binary.LittleEndian, &ele.DownCount)
		pos += 2

		ele.Open = float64(price_open_diff+pre_diff_base) / 1000.0
		price_open_diff += pre_diff_base

		ele.Close = float64(price_open_diff+price_close_diff) / 1000.0
		ele.High = float64(price_open_diff+price_high_diff) / 1000.0
		ele.Low = float64(price_open_diff+price_low_diff) / 1000.0

		pre_diff_base = price_open_diff + price_close_diff

		obj.reply.List = append(obj.reply.List, ele)
	}
	return err
}

func (obj *IndexBarsPackage) Reply() interface{} {
	return obj.reply
}
