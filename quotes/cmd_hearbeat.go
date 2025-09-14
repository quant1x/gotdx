package quotes

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"

	"github.com/quant1x/gotdx/internal"
	"github.com/quant1x/gotdx/proto"
)

// 心跳包, command: 0004
// 0c76002800 02 0200 0200 0400
// b1cb74000c760028000004000a000a00 00000000000000000000

type HeartBeatPackage struct {
	reqHeader  *StdRequestHeader
	request    *HeartBeatRequest
	respHeader *StdResponseHeader
	reply      *HeartBeatReply

	contentHex string
}

type HeartBeatRequest struct {
}

type HeartBeatReply struct {
	Info string // 10个字节的消息, 未解
}

func NewHeartBeat() *HeartBeatPackage {
	obj := new(HeartBeatPackage)
	obj.reqHeader = new(StdRequestHeader)
	obj.respHeader = new(StdResponseHeader)
	obj.request = new(HeartBeatRequest)
	obj.reply = new(HeartBeatReply)

	obj.reqHeader.ZipFlag = proto.FlagNotZipped
	obj.reqHeader.SeqID = internal.SequenceId()
	obj.reqHeader.PacketType = 0x02
	obj.reqHeader.Method = proto.STD_MSG_HEARTBEAT
	return obj
}

func (obj *HeartBeatPackage) Serialize() ([]byte, error) {
	b, err := hex.DecodeString(obj.contentHex)

	obj.reqHeader.PkgLen1 = 2 + uint16(len(b))
	obj.reqHeader.PkgLen2 = 2 + uint16(len(b))

	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.LittleEndian, obj.reqHeader)

	buf.Write(b)
	return buf.Bytes(), err
}

func (obj *HeartBeatPackage) UnSerialize(header interface{}, data []byte) error {
	obj.respHeader = header.(*StdResponseHeader)
	serverInfo := internal.Utf8ToGbk(data[:])
	obj.reply.Info = serverInfo
	return nil
}

func (obj *HeartBeatPackage) Reply() interface{} {
	return obj.reply
}
