package proto

import (
	"bytes"
	"encoding/binary"
	"gotdx/utils"
)

// Hello2 创建握手消息2
type Hello2 struct {
	ReqHeader
	content []byte
	Reply   *Hello2Reply
}
type Hello2Reply struct {
	Info       string
	serverTime string
}

func NewHello2() *Hello2 {
	obj := &Hello2{}
	obj.Reply = new(Hello2Reply)
	obj.Zip = 0x0c
	obj.SeqID = seqID()
	obj.PacketType = 0x01
	obj.Method = KMSG_CMD2
	obj.content = []byte{0xd5, 0xd0, 0xc9, 0xcc, 0xd6, 0xa4, 0xa8, 0xaf, 0x00, 0x00, 0x00, 0x8f, 0xc2, 0x25, 0x40, 0x13, 0x00, 0x00, 0xd5, 0x00, 0xc9, 0xcc, 0xbd, 0xf0, 0xd7, 0xea, 0x00, 0x00, 0x00, 0x02}
	return obj
}

func (obj *Hello2) Serialize() ([]byte, error) {
	obj.PkgLen1 = 2 + uint16(len(obj.content))
	obj.PkgLen2 = 2 + uint16(len(obj.content))

	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, obj.ReqHeader)
	buf.Write(obj.content)
	return buf.Bytes(), err
}

/*
0100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000011f85e34068747470733a2f2f626967352e6e65776f6e652e636f6d2e636e2f7a797968742f7a645f7a737a712e7a6970000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004150503a414c4c0d0a54494d453a303a30312d31353a30352c31353a30362d32333a35390d0a20202020c4facab9d3c3b5c4b0e6b1bebcb4bdabcda3d3c3a3acceaac1cbc4fab5c4d5fdb3a3cab9d3c32cc7ebbea1bfecc9fdd6c1d5d0c9ccd6a4c8af5043b0e6a1a30d0a20202020c8e7b9fbb2bbc4dcd7d4b6afc9fdbcb6a3acc7ebb5bdb9d9cdf868747470733a2f2f7777772e636d736368696e612e636f6d2fcfc2d4d8b0b2d7b0a3acd0bbd0bbc4fab5c4d6a7b3d6a3a100                                                                   年月日              年月日
*/
func (obj *Hello2) UnSerialize(header interface{}, data []byte) error {
	serverInfo := utils.Utf8ToGbk(data[58:])
	//fmt.Println(fmt.Sprintf("服务器:%s;", serverInfo))
	//fmt.Println(hex.EncodeToString(data))
	obj.Reply.Info = serverInfo
	return nil
}
