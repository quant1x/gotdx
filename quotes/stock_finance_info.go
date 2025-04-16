package quotes

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"gitee.com/quant1x/exchange"
	"gitee.com/quant1x/gotdx/internal"
	"gitee.com/quant1x/gotdx/proto"
	"gitee.com/quant1x/gox/api"
	"gitee.com/quant1x/gox/encoding/binary/cstruct"
)

// FinanceInfoPackage 财务数据
type FinanceInfoPackage struct {
	reqHeader  *StdRequestHeader
	respHeader *StdResponseHeader
	request    *FinanceInfoRequest
	reply      *FinanceInfo
	contentHex string
}

type FinanceInfoRequest struct {
	Count  uint16  // 总数
	Market uint8   // 市场代码
	Code   [6]byte // 股票代码
}

// FinanceInfoReply 响应包结构
type FinanceInfoReply struct {
	Count uint16 //  总数
	//MarketType uint8   `struc:"uint8,little"`
	//Code   [6]byte `struc:"[6]byte,little"`
	First RawFinanceInfo
	//List  [2]RawFinanceInfo
}

// RawFinanceInfo 响应包结构
//
//	一次返回145个字节, 现在有136个字节, 空余9个字节分别是总数、市场和代码
//	Num每加1, 返回字节数增加143
type RawFinanceInfo struct {
	//Unknown1           [2]byte `struc:"[2]byte,little"`
	Market             uint8   `struc:"uint8,little"`
	Code               [6]byte `struc:"[6]byte,little"`
	LiuTongGuBen       float32 `struc:"float32,little"` // 流通股本
	Province           uint16  `struc:"uint16,little"`  // 所属省份
	Industry           uint16  `struc:"uint16,little"`
	UpdatedDate        uint32  `struc:"uint32,little"`
	IPODate            uint32  `struc:"uint32,little"`
	ZongGuBen          float32 `struc:"float32,little"`
	GuoJiaGu           float32 `struc:"float32,little"`
	FaQiRenFaRenGu     float32 `struc:"float32,little"`
	FaRenGu            float32 `struc:"float32,little"`
	BGu                float32 `struc:"float32,little"` // B股
	HGu                float32 `struc:"float32,little"` // H股
	ZhiGongGu          float32 `struc:"float32,little"`
	ZongZiChan         float32 `struc:"float32,little"`
	LiuDongZiChan      float32 `struc:"float32,little"`
	GuDingZiChan       float32 `struc:"float32,little"`
	WuXingZiChan       float32 `struc:"float32,little"`
	GuDongRenShu       float32 `struc:"float32,little"`
	LiuDongFuZhai      float32 `struc:"float32,little"`
	ChangQiFuZhai      float32 `struc:"float32,little"`
	ZiBenGongJiJin     float32 `struc:"float32,little"`
	JingZiChan         float32 `struc:"float32,little"`
	ZhuYingShouRu      float32 `struc:"float32,little"`
	ZhuYingLiRun       float32 `struc:"float32,little"`
	YingShouzhangkuan  float32 `struc:"float32,little"`
	YingYeLiRun        float32 `struc:"float32,little"`
	TouZiShouYu        float32 `struc:"float32,little"`
	JingYingXianJinLiu float32 `struc:"float32,little"`
	ZongXianJinLiu     float32 `struc:"float32,little"`
	CunHuo             float32 `struc:"float32,little"`
	LiRunZongHe        float32 `struc:"float32,little"`
	ShuiHouLiRun       float32 `struc:"float32,little"`
	JingLiRun          float32 `struc:"float32,little"`
	WeiFenLiRun        float32 `struc:"float32,little"`
	BaoLiu1            float32 `struc:"float32,little"`
	BaoLiu2            float32 `struc:"float32,little"`
	//BaoLiu3            [7]byte `struc:"[7]byte,little"`
}

