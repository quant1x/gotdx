package quotes

import (
	"bytes"
	"encoding/binary"

	"github.com/quant1x/gotdx/internal"
	"github.com/quant1x/gotdx/proto"
)

const (
	TDX_SECURITY_LIST_MAX = 1000 // 单次最大获取多少条股票数据
)

// SecurityListPackage 股票列表
type SecurityListPackage struct {
	reqHeader  *StdRequestHeader
	respHeader *StdResponseHeader
	request    *SecurityListRequest
	reply      *SecurityListReply

	contentHex string
}

type SecurityListRequest struct {
	Market uint16
	Start  uint16
}

type SecurityListReply struct {
	Count uint16
	List  []Security
}

type Security struct {
	Code      string
	VolUnit   uint16
	Reversed1 [4]byte `dataframe:"-"`
	//R1           uint32
	//P1           float64
	DecimalPoint int8
	Name         string
	PreClose     float64
	Reversed2    [4]byte `dataframe:"-"`
	//R2           uint32
	//P2           float64
}

func NewSecurityListPackage() *SecurityListPackage {
	obj := new(SecurityListPackage)
	obj.reqHeader = new(StdRequestHeader)
	obj.respHeader = new(StdResponseHeader)
	obj.request = new(SecurityListRequest)
	obj.reply = new(SecurityListReply)

	obj.reqHeader.ZipFlag = proto.FlagNotZipped
	obj.reqHeader.SeqID = internal.SequenceId()
	obj.reqHeader.PacketType = 0x01
	obj.reqHeader.Method = proto.STD_MSG_SECURITY_LIST
	return obj
}

func (obj *SecurityListPackage) SetParams(req *SecurityListRequest) {
	obj.request = req
}

func (obj *SecurityListPackage) Serialize() ([]byte, error) {
	obj.reqHeader.PkgLen1 = 2 + 4
	obj.reqHeader.PkgLen2 = 2 + 4

	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, obj.reqHeader)
	err = binary.Write(buf, binary.LittleEndian, obj.request)

	//b, err := hex.DecodeString(obj.contentHex)
	//buf.Write(b)

	//err = binary.Write(buf, binary.LittleEndian, uint16(len(obj.stocks)))

	return buf.Bytes(), err
}

func (obj *SecurityListPackage) UnSerialize(header interface{}, data []byte) error {
	obj.respHeader = header.(*StdResponseHeader)

	pos := 0
	err := binary.Read(bytes.NewBuffer(data[pos:pos+2]), binary.LittleEndian, &obj.reply.Count)
	pos += 2
	for index := uint16(0); index < obj.reply.Count; index++ {
		ele := Security{}
		var code [6]byte
		_ = binary.Read(bytes.NewBuffer(data[pos:pos+6]), binary.LittleEndian, &code)
		pos += 6
		ele.Code = string(code[:])

		_ = binary.Read(bytes.NewBuffer(data[pos:pos+2]), binary.LittleEndian, &ele.VolUnit)
		pos += 2

		var name [8]byte
		_ = binary.Read(bytes.NewBuffer(data[pos:pos+8]), binary.LittleEndian, &name)
		ele.Name = internal.Utf8ToGbk(name[:])
		pos += 8

		_ = binary.Read(bytes.NewBuffer(data[pos:pos+4]), binary.LittleEndian, &ele.Reversed1)
		//_ = binary.Read(bytes.NewBuffer(data[pos:pos+4]), binary.LittleEndian, &ele.R1)
		//ele.P1 = getVolume(int(ele.R1))
		pos += 4

		_ = binary.Read(bytes.NewBuffer(data[pos:pos+1]), binary.LittleEndian, &ele.DecimalPoint)
		pos += 1
		var rawPreClose uint32
		_ = binary.Read(bytes.NewBuffer(data[pos:pos+4]), binary.LittleEndian, &rawPreClose)
		ele.PreClose = internal.IntToFloat64(int(rawPreClose))
		pos += 4

		_ = binary.Read(bytes.NewBuffer(data[pos:pos+4]), binary.LittleEndian, &ele.Reversed2)
		//_ = binary.Read(bytes.NewBuffer(data[pos:pos+4]), binary.LittleEndian, &ele.R2)
		//ele.P2 = getVolume(int(ele.R2))
		pos += 4

		obj.reply.List = append(obj.reply.List, ele)
	}
	return err
}

func (obj *SecurityListPackage) Reply() interface{} {
	return obj.reply
}
