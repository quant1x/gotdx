package quotes

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"

	"gitee.com/quant1x/gotdx/internal"
	"gitee.com/quant1x/gotdx/proto"
	"gitee.com/quant1x/gox/encoding/binary/struc"
)

var (
	XDXR_CATEGORY_MAPPING = map[int]string{
		1:  "除权除息",
		2:  "送配股上市",
		3:  "非流通股上市",
		4:  "未知股本变动",
		5:  "股本变化",
		6:  "增发新股",
		7:  "股份回购",
		8:  "增发新股上市",
		9:  "转配股上市",
		10: "可转债上市",
		11: "扩缩股",
		12: "非流通股缩股",
		13: "送认购权证",
		14: "送认沽权证",
	}
)

// XdxrInfoPackage 除权除息
type XdxrInfoPackage struct {
	reqHeader  *StdRequestHeader
	respHeader *StdResponseHeader
	request    *XdxrInfoRequest
	response   *XdxrInfoReply
	reply      []XdxrInfo
	contentHex string
}

type XdxrInfoRequest struct {
	//Count  uint16  // 总数
	Market uint8   // 市场代码
	Code   [6]byte // 股票代码
}

// XdxrInfoReply 响应包结构
type XdxrInfoReply struct {
	Unknown []byte        `struc:"[9]byte,little"`            // 未知
	Count   uint16        `struc:"uint16,little,sizeof=List"` //  总数
	List    []RawXdxrInfo `struc:"[29]byte, little"`          // [29]byte和title中间必须要有一个空格
}

type RawXdxrInfo struct {
	Market   int    `struc:"uint8,little"`   // 市场代码
	Code     string `struc:"[6]byte,little"` // 股票代码
	Unknown  int    `struc:"uint8,little"`   // 未知
	Date     uint32 `struc:"uint32,little"`  // 日期
	Category int    `struc:"uint8,little"`   // 类型
	Data     []byte `struc:"[16]byte,little"`
}

type XdxrInfo struct {
	Date          string  // 日期
	Category      int     // 类型
	Name          string  // 类型名称
	FenHong       float64 // 分红
	PeiGuJia      float64 // 配股价
	SongZhuanGu   float64 // 送转股
	PeiGu         float64 // 配股
	SuoGu         float64 // 缩股
	QianLiuTong   float64 // 前流通
	HouLiuTong    float64 // 后流通
	QianZongGuBen float64 // 前总股本
	HouZongGuBen  float64 // 后总股本
	FenShu        float64 // 份数
	XingQuanJia   float64 // 行权价
}

// IsCapitalChange 是否股本变化
func (x *XdxrInfo) IsCapitalChange() bool {
	switch x.Category {
	case 1, 11, 12, 13, 14:
		return false
	default:
		if x.HouLiuTong > 0 && x.HouZongGuBen > 0 {
			return true
		}
	}
	return false
}

// Adjust 返回复权回调函数 factor
func (x *XdxrInfo) Adjust() func(p float64) float64 {
	songZhuangu := x.SongZhuanGu
	peiGu := x.PeiGu
	suoGu := x.SuoGu
	xdxrGuShu := (songZhuangu + peiGu - suoGu) / 10
	fenHong := x.FenHong
	peiGuJia := x.PeiGuJia
	xdxrFenHong := (peiGuJia*peiGu - fenHong) / 10

	factor := func(p float64) float64 {
		v := (p + xdxrFenHong) / (1 + xdxrGuShu)
		//return num.Decimal(v)
		return v
	}
	return factor
}

func NewXdxrInfoPackage() *XdxrInfoPackage {
	pkg := new(XdxrInfoPackage)
	pkg.reqHeader = new(StdRequestHeader)
	pkg.respHeader = new(StdResponseHeader)
	pkg.request = new(XdxrInfoRequest)
	pkg.response = new(XdxrInfoReply)

	//0c 1f 18 76 00 01 0b 00 0b 00 10 00 01 00
	//0c
	pkg.reqHeader.ZipFlag = proto.FlagNotZipped
	//1f 18 76 00
	pkg.reqHeader.SeqID = internal.SequenceId()
	//01
	pkg.reqHeader.PacketType = 0x01
	//0b 00
	//PkgLen1    uint16
	pkg.reqHeader.PkgLen1 = 0x000b
	//0b 00
	//PkgLen2    uint16
	pkg.reqHeader.PkgLen2 = 0x000b
	//10 00
	pkg.reqHeader.Method = proto.STD_MSG_XDXR_INFO
	pkg.contentHex = "0100" // 未解
	return pkg
}

