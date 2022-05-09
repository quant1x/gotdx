package proto

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"gotdx/utils"
)

// Hello1 创建握手消息1
type Hello1 struct {
	ReqHeader
	contentHex string
	Reply      *Hello1Reply
}
type Hello1Reply struct {
	Info       string
	serverTime string
}

func NewHello1() *Hello1 {
	obj := &Hello1{}
	obj.Zip = 0x0c
	obj.SeqID = seqID()
	obj.PacketType = 0x01
	obj.Method = KMSG_CMD1
	obj.contentHex = "01"
	return obj
}

func (obj *Hello1) Serialize() ([]byte, error) {
	b, err := hex.DecodeString(obj.contentHex)

	obj.PkgLen1 = 2 + uint16(len(b))
	obj.PkgLen2 = 2 + uint16(len(b))

	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.LittleEndian, obj.ReqHeader)

	buf.Write(b)
	return buf.Bytes(), err
}

/*
00e60708051 50 f0 00 d3 a02b2020c03840384038403840384033a02b2020c0384038403840384038403 00 5a8a3401 f94a0100 5a8a3401 fd4a0100ff00e 700000101013f
            分  时    秒                                                                      日期
*/
func (obj *Hello1) UnSerialize(header interface{}, data []byte) error {
	obj.Reply = new(Hello1Reply)
	serverInfo := utils.Utf8ToGbk(data[68:])
	obj.Reply.Info = serverInfo
	return nil
}
