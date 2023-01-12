package proto

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
)

type SecurityCounts struct {
	reqHeader  *RequestHeader
	respHeader *ResponseHeader
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

func NewGetSecurityCount() *SecurityCounts {
	obj := new(SecurityCounts)
	obj.reqHeader = new(RequestHeader)
	obj.respHeader = new(ResponseHeader)
	obj.request = new(SecurityCountRequest)
	obj.reply = new(SecurityCountReply)

	obj.reqHeader.Zip = 0x0c
	obj.reqHeader.SeqID = seqID()
	obj.reqHeader.PacketType = 0x01
	obj.reqHeader.Method = KMSG_SECURITYCOUNT
	obj.contentHex = "75c73301" // 未解
	return obj
}

func (obj *SecurityCounts) SetParams(req *SecurityCountRequest) {
	obj.request = req
}

func (obj *SecurityCounts) Serialize() ([]byte, error) {
	obj.reqHeader.PkgLen1 = 2 + 4 + 2
	obj.reqHeader.PkgLen2 = 2 + 4 + 2

	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, obj.reqHeader)
	err = binary.Write(buf, binary.LittleEndian, obj.request)
	b, err := hex.DecodeString(obj.contentHex)
	buf.Write(b)
	return buf.Bytes(), err
}

func (obj *SecurityCounts) UnSerialize(header interface{}, data []byte) error {
	obj.respHeader = header.(*ResponseHeader)

	obj.reply.Count = binary.LittleEndian.Uint16(data[:2])
	return nil
}

func (obj *SecurityCounts) Reply() *SecurityCountReply {
	return obj.reply
}
