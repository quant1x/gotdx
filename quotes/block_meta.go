package quotes

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"gitee.com/quant1x/gotdx/proto"
	"gitee.com/quant1x/gotdx/util"
	"github.com/mymmsc/gox/encoding/binary/struc"
)

// 板块相关参数
const (
	BLOCK_ZHISHU      = "block_zs.dat" // 指数
	BLOCK_FENGGE      = "block_fg.dat" // 风格
	BLOCK_GAINIAN     = "block_gn.dat" // 概念
	BLOCK_DEFAULT     = "block.dat"    // 早期的板块数据文件, 与block_zs.dat
	BLOCK_CHUNKS_SIZE = 0x7530         // 板块文件默认一个请求包最大数据
)

// BlockMetaPackage 板块信息
type BlockMetaPackage struct {
	reqHeader  *StdRequestHeader
	respHeader *StdResponseHeader
	request    *BlockMetaRequest
	response   *BlockMeta
	contentHex string
}

// BlockMetaRequest 请求包
type BlockMetaRequest struct {
	BlockFile [40]byte // 板块文件名
}

// BlockMeta 响应包结构
type BlockMeta struct {
	Size      uint32   `struc:"uint32,little"`   // 尺寸
	C1        byte     `struc:"byte,little"`     // C1
	HashValue [32]byte `struc:"[32]byte,little"` // hash值
	C2        byte     `struc:"byte,little"`     // C2
}

func NewBlockMetaPackage() *BlockMetaPackage {
	pkg := new(BlockMetaPackage)
	pkg.reqHeader = new(StdRequestHeader)
	pkg.respHeader = new(StdResponseHeader)
	pkg.request = new(BlockMetaRequest)
	pkg.response = new(BlockMeta)

	//0c 1f 18 76 00 01 0b 00 0b 00 10 00 01 00
	//0c
	pkg.reqHeader.ZipFlag = 0x0c
	//1f 18 76 00
	pkg.reqHeader.SeqID = util.SeqID()
	//01
	pkg.reqHeader.PacketType = 0x01
	//0b 00
	//PkgLen1    uint16
	pkg.reqHeader.PkgLen1 = 0x002a
	//0b 00
	//PkgLen2    uint16
	pkg.reqHeader.PkgLen2 = 0x002a
	//10 00
	pkg.reqHeader.Method = proto.STD_MSG_BLOCK_META
	//pkg.contentHex = "0100" // 未解
	return pkg
}

func (obj *BlockMetaPackage) SetParams(req *BlockMetaRequest) {
	obj.request = req
}

func (obj *BlockMetaPackage) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, obj.reqHeader)
	b, err := hex.DecodeString(obj.contentHex)
	buf.Write(b)
	err = binary.Write(buf, binary.LittleEndian, obj.request)
	return buf.Bytes(), err
}

func (obj *BlockMetaPackage) UnSerialize(header interface{}, data []byte) error {
	obj.respHeader = header.(*StdResponseHeader)
	// 构造流
	buf := bytes.NewBuffer(data)
	var reply BlockMeta
	err := struc.Unpack(buf, &reply)
	if err != nil {
		return err
	}
	obj.response = &reply
	return nil
}

func (obj *BlockMetaPackage) Reply() interface{} {
	return obj.response
}
