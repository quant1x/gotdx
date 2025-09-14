package quotes

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"strings"

	"github.com/quant1x/exchange"
	"github.com/quant1x/gotdx/internal"
	"github.com/quant1x/gotdx/proto"
	"github.com/quant1x/x/encoding/binary/struc"
	"github.com/quant1x/x/util/linkedhashmap"
)

// CompanyInfoContentPackage 企业基本信息
type CompanyInfoContentPackage struct {
	reqHeader  *StdRequestHeader
	respHeader *StdResponseHeader
	request    *CompanyInfoContentRequest
	reply      *CompanyInfoContent
	contentHex string
}

type CompanyInfoContentRequest struct {
	Market   uint16   // 市场代码
	Code     [6]byte  // 股票代码
	Unknown1 uint16   // 未知数据
	Filename [80]byte // 文件名
	Offset   uint32   // 偏移量
	Length   uint32   // 长度
	Unknown2 uint32   // 未知数据
}

// CompanyInfoContentReply 响应包结构,
type CompanyInfoContentReply struct {
	Market   uint16 `struc:"uint16,little"`  // 市场代码
	Code     string `struc:"[6]byte,little"` // 股票代码
	Unknown1 []byte `struc:"[2]byte,little"` // 未知
	Length   uint16 `struc:"uint16,little"`  // 词条总数
	Data     []byte `struc:"sizefrom=Length"`
}

type CompanyInfoContent struct {
	Market  exchange.MarketType `dataframe:"market"`  // 市场代码
	Code    string              `dataframe:"code"`    // 短码
	Name    string              `dataframe:"name"`    // 名称
	Length  uint32              `dataframe:"length"`  // 长度
	Content string              `dataframe:"content"` // 内容
}

func (this *CompanyInfoContent) Map(unit string) *linkedhashmap.Map {
	mapInfo := linkedhashmap.New()
	c := strings.ReplaceAll(this.Content, "-\\u003e", "->")
	//arr := strings.Split(c, "\\r\\n\\r\\n")
	arr := strings.Split(c, "\r\n\r\n")
	for i, block := range arr {
		block = strings.TrimSpace(block)
		//v = strings.ReplaceAll(v, " ", "")
		//v = strings.ReplaceAll(v, "│\\r\\n││", "")
		//v = strings.Trim(v, "│")
		//fmt.Println(i, block)
		if i > 0 && strings.Index(block, unit) >= 0 {
			arr := strings.Split(block, "\r\n")
			block = ""
			for _, v := range arr {
				if strings.Index(v, unit) >= 0 {
					continue
				}
				if strings.Index(v, "┌") >= 0 && strings.Index(v, "┐") >= 0 {
					continue
				}
				if strings.Index(v, "└") >= 0 && strings.Index(v, "┘") >= 0 {
					continue
				}
				if strings.Index(v, "├") >= 0 && strings.Index(v, "┤") >= 0 {
					continue
				}
				v = strings.TrimLeft(v, "│")
				v = strings.TrimRight(v, "│")
				v = strings.TrimSpace(v)
				v = strings.ReplaceAll(v, "│", "|")
				//v = strings.TrimLeft(v, "|")
				if v[0] == '|' {
					block += v[1:]
				} else {
					block += "|" + v
				}
			}
			list := strings.Split(block[1:], "|")

			for k := 0; k < len(list); k += 2 {
				key := strings.TrimSpace(list[k])
				value := strings.TrimSpace(list[k+1])
				mapInfo.Put(key, value)
			}
			//mapInfo.Each(func(key interface{}, value interface{}) {
			//	fmt.Println(key, value)
			//})

			break
		}
	}
	return mapInfo
}

func NewCompanyInfoContentPackage() *CompanyInfoContentPackage {
	pkg := new(CompanyInfoContentPackage)
	pkg.reqHeader = new(StdRequestHeader)
	pkg.respHeader = new(StdResponseHeader)
	//pkg.request = new(CompanyInfoContentRequest)
	//pkg.reply = new(CompanyInfoContent)

	//0c 1f 18 76 00 01 0b 00 0b 00 10 00 01 00
	//0c
	pkg.reqHeader.ZipFlag = proto.FlagNotZipped
	//1f 18 76 00
	pkg.reqHeader.SeqID = internal.SequenceId()
	//01
	pkg.reqHeader.PacketType = 0x01
	//0b 00
	//PkgLen1    uint16
	pkg.reqHeader.PkgLen1 = 0x0068
	//0b 00
	//PkgLen2    uint16
	pkg.reqHeader.PkgLen2 = 0x0068
	//10 00
	pkg.reqHeader.Method = proto.STD_MSG_COMPANY_CONTENT
	//pkg.contentHex = "0100" // 未解
	return pkg
}

func (obj *CompanyInfoContentPackage) SetParams(req *CompanyInfoContentRequest) {
	obj.request = req
}

func (obj *CompanyInfoContentPackage) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, obj.reqHeader)
	b, err := hex.DecodeString(obj.contentHex)
	buf.Write(b)
	err = binary.Write(buf, binary.LittleEndian, obj.request)
	return buf.Bytes(), err
}

func (obj *CompanyInfoContentPackage) UnSerialize(header interface{}, data []byte) error {
	obj.respHeader = header.(*StdResponseHeader)

	var reply CompanyInfoContentReply
	buf := bytes.NewBuffer(data)
	err := struc.Unpack(buf, &reply)
	if err != nil {
		return err
	}
	response := CompanyInfoContent{
		Market:  exchange.MarketType(reply.Market),
		Code:    reply.Code,
		Length:  uint32(reply.Length),
		Content: internal.Utf8ToGbk(reply.Data),
	}
	obj.reply = &response
	return nil
}

func (obj *CompanyInfoContentPackage) Reply() interface{} {
	return obj.reply
}
