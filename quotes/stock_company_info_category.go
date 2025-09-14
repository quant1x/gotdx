package quotes

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"

	"github.com/quant1x/gotdx/internal"
	"github.com/quant1x/gotdx/proto"
	"github.com/quant1x/x/encoding/binary/struc"
)

// CompanyInfoCategoryPackage 企业基本信息
type CompanyInfoCategoryPackage struct {
	reqHeader  *StdRequestHeader
	respHeader *StdResponseHeader
	request    *CompanyInfoCategoryRequest
	reply      []CompanyInfoCategory
	contentHex string
}

type CompanyInfoCategoryRequest struct {
	Market  uint16  // 市场代码
	Code    [6]byte // 股票代码
	Unknown uint32  // 未知数据
}

// CompanyInfoCategoryReply 响应包结构,
type CompanyInfoCategoryReply struct {
	Count uint16                   `struc:"uint16,little,sizeof=Data"` // 词条总数
	Data  []RawCompanyInfoCategory `struc:"[152]byte, little"`         // 词条数据
}

// RawCompanyInfoCategory 响应包结构
type RawCompanyInfoCategory struct {
	Name     []byte `struc:"[64]byte,little"` // 名称
	Filename []byte `struc:"[80]byte,little"` // 文件名
	Offset   uint32 `struc:"uint32,little"`   // 偏移量
	Length   uint32 `struc:"uint32,little"`   // 长度
}

type CompanyInfoCategory struct {
	Name     string `struc:"[64]byte,little" dataframe:"name"`     // 名称
	Filename string `struc:"[80]byte,little" dataframe:"filename"` // 文件名
	Offset   uint32 `struc:"uint32,little" dataframe:"offset"`     // 偏移量
	Length   uint32 `struc:"uint32,little" dataframe:"length"`     // 长度
}

func NewCompanyInfoCategoryPackage() *CompanyInfoCategoryPackage {
	pkg := new(CompanyInfoCategoryPackage)
	pkg.reqHeader = new(StdRequestHeader)
	pkg.respHeader = new(StdResponseHeader)
	//pkg.request = new(CompanyInfoCategoryRequest)
	//pkg.reply = new(CompanyInfoCategory)

	//0c 1f 18 76 00 01 0b 00 0b 00 10 00 01 00
	//0c
	pkg.reqHeader.ZipFlag = proto.FlagNotZipped
	//1f 18 76 00
	pkg.reqHeader.SeqID = internal.SequenceId()
	//01
	pkg.reqHeader.PacketType = 0x01
	//0b 00
	//PkgLen1    uint16
	pkg.reqHeader.PkgLen1 = 0x000e
	//0b 00
	//PkgLen2    uint16
	pkg.reqHeader.PkgLen2 = 0x000e
	//10 00
	pkg.reqHeader.Method = proto.STD_MSG_COMPANY_CATEGORY
	//pkg.contentHex = "0100" // 未解
	return pkg
}

func (obj *CompanyInfoCategoryPackage) SetParams(req *CompanyInfoCategoryRequest) {
	obj.request = req
}

func (obj *CompanyInfoCategoryPackage) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, obj.reqHeader)
	b, err := hex.DecodeString(obj.contentHex)
	buf.Write(b)
	err = binary.Write(buf, binary.LittleEndian, obj.request)
	return buf.Bytes(), err
}

func (obj *CompanyInfoCategoryPackage) UnSerialize(header interface{}, data []byte) error {
	obj.respHeader = header.(*StdResponseHeader)

	var reply CompanyInfoCategoryReply
	buf := bytes.NewBuffer(data)
	err := struc.Unpack(buf, &reply)
	if err != nil {
		return err
	}
	//category := make(map[string]CompanyInfoCategory)
	list := []CompanyInfoCategory{}
	for _, v := range reply.Data {
		info := CompanyInfoCategory{
			Name:     internal.Utf8ToGbk(v.Name[:]),
			Filename: internal.Utf8ToGbk(v.Filename[:]),
			Offset:   v.Offset,
			Length:   v.Length,
		}
		//category[info.Name] = info
		list = append(list, info)
	}
	obj.reply = list
	return nil
}

func (obj *CompanyInfoCategoryPackage) Reply() interface{} {
	return obj.reply
}
