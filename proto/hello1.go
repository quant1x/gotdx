package proto

import (
	"bytes"
	"encoding/binary"
	"gotdx/utils"
)

// Hello1 创建握手消息1
type Hello1 struct {
	ReqHeader
	content []byte
	Reply   *Hello1Reply
}
type Hello1Reply struct {
	Info       string
	serverTime string
}

func NewHello1() *Hello1 {
	obj := &Hello1{}
	obj.Reply = new(Hello1Reply)
	obj.Zip = 0x0c
	obj.SeqID = seqID()
	obj.PacketType = 0x01
	obj.Method = KMSG_CMD1
	obj.content = []byte{0x01}
	return obj
}

func (obj *Hello1) Serialize() ([]byte, error) {
	obj.PkgLen1 = 2 + uint16(len(obj.content))
	obj.PkgLen2 = 2 + uint16(len(obj.content))

	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, obj.ReqHeader)
	buf.Write(obj.content)
	return buf.Bytes(), err
}

/*
00e60708051 50 f0 00 d3 a02b2020c03840384038403840384033a02b2020c0384038403840384038403 00 5a8a3401 f94a0100 5a8a3401 fd4a0100ff00e 700000101013f
            分  时    秒                                                                      日期
*/
func (obj *Hello1) UnSerialize(header interface{}, data []byte) error {
	serverInfo := utils.Utf8ToGbk(data[68:])
	//fmt.Println(fmt.Sprintf("服务器:%s;", serverInfo))
	//fmt.Println(hex.EncodeToString(data))
	obj.Reply.Info = serverInfo
	return nil
}
