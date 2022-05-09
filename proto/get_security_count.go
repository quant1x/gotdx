package proto

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
)

// GetSecurityCount
type GetSecurityCount struct {
	ReqHeader
	content string
	Reply   *GetSecurityCountReply

	market uint16
}

type GetSecurityCountReply struct {
	Count uint16
}

func NewGetSecurityCount() *GetSecurityCount {
	obj := &GetSecurityCount{}
	obj.Zip = 0x0c
	obj.SeqID = seqID()
	obj.PacketType = 0x01
	obj.Method = KMSG_SECURITYCOUNT
	obj.content = "75c73301" // 未解
	return obj
}
func (obj *GetSecurityCount) SetParams(market uint16) {
	obj.market = market
}

func (obj *GetSecurityCount) Serialize() ([]byte, error) {
	obj.PkgLen1 = 2 + uint16(len(obj.content)) + 2
	obj.PkgLen2 = 2 + uint16(len(obj.content)) + 2

	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, obj.ReqHeader)
	err = binary.Write(buf, binary.LittleEndian, obj.market)
	b, err := hex.DecodeString(obj.content)
	buf.Write(b)
	return buf.Bytes(), err
}

func (obj *GetSecurityCount) UnSerialize(header interface{}, data []byte) error {
	obj.Reply = new(GetSecurityCountReply)
	obj.Reply.Count = binary.LittleEndian.Uint16(data[:2])
	return nil
}
