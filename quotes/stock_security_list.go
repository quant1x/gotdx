package quotes

import (
	"bytes"
	"encoding/binary"
	"gitee.com/quant1x/gotdx/proto"
)

// 股票列表
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
	Code         string
	VolUnit      uint16
	DecimalPoint int8
	Name         string
	PreClose     float64
}

func NewSecurityListPackage() *SecurityListPackage {
	obj := new(SecurityListPackage)
	obj.reqHeader = new(StdRequestHeader)
	obj.respHeader = new(StdResponseHeader)
	obj.request = new(SecurityListRequest)
	obj.reply = new(SecurityListReply)

	obj.reqHeader.Zip = 0x0c
	obj.reqHeader.SeqID = seqID()
	obj.reqHeader.PacketType = 0x01
	obj.reqHeader.Method = proto.KMSG_SECURITYLIST
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

	//fmt.Println(hex.EncodeToString(data))
	pos := 0
	err := binary.Read(bytes.NewBuffer(data[pos:pos+2]), binary.LittleEndian, &obj.reply.Count)
	pos += 2
	for index := uint16(0); index < obj.reply.Count; index++ {
		ele := Security{}
		var code [6]byte
		binary.Read(bytes.NewBuffer(data[pos:pos+6]), binary.LittleEndian, &code)
		pos += 6
		ele.Code = string(code[:])

		binary.Read(bytes.NewBuffer(data[pos:pos+2]), binary.LittleEndian, &ele.VolUnit)
		pos += 2

		var name [8]byte
		binary.Read(bytes.NewBuffer(data[pos:pos+8]), binary.LittleEndian, &name)
		pos += 8

		ele.Code = Utf8ToGbk(name[:])

		pos += 4
		binary.Read(bytes.NewBuffer(data[pos:pos+1]), binary.LittleEndian, &ele.DecimalPoint)
		pos += 1
		var precloseraw uint32
		binary.Read(bytes.NewBuffer(data[pos:pos+4]), binary.LittleEndian, &precloseraw)
		pos += 4

		ele.PreClose = getvolume(int(precloseraw))
		pos += 4

		obj.reply.List = append(obj.reply.List, ele)
	}
	return err
}

func (obj *SecurityListPackage) Reply() interface{} {
	return obj.reply
}
