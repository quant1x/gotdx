package quotes

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"gitee.com/quant1x/gotdx/internal"
	"gitee.com/quant1x/gotdx/proto"
)

// SecurityCountPackage 市场股票数量
type SecurityCountPackage struct {
	reqHeader  *StdRequestHeader
	respHeader *StdResponseHeader
	request    *SecurityCountRequest
	reply      *SecurityCountReply
	contentHex string
}

type SecurityCountRequest struct {
	Market uint16
}

type SecurityCountReply struct {
	Count uint16
}

func NewSecurityCountPackage() *SecurityCountPackage {
	obj := new(SecurityCountPackage)
	obj.reqHeader = new(StdRequestHeader)
	obj.respHeader = new(StdResponseHeader)
	obj.request = new(SecurityCountRequest)
	obj.reply = new(SecurityCountReply)

	obj.reqHeader.ZipFlag = proto.FlagNotZipped
	obj.reqHeader.SeqID = internal.SeqID()
	obj.reqHeader.PacketType = 0x01
	obj.reqHeader.Method = proto.STD_MSG_SECURITY_COUNT
	obj.contentHex = "75c73301" // 未解
	return obj
}

func (obj *SecurityCountPackage) SetParams(req *SecurityCountRequest) {
	obj.request = req
}

func (obj *SecurityCountPackage) Serialize() ([]byte, error) {
	obj.reqHeader.PkgLen1 = 2 + 4 + 2
	obj.reqHeader.PkgLen2 = 2 + 4 + 2

	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, obj.reqHeader)
	err = binary.Write(buf, binary.LittleEndian, obj.request)
	b, err := hex.DecodeString(obj.contentHex)
	buf.Write(b)
	return buf.Bytes(), err
}

func (obj *SecurityCountPackage) UnSerialize(header interface{}, data []byte) error {
	obj.respHeader = header.(*StdResponseHeader)

	obj.reply.Count = binary.LittleEndian.Uint16(data[:2])
	return nil
}

func (obj *SecurityCountPackage) Reply() interface{} {
	return obj.reply
}
