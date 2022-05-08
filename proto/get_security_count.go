package proto

import (
	"bytes"
	"encoding/binary"
)

// SecurityCount
type SecurityCount struct {
	ReqHeader
	content []byte
	Reply   *SecurityCountReply

	market uint16
}

type SecurityCountReply struct {
	Count uint16
}

func NewSecurityCount() *SecurityCount {
	obj := &SecurityCount{}
	obj.Reply = new(SecurityCountReply)
	obj.Zip = 0x0c
	obj.SeqID = seqID()
	obj.PacketType = 0x01
	obj.Method = KMSG_SECURITYCOUNT
	obj.content = []byte{0x75, 0xc7, 0x33, 0x01}
	return obj
}
func (obj *SecurityCount) SetParams(market uint16) {
	obj.market = market
}

func (obj *SecurityCount) Serialize() ([]byte, error) {
	obj.PkgLen1 = 2 + uint16(len(obj.content)) + 2
	obj.PkgLen2 = 2 + uint16(len(obj.content)) + 2

	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, obj.ReqHeader)
	err = binary.Write(buf, binary.LittleEndian, obj.market)
	buf.Write(obj.content)
	return buf.Bytes(), err
}

/*
0100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000011f85e34068747470733a2f2f626967352e6e65776f6e652e636f6d2e636e2f7a797968742f7a645f7a737a712e7a6970000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004150503a414c4c0d0a54494d453a303a30312d31353a30352c31353a30362d32333a35390d0a20202020c4facab9d3c3b5c4b0e6b1bebcb4bdabcda3d3c3a3acceaac1cbc4fab5c4d5fdb3a3cab9d3c32cc7ebbea1bfecc9fdd6c1d5d0c9ccd6a4c8af5043b0e6a1a30d0a20202020c8e7b9fbb2bbc4dcd7d4b6afc9fdbcb6a3acc7ebb5bdb9d9cdf868747470733a2f2f7777772e636d736368696e612e636f6d2fcfc2d4d8b0b2d7b0a3acd0bbd0bbc4fab5c4d6a7b3d6a3a100                                                                   年月日              年月日
*/
func (obj *SecurityCount) UnSerialize(header interface{}, data []byte) error {
	//serverInfo := utils.Utf8ToGbk(data[58:])
	//fmt.Println(fmt.Sprintf("服务器:%s;", serverInfo))
	//fmt.Println(hex.EncodeToString(data))
	//obj.Reply.Info = serverInfo

	//var tmp uint16
	//bytesBuffer := bytes.NewBuffer(data)
	//err := binary.Write(bytesBuffer, binary.LittleEndian, &tmp)
	//binary.LittleEndian.Uint16(data)
	obj.Reply.Count = binary.LittleEndian.Uint16(data[:2])
	return nil
}