type RawFinanceInfo1 struct {
	//Unknown1           [2]byte `struc:"[2]byte,little"`
	//MarketType             uint8   `struc:"uint8,little"`
	//Code               [6]byte `struc:"[6]byte,little"`
	LiuTongGuBen       float32 `struc:"float32,little"`
	Province           uint16  `struc:"uint16,little"`
	Industry           uint16  `struc:"uint16,little"`
	UpdatedDate        uint32  `struc:"uint32,little"`
	IPODate            uint32  `struc:"uint32,little"`
	ZongGuBen          float32 `struc:"float32,little"`
	GuoJiaGu           float32 `struc:"float32,little"`
	FaQiRenFaRenGu     float32 `struc:"float32,little"`
	FaRenGu            float32 `struc:"float32,little"`
	BGu                float32 `struc:"float32,little"`
	HGu                float32 `struc:"float32,little"`
	ZhiGongGu          float32 `struc:"float32,little"`
	ZongZiChan         float32 `struc:"float32,little"`
	LiuDongZiChan      float32 `struc:"float32,little"`
	GuDingZiChan       float32 `struc:"float32,little"`
	WuXingZiChan       float32 `struc:"float32,little"`
	GuDongRenShu       float32 `struc:"float32,little"`
	LiuDongFuZhai      float32 `struc:"float32,little"`
	ChangQiFuZhai      float32 `struc:"float32,little"`
	ZiBenGongJiJin     float32 `struc:"float32,little"`
	JingZiChan         float32 `struc:"float32,little"`
	ZhuYingShouRu      float32 `struc:"float32,little"`
	ZhuYingLiRun       float32 `struc:"float32,little"`
	YingShouZhangKuan  float32 `struc:"float32,little"`
	YingYeLiRun        float32 `struc:"float32,little"`
	TouZiShouYu        float32 `struc:"float32,little"`
	JingYingXianJinLiu float32 `struc:"float32,little"`
	ZongXianJinLiu     float32 `struc:"float32,little"`
	CunHuo             float32 `struc:"float32,little"`
	LiRunZongHe        float32 `struc:"float32,little"`
	ShuiHouLiRun       float32 `struc:"float32,little"`
	JingLiRun          float32 `struc:"float32,little"`
	WeiFenLiRun        float32 `struc:"float32,little"`
	BaoLiu1            float32 `struc:"float32,little"`
	BaoLiu2            float32 `struc:"float32,little"`
	BaoLiu3            [7]byte `struc:"[7]byte,little"`
}

type FinanceInfo struct {
	Code               string  `json:"code"`
	LiuTongGuBen       float64 `struc:"float32,little" json:"liu_tong_gu_ben"`
	Province           uint16  `struc:"uint16,little" json:"province"`
	Industry           uint16  `struc:"uint16,little" json:"industry"`
	UpdatedDate        uint32  `struc:"uint32,little" json:"updatedDate"`
	IPODate            uint32  `struc:"uint32,little" json:"ipo_date"`
	ZongGuBen          float64 `struc:"float32,little" json:"zong_gu_ben"`
	GuoJiaGu           float64 `struc:"float32,little" json:"guo_jia_gu"`
	FaQiRenFaRenGu     float64 `struc:"float32,little" json:"fa_qi_ren_fa_ren_gu"`
	FaRenGu            float64 `struc:"float32,little" json:"fa_ren_gu"`
	BGu                float64 `struc:"float32,little" json:"b_gu"`
	HGu                float64 `struc:"float32,little" json:"h_gu"`
	ZhiGongGu          float64 `struc:"float32,little" json:"zhi_gong_gu"`
	ZongZiChan         float64 `struc:"float32,little" json:"zong_zi_chan"`
	LiuDongZiChan      float64 `struc:"float32,little" json:"liu_dong_zi_chan"`
	GuDingZiChan       float64 `struc:"float32,little" json:"gu_ding_zi_chan"`
	WuXingZiChan       float64 `struc:"float32,little" json:"wu_xing_zi_chan"`
	GuDongRenShu       float64 `struc:"float32,little" json:"gu_dong_ren_shu"`
	LiuDongFuZhai      float64 `struc:"float32,little" json:"liu_dong_fu_zhai"`
	ChangQiFuZhai      float64 `struc:"float32,little" json:"chang_qi_fu_zhai"`
	ZiBenGongJiJin     float64 `struc:"float32,little" json:"zi_ben_gong_ji_jin"`
	JingZiChan         float64 `struc:"float32,little" json:"jing_zi_chan"`
	ZhuYingShouRu      float64 `struc:"float32,little" json:"zhu_ying_shou_ru"`
	ZhuYingLiRun       float64 `struc:"float32,little" json:"zhu_ying_li_run"`
	YingShouZhangKuan  float64 `struc:"float32,little" json:"ying_shou_zhang_kuan"`
	YingYeLiRun        float64 `struc:"float32,little" json:"ying_ye_li_run"`
	TouZiShouYu        float64 `struc:"float32,little" json:"tou_zi_shou_yu"`
	JingYingXianJinLiu float64 `struc:"float32,little" json:"jing_ying_xian_jin_liu"`
	ZongXianJinLiu     float64 `struc:"float32,little" json:"zong_xian_jin_liu"`
	CunHuo             float64 `struc:"float32,little" json:"cun_huo"`
	LiRunZongHe        float64 `struc:"float32,little" json:"li_run_zong_he"`
	ShuiHouLiRun       float64 `struc:"float32,little" json:"shui_hou_li_run"`
	JingLiRun          float64 `struc:"float32,little" json:"jing_li_run"`
	WeiFenLiRun        float64 `struc:"float32,little" json:"wei_fen_li_run"`
	MeiGuJingZiChan    float64 `struc:"float32,little" json:"mei_gu_jing_zi_chan"`
	BaoLiu2            float64 `struc:"float32,little" json:"bao_liu_2"`
}

// IsDelisting 是否退市
func (info FinanceInfo) IsDelisting() bool {
	return info.IPODate == 0 && info.ZongGuBen == 0 && info.LiuTongGuBen == 0
}

