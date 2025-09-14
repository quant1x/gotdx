package std

// 获取股票列表
import (
	"gitee.com/quant1x/exchange"
	"gitee.com/quant1x/gotdx/internal"
)

// 请求包结构
type FinanceInfoRequest struct {
	// struc不允许slice解析，只允许包含长度的array，该长度可根据hex字符串计算
	Unknown1 []byte `struc:"[14]byte"`
	// pytdx中使用struct.Pack进行反序列化
	// 其中<H等价于这里的struc:"uint16,little"
	// <I等价于struc:"uint32,little"
	Market exchange.MarketType `struc:"uint8,little" json:"market"`
	Code   string              `struc:"[6]byte,little" json:"code"`
}

// 请求包序列化输出
func (req *FinanceInfoRequest) Marshal() ([]byte, error) {
	return DefaultMarshal(req)
}

// 响应包结构
type FinanceInfoResponseRaw struct {
	Unknown1 []byte `struc:"[2]byte,little" json:"unknown1"`
	Market   int    `struc:"uint8,little" json:"market"`
	Code     string `struc:"[6]byte,little" json:"code"`
	//Info     FinanceInfo `struc:"[136]byte,little" json:"info"`
	LiuTongGuBen       float32 `struc:"float32,little" json:"liuTongGuBen"`
	Province           uint16  `struc:"uint16,little" json:"province"`
	Industry           uint16  `struc:"uint16,little" json:"industry"`
	UpdatedDate        uint32  `struc:"uint32,little" json:"updatedDate"`
	IPODate            uint32  `struc:"uint32,little" json:"ipo_date"`
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
	Yingshouzhangkuan  float32 `struc:"float32,little"`
	YingyeLiRun        float32 `struc:"float32,little"`
	TouZiShouYu        float32 `struc:"float32,little"`
	JingYingxianJinLiu float32 `struc:"float32,little"`
	ZongXianJinLiu     float32 `struc:"float32,little"`
	CunHuo             float32 `struc:"float32,little"`
	LiRunZongHe        float32 `struc:"float32,little"`
	ShuiHouLiRun       float32 `struc:"float32,little"`
	JingLiRun          float32 `struc:"float32,little"`
	WeiFenLiRun        float32 `struc:"float32,little"`
	BaoLiu1            float32 `struc:"float32,little"`
	BaoLiu2            float32 `struc:"float32,little"`
}

func (resp *FinanceInfoResponseRaw) Unmarshal(data []byte) error {
	return DefaultUnmarshal(data, &resp)
}

type FinanceInfo struct {
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
	YingyeLiRun        float64 `struc:"float32,little" json:"yingye_li_run"`
	TouZiShouYu        float64 `struc:"float32,little" json:"tou_zi_shou_yu"`
	JingYingxianJinLiu float64 `struc:"float32,little" json:"jing_yingxian_jin_liu"`
	ZongXianJinLiu     float64 `struc:"float32,little" json:"zong_xian_jin_liu"`
	CunHuo             float64 `struc:"float32,little" json:"cun_huo"`
	LiRunZongHe        float64 `struc:"float32,little" json:"li_run_zong_he"`
	ShuiHouLiRun       float64 `struc:"float32,little" json:"shui_hou_li_run"`
	JingLiRun          float64 `struc:"float32,little" json:"jing_li_run"`
	WeiFenLiRun        float64 `struc:"float32,little" json:"wei_fen_li_run"`
	MeiGuJingZiChan    float64 `struc:"float32,little" json:"mei_gu_jing_zi_chan"`
	BaoLiu2            float64 `struc:"float32,little" json:"bao_liu_2"`
}

// 响应包结构
type FinanceInfoResponse struct {
	FinanceInfo
}

// 内部套用原始结构解析，外部为经过解析之后的响应信息
func (resp *FinanceInfoResponse) Unmarshal(data []byte) error {
	var raw FinanceInfoResponseRaw
	err := raw.Unmarshal(data)
	if err != nil {
		return err
	}
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
	resp.YingShouZhangKuan = internal.NumberToFloat64(raw.Yingshouzhangkuan) * 10000
	resp.YingyeLiRun = internal.NumberToFloat64(raw.YingyeLiRun) * 10000
	resp.TouZiShouYu = internal.NumberToFloat64(raw.TouZiShouYu) * 10000
	resp.JingYingxianJinLiu = internal.NumberToFloat64(raw.JingYingxianJinLiu) * 10000
	resp.ZongXianJinLiu = internal.NumberToFloat64(raw.ZongXianJinLiu) * 10000
	resp.CunHuo = internal.NumberToFloat64(raw.CunHuo) * 10000
	resp.LiRunZongHe = internal.NumberToFloat64(raw.LiRunZongHe) * 10000
	resp.ShuiHouLiRun = internal.NumberToFloat64(raw.ShuiHouLiRun) * 10000
	resp.JingLiRun = internal.NumberToFloat64(raw.JingLiRun) * 10000
	resp.WeiFenLiRun = internal.NumberToFloat64(raw.WeiFenLiRun) * 10000
	resp.MeiGuJingZiChan = internal.NumberToFloat64(raw.BaoLiu1) * 10000
	resp.BaoLiu2 = internal.NumberToFloat64(raw.BaoLiu2)
	return nil
}

// todo: 检测market是否为合法值
func NewFinanceInfoRequest(market exchange.MarketType, code string) (*FinanceInfoRequest, error) {
	request := &FinanceInfoRequest{
		Unknown1: internal.HexString2Bytes("0c 1f 18 76 00 01 0b 00 0b 00 10 00 01 00"),
		Market:   market,
		Code:     code,
	}
	return request, nil
}

func NewFinanceInfo(market exchange.MarketType, code string) (*FinanceInfoRequest, *FinanceInfoResponse, error) {
	var response FinanceInfoResponse
	var request, err = NewFinanceInfoRequest(market, code)
	return request, &response, err
}
