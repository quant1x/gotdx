package quotes

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"gitee.com/quant1x/gotdx/internal"
	"gitee.com/quant1x/gotdx/proto"
	"gitee.com/quant1x/gox/encoding/binary/struc"
)

// BlockInfoPackage 板块信息
type BlockInfoPackage struct {
	reqHeader  *StdRequestHeader
	respHeader *StdResponseHeader
	request    *BlockInfoRequest
	response   *BlockInfoResponse
	contentHex string
}

// BlockInfoRequest 请求包
type BlockInfoRequest struct {
	Start     uint32    `struc:"uint32,little"`
	Size      uint32    `struc:"uint32,little"`
	BlockFile [100]byte `struc:"[100]byte,little"` // 板块文件名
}

type BlockInfo struct {
	BlockName  string
	BlockType  uint16
	StockCount uint16
	Codelist   []string
}

type BlockInfoResponse struct {
	Size uint32 `struc:"uint32,little"`
	Data []byte `struc:"sizefrom=Size"`
}

type BlockInfoReply struct {
	BlockNum uint16      `struc:"uint16,little"` // 板块个数
	Block    []BlockInfo // 板块列表
}

func NewBlockInfoPackage() *BlockInfoPackage {
	pkg := new(BlockInfoPackage)
	pkg.reqHeader = new(StdRequestHeader)
	pkg.respHeader = new(StdResponseHeader)
	pkg.request = new(BlockInfoRequest)
	pkg.response = new(BlockInfoResponse)

	//0c 1f 18 76 00 01 0b 00 0b 00 10 00 01 00
	//0c
	pkg.reqHeader.ZipFlag = 0x0c
	//1f 18 76 00
	pkg.reqHeader.SeqID = internal.SeqID()
	//01
	pkg.reqHeader.PacketType = 0x01
	//0b 00
	//PkgLen1    uint16
	pkg.reqHeader.PkgLen1 = 0x006e
	//0b 00
	//PkgLen2    uint16
	pkg.reqHeader.PkgLen2 = 0x006e
	//10 00
	pkg.reqHeader.Method = proto.STD_MSG_BLOCK_DATA
	//pkg.contentHex = "0100" // 未解
	return pkg
}

func (obj *BlockInfoPackage) SetParams(req *BlockInfoRequest) {
	obj.request = req
}

func (obj *BlockInfoPackage) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, obj.reqHeader)
	b, err := hex.DecodeString(obj.contentHex)
	buf.Write(b)
	err = binary.Write(buf, binary.LittleEndian, obj.request)
	return buf.Bytes(), err
}

func (obj *BlockInfoPackage) UnSerialize(header interface{}, data []byte) error {
	obj.respHeader = header.(*StdResponseHeader)
	// 构造流
	buf := bytes.NewBuffer(data)
	var reply BlockInfoResponse
	err := struc.Unpack(buf, &reply)
	if err != nil {
		return err
	}
	obj.response = &reply
	return nil
}

func (obj *BlockInfoPackage) Reply() interface{} {
	return obj.response
}