func NewFinanceInfoPackage() *FinanceInfoPackage {
	pkg := new(FinanceInfoPackage)
	pkg.reqHeader = new(StdRequestHeader)
	pkg.respHeader = new(StdResponseHeader)
	pkg.request = new(FinanceInfoRequest)
	pkg.reply = new(FinanceInfo)

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
	pkg.reqHeader.Method = proto.STD_MSG_FINANCE_INFO
	//pkg.contentHex = "0100" // 未解
	return pkg
}

func (obj *FinanceInfoPackage) SetParams(req *FinanceInfoRequest) {
	obj.request = req
}

func (obj *FinanceInfoPackage) Serialize() ([]byte, error) {
	//obj.reqHeader.PkgLen1 = 2 + 4 + 2
	//obj.reqHeader.PkgLen2 = 2 + 4 + 2

	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, obj.reqHeader)
	b, err := hex.DecodeString(obj.contentHex)
	buf.Write(b)
	err = binary.Write(buf, binary.LittleEndian, obj.request)
	return buf.Bytes(), err
}

func (obj *FinanceInfoPackage) UnSerialize(header interface{}, data []byte) error {
	obj.respHeader = header.(*StdResponseHeader)

	var reply FinanceInfoReply
	err := cstruct.Unpack(data, &reply)
	if err != nil {
		return err
	}
	var resp FinanceInfo
	raw := reply.First
	resp.Code = exchange.GetSecurityCode(obj.request.Market, api.Bytes2String(obj.request.Code[:]))
	resp.LiuTongGuBen = internal.NumberToFloat64(raw.LiuTongGuBen) * 10000
	resp.Province = raw.Province
	resp.Industry = raw.Industry
	resp.UpdatedDate = raw.UpdatedDate
	resp.IPODate = raw.IPODate
	resp.ZongGuBen = internal.NumberToFloat64(raw.ZongGuBen) * 10000
	resp.GuoJiaGu = internal.NumberToFloat64(raw.GuoJiaGu) * 10000
	resp.FaQiRenFaRenGu = internal.NumberToFloat64(raw.FaQiRenFaRenGu) * 10000
	resp.FaRenGu = internal.NumberToFloat64(raw.FaRenGu) * 10000
	resp.BGu = internal.NumberToFloat64(raw.BGu) * 10000
	resp.HGu = internal.NumberToFloat64(raw.HGu) * 10000
	resp.ZhiGongGu = internal.NumberToFloat64(raw.ZhiGongGu) * 10000
	resp.ZongZiChan = internal.NumberToFloat64(raw.ZongZiChan) * 10000
	resp.LiuDongZiChan = internal.NumberToFloat64(raw.LiuDongZiChan) * 10000
	resp.GuDingZiChan = internal.NumberToFloat64(raw.GuDingZiChan) * 10000
	resp.WuXingZiChan = internal.NumberToFloat64(raw.WuXingZiChan) * 10000
	resp.GuDongRenShu = internal.NumberToFloat64(raw.GuDongRenShu)
	resp.LiuDongFuZhai = internal.NumberToFloat64(raw.LiuDongFuZhai) * 10000
	resp.ChangQiFuZhai = internal.NumberToFloat64(raw.ChangQiFuZhai) * 10000
	resp.ZiBenGongJiJin = internal.NumberToFloat64(raw.ZiBenGongJiJin) * 10000
	resp.JingZiChan = internal.NumberToFloat64(raw.JingZiChan) * 10000
	resp.ZhuYingShouRu = internal.NumberToFloat64(raw.ZhuYingShouRu) * 10000
	resp.ZhuYingLiRun = internal.NumberToFloat64(raw.ZhuYingLiRun) * 10000
	resp.YingShouZhangKuan = internal.NumberToFloat64(raw.YingShouzhangkuan) * 10000
	resp.YingYeLiRun = internal.NumberToFloat64(raw.YingYeLiRun) * 10000
	resp.TouZiShouYu = internal.NumberToFloat64(raw.TouZiShouYu) * 10000
	resp.JingYingXianJinLiu = internal.NumberToFloat64(raw.JingYingXianJinLiu) * 10000
	resp.ZongXianJinLiu = internal.NumberToFloat64(raw.ZongXianJinLiu) * 10000
	resp.CunHuo = internal.NumberToFloat64(raw.CunHuo) * 10000
	resp.LiRunZongHe = internal.NumberToFloat64(raw.LiRunZongHe) * 10000
	resp.ShuiHouLiRun = internal.NumberToFloat64(raw.ShuiHouLiRun) * 10000
	resp.JingLiRun = internal.NumberToFloat64(raw.JingLiRun) * 10000
	resp.WeiFenLiRun = internal.NumberToFloat64(raw.WeiFenLiRun) * 10000
	resp.MeiGuJingZiChan = internal.NumberToFloat64(raw.BaoLiu1) * 10000
	resp.BaoLiu2 = internal.NumberToFloat64(raw.BaoLiu2)
	obj.reply = &resp
	return nil
}

func (obj *FinanceInfoPackage) Reply() interface{} {
	return obj.reply
}