func (obj *XdxrInfoPackage) SetParams(req *XdxrInfoRequest) {
	//req.Count = 1
	obj.request = req
}

func (obj *XdxrInfoPackage) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, obj.reqHeader)
	b, err := hex.DecodeString(obj.contentHex)
	buf.Write(b)
	err = binary.Write(buf, binary.LittleEndian, obj.request)
	return buf.Bytes(), err
}

func (obj *XdxrInfoPackage) UnSerialize(header interface{}, data []byte) error {
	obj.respHeader = header.(*StdResponseHeader)
	// 构造流
	buf := bytes.NewBuffer(data)
	var reply XdxrInfoReply
	err := struc.Unpack(buf, &reply)
	if err != nil {
		return err
	}
	var list = []XdxrInfo{}
	for _, v := range reply.List {
		year, month, day, hour, minute := internal.GetDatetimeFromUint32(9, v.Date, 0)
		xdxr := XdxrInfo{
			//Date           string // 日期
			Date: fmt.Sprintf("%04d-%02d-%02d", year, month, day),
			//Category       int    // 类型
			Category: v.Category,
			//Name           string // 类型名称
			Name: XDXR_CATEGORY_MAPPING[v.Category],
			//FenHong        int    // 分红
			//PeiGuJia       int    // 配股价
			//SongZhuanGu    int    // 送转股
			//PeiGu          int    // 配股
			//SuoGu          int    // 锁骨
			//QianLiuTong int    // 盘前流通
			//HouLiuTong  int    // 盘后流通
			//QianZongGuBen  int    // 前总股本
			//HouZongGuBen   int    // 后总股本
			//FenShu         int    // 份数
			//XingGuanJia    int    // 行权价
		}
		switch xdxr.Category {
		case 1:
			var f float32
			pos := 0
			_ = binary.Read(bytes.NewBuffer(v.Data[pos:pos+4]), binary.LittleEndian, &f)
			xdxr.FenHong = float64(f)
			pos += 4
			_ = binary.Read(bytes.NewBuffer(v.Data[pos:pos+4]), binary.LittleEndian, &f)
			xdxr.PeiGuJia = float64(f)
			pos += 4
			_ = binary.Read(bytes.NewBuffer(v.Data[pos:pos+4]), binary.LittleEndian, &f)
			xdxr.SongZhuanGu = float64(f)
			pos += 4
			_ = binary.Read(bytes.NewBuffer(v.Data[pos:pos+4]), binary.LittleEndian, &f)
			xdxr.PeiGu = float64(f)
		case 11, 12:
			var f float32
			pos := 8
			_ = binary.Read(bytes.NewBuffer(v.Data[pos:pos+4]), binary.LittleEndian, &f)
			xdxr.SuoGu = float64(f)
		case 13, 14:
			var f float32
			pos := 0
			_ = binary.Read(bytes.NewBuffer(v.Data[pos:pos+4]), binary.LittleEndian, &f)
			xdxr.XingQuanJia = float64(f)
			pos = 8
			_ = binary.Read(bytes.NewBuffer(v.Data[pos:pos+4]), binary.LittleEndian, &f)
			xdxr.FenShu = float64(f)
		default:
			var i uint32
			pos := 0
			_ = binary.Read(bytes.NewBuffer(v.Data[pos:pos+4]), binary.LittleEndian, &i)
			xdxr.QianLiuTong = __get_v(i)
			pos += 4
			_ = binary.Read(bytes.NewBuffer(v.Data[pos:pos+4]), binary.LittleEndian, &i)
			xdxr.QianZongGuBen = __get_v(i)
			pos += 4
			_ = binary.Read(bytes.NewBuffer(v.Data[pos:pos+4]), binary.LittleEndian, &i)
			xdxr.HouLiuTong = __get_v(i)
			pos += 4
			_ = binary.Read(bytes.NewBuffer(v.Data[pos:pos+4]), binary.LittleEndian, &i)
			xdxr.HouZongGuBen = __get_v(i)
		}
		list = append(list, xdxr)
		_ = hour
		_ = minute
	}
	obj.reply = list
	return nil
}

func (obj *XdxrInfoPackage) Reply() interface{} {
	return obj.reply
}

func __get_v(v uint32) float64 {
	if v == 0 {
		return 0
	}
	return internal.IntToFloat64(int(v))
}
