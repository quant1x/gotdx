package quotes

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"gitee.com/quant1x/gotdx/internal"
	"gitee.com/quant1x/gotdx/proto"
	"gitee.com/quant1x/gotdx/trading"
	"math"
)

const (
	TDX_SECURITY_QUOTES_MAX = 80 // 单次最大获取80条实时数据
)

type TradeState int8

const (
	TDX_SECURITY_TRADE_STATE_DELISTING TradeState = iota // 终止上市
	TDX_SECURITY_TRADE_STATE_NORMAL                      // 正常交易
	TDX_SECURITY_TRADE_STATE_SUSPEND                     // 停牌
)

// SecurityQuotesPackage 盘口五档报价
type SecurityQuotesPackage struct {
	reqHeader  *StdRequestHeader
	respHeader *StdResponseHeader
	request    *SecurityQuotesRequest
	reply      *SecurityQuotesReply
	mapCode    map[string]Stock // 序号和代码临时映射关系

	contentHex string
}

type Stock struct {
	Market uint8
	Code   string
}

type SecurityQuotesRequest struct {
	Count     uint16
	StockList []Stock
}

type SecurityQuotesReply struct {
	Count uint16
	List  []SecurityQuote
}

type SecurityQuote struct {
	State           TradeState // 交易状态
	Market          uint8      // 市场
	Code            string     // 代码
	Active1         uint16     // 活跃度
	Price           float64    // 现价
	LastClose       float64    // 昨收
	Open            float64    // 开盘
	High            float64    // 最高
	Low             float64    // 最低
	ServerTime      string     // 时间
	ReversedBytes0  int        // 保留(时间 ServerTime)
	ReversedBytes1  int        // 保留
	Vol             int        // 总量
	CurVol          int        // 个股-现成交量,板块指数-现成交额
	Amount          float64    // 总金额
	SVol            int        // 个股有效-内盘
	BVol            int        // 个股有效-外盘
	IndexOpenAmount int        // 指数有效-集合竞价成交金额=开盘成交金额
	StockOpenAmount int        // 个股有效-集合竞价成交金额=开盘成交金额
	OpenVolume      int        // 集合竞价-开盘量, 单位是股
	CloseVolume     int        // 集合竞价-收盘量, 单位是股
	IndexUp         int        // 指数有效-上涨数
	IndexUpLimit    int        // 指数有效-涨停数
	IndexDown       int        // 指数有效-下跌数
	IndexDownLimit  int        // 指数有效-跌停数
	Bid1            float64    // 个股-委买价1
	Ask1            float64    // 个股-委卖价1
	BidVol1         int        // 个股-委买量1 板块-上涨数
	AskVol1         int        // 个股-委卖量1 板块-下跌数
	Bid2            float64    // 个股-委买价2
	Ask2            float64    // 个股-委卖价2
	BidVol2         int        // 个股-委买量2 板块-涨停数
	AskVol2         int        // 个股-委卖量2 板块-跌停数
	Bid3            float64    // 个股-委买价3
	Ask3            float64    // 个股-委卖价3
	BidVol3         int        // 个股-委买量3
	AskVol3         int        // 个股-委卖量3
	Bid4            float64    // 个股-委买价4
	Ask4            float64    // 个股-委卖价4
	BidVol4         int        // 个股-委买量4
	AskVol4         int        // 个股-委卖量4
	Bid5            float64    // 个股-委买价5
	Ask5            float64    // 个股-委卖价5
	BidVol5         int        // 个股-委买量5
	AskVol5         int        // 个股-委卖量5
	ReversedBytes4  uint16     // 保留
	ReversedBytes5  int        // 保留
	ReversedBytes6  int        // 保留
	ReversedBytes7  int        // 保留
	ReversedBytes8  int        // 保留
	Rate            float64    // 涨速
	Active2         uint16     // 活跃度, 如果是指数则为0, 个股同Active1
}

type Level struct {
	Price float64
	Vol   int
}

func NewSecurityQuotesPackage() *SecurityQuotesPackage {
	obj := new(SecurityQuotesPackage)
	obj.reqHeader = new(StdRequestHeader)
	obj.respHeader = new(StdResponseHeader)
	obj.request = new(SecurityQuotesRequest)
	obj.reply = new(SecurityQuotesReply)

	obj.reqHeader.ZipFlag = proto.FlagNotZipped
	obj.reqHeader.SeqID = internal.SeqID()
	obj.reqHeader.PacketType = 0x01
	obj.reqHeader.Method = proto.STD_MSG_SECURITY_QUOTES_old
	obj.contentHex = "0500000000000000" // 1.3.5以前的版本
	obj.mapCode = make(map[string]Stock)
	return obj
}

func (obj *SecurityQuotesPackage) SetParams(req *SecurityQuotesRequest) {
	req.Count = uint16(len(req.StockList))
	obj.request = req
	for i := 0; i < len(req.StockList); i++ {
		v := req.StockList[i]
		securityCode := proto.GetMarketFlag(v.Market) + v.Code
		obj.mapCode[securityCode] = v
	}
}

func (obj *SecurityQuotesPackage) Serialize() ([]byte, error) {
	obj.reqHeader.PkgLen1 = 2 + uint16(len(obj.request.StockList)*7) + 10
	obj.reqHeader.PkgLen2 = 2 + uint16(len(obj.request.StockList)*7) + 10

	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, obj.reqHeader)
	b, err := hex.DecodeString(obj.contentHex)
	buf.Write(b)

	err = binary.Write(buf, binary.LittleEndian, obj.request.Count)
	for _, stock := range obj.request.StockList {
		code := make([]byte, 6)
		copy(code, stock.Code)
		tmp := []byte{stock.Market}
		tmp = append(tmp, code...)
		buf.Write(tmp)
	}

	return buf.Bytes(), err
}

func (obj *SecurityQuotesPackage) UnSerialize(header interface{}, data []byte) error {
	obj.respHeader = header.(*StdResponseHeader)

	pos := 0
	var _tmp uint16
	_ = binary.Read(bytes.NewBuffer(data[pos:pos+2]), binary.LittleEndian, &_tmp)
	pos += 2 // 跳过两个字节
	_ = binary.Read(bytes.NewBuffer(data[pos:pos+2]), binary.LittleEndian, &obj.reply.Count)
	pos += 2
	for index := uint16(0); index < obj.reply.Count; index++ {
		ele := SecurityQuote{}
		_ = binary.Read(bytes.NewBuffer(data[pos:pos+1]), binary.LittleEndian, &ele.Market)
		pos += 1
		var code [6]byte
		_ = binary.Read(bytes.NewBuffer(data[pos:pos+6]), binary.LittleEndian, &code)
		//enc := mahonia.NewDecoder("gbk")
		//ele.Code = enc.ConvertString(string(code[:]))
		ele.Code = internal.Utf8ToGbk(code[:])
		pos += 6
		_ = binary.Read(bytes.NewBuffer(data[pos:pos+2]), binary.LittleEndian, &ele.Active1)
		pos += 2

		price := internal.DecodeVarint(data, &pos)
		ele.Price = obj.getPrice(price, 0)
		ele.LastClose = obj.getPrice(price, internal.DecodeVarint(data, &pos))
		ele.Open = obj.getPrice(price, internal.DecodeVarint(data, &pos))
		ele.High = obj.getPrice(price, internal.DecodeVarint(data, &pos))
		ele.Low = obj.getPrice(price, internal.DecodeVarint(data, &pos))

		ele.ReversedBytes0 = internal.DecodeVarint(data, &pos)
		if ele.ReversedBytes0 > 0 {
			//ele.ServerTime = timeFromStr(fmt.Sprintf("%d", ele.ReversedBytes0))
			ele.ServerTime = internal.TimeFromInt(ele.ReversedBytes0)
		} else {
			ele.ServerTime = "0"
			// 如果出现这种情况, 可能是退市或者其实交易状态异常的数据, 摘牌的情况下, 证券代码是错的
			//ele.Code = proto.StockDelisting
			// 证券代码可能部证券, 上海交易所的退市代码有机会填写成600839
		}
		ele.ReversedBytes1 = internal.DecodeVarint(data, &pos)

		ele.Vol = internal.DecodeVarint(data, &pos)
		ele.CurVol = internal.DecodeVarint(data, &pos)

		var amountraw uint32
		_ = binary.Read(bytes.NewBuffer(data[pos:pos+4]), binary.LittleEndian, &amountraw)
		pos += 4
		ele.Amount = internal.IntToFloat64(int(amountraw))

		ele.SVol = internal.DecodeVarint(data, &pos)
		ele.BVol = internal.DecodeVarint(data, &pos)

		// 开盘金额需要 * 100
		ele.IndexOpenAmount = internal.DecodeVarint(data, &pos) * 100
		ele.StockOpenAmount = internal.DecodeVarint(data, &pos) * 100
		// 确定当前数据是指数或者板块
		isIndexOrBlock := proto.AssertIndexByMarketAndCode(ele.Market, ele.Code)
		tmpOpenVolume := float64(0)
		if isIndexOrBlock {
			// 指数或者板块, 单位是"股"
			tmpOpenVolume = math.Round(float64(ele.IndexOpenAmount) / ele.Open)
		} else {
			// 个股, 单位是"股"
			tmpOpenVolume = math.Round(float64(ele.StockOpenAmount) / ele.Open)
		}
		if internal.Float64IsNaN(tmpOpenVolume) {
			tmpOpenVolume = 0.00
		}
		ele.OpenVolume = int(tmpOpenVolume)
		var bidLevels []Level
		var askLevels []Level
		for i := 0; i < 5; i++ {
			bidele := Level{Price: obj.getPrice(internal.DecodeVarint(data, &pos), price)}
			offerele := Level{Price: obj.getPrice(internal.DecodeVarint(data, &pos), price)}
			bidele.Vol = internal.DecodeVarint(data, &pos)
			offerele.Vol = internal.DecodeVarint(data, &pos)
			bidLevels = append(bidLevels, bidele)
			askLevels = append(askLevels, offerele)
		}
		ele.Bid1 = bidLevels[0].Price
		ele.Bid2 = bidLevels[1].Price
		ele.Bid3 = bidLevels[2].Price
		ele.Bid4 = bidLevels[3].Price
		ele.Bid5 = bidLevels[4].Price
		ele.Ask1 = askLevels[0].Price
		ele.Ask2 = askLevels[1].Price
		ele.Ask3 = askLevels[2].Price
		ele.Ask4 = askLevels[3].Price
		ele.Ask5 = askLevels[4].Price

		ele.BidVol1 = bidLevels[0].Vol
		ele.BidVol2 = bidLevels[1].Vol
		ele.BidVol3 = bidLevels[2].Vol
		ele.BidVol4 = bidLevels[3].Vol
		ele.BidVol5 = bidLevels[4].Vol

		ele.AskVol1 = askLevels[0].Vol
		ele.AskVol2 = askLevels[1].Vol
		ele.AskVol3 = askLevels[2].Vol
		ele.AskVol4 = askLevels[3].Vol
		ele.AskVol5 = askLevels[4].Vol
		//fmt.Printf("pos: %d\n", pos)
		//fmt.Println(hex.EncodeToString(data[:pos]))

		_ = binary.Read(bytes.NewBuffer(data[pos:pos+2]), binary.LittleEndian, &ele.ReversedBytes4)
		pos += 2
		ele.ReversedBytes5 = internal.DecodeVarint(data, &pos)
		ele.ReversedBytes6 = internal.DecodeVarint(data, &pos)
		ele.ReversedBytes7 = internal.DecodeVarint(data, &pos)
		ele.ReversedBytes8 = internal.DecodeVarint(data, &pos)

		var reversedbytes9 int16
		_ = binary.Read(bytes.NewBuffer(data[pos:pos+2]), binary.LittleEndian, &reversedbytes9)
		pos += 2
		ele.Rate = float64(reversedbytes9) / 100.0
		_ = binary.Read(bytes.NewBuffer(data[pos:pos+2]), binary.LittleEndian, &ele.Active2)
		pos += 2

		// 交易状态判断
		if ele.LastClose == float64(0) && ele.Open == float64(0) {
			// 设置为退市状态
			ele.State = TDX_SECURITY_TRADE_STATE_DELISTING
		} else {
			// 如果不是退市状态, 从临时映射中删除
			securityCode := proto.GetMarketFlag(ele.Market) + ele.Code
			delete(obj.mapCode, securityCode)
			// 如果开盘价非0, 交易状态正常
			if ele.Open != float64(0) {
				ele.State = TDX_SECURITY_TRADE_STATE_NORMAL
			} else {
				// 开盘价等于0, 停牌
				ele.State = TDX_SECURITY_TRADE_STATE_SUSPEND
			}
		}

		if isIndexOrBlock {
			ele.IndexUp = ele.BidVol1
			ele.IndexDown = ele.AskVol1
			ele.IndexUpLimit = ele.BidVol2
			ele.IndexUpLimit = ele.AskVol2
		}
		upDateInRealTime, status := trading.CanUpdateInRealtime()
		if !upDateInRealTime && status == trading.ExchangeClosing {
			// 收盘
			if isIndexOrBlock {
				ele.CloseVolume = int(float64(ele.CurVol*100) / ele.Price)
			} else {
				ele.CloseVolume = ele.CurVol * 100
			}
		}
		obj.reply.List = append(obj.reply.List, ele)
	}
	// 修正停牌的证券代码
	for i := 0; len(obj.mapCode) > 0 && i < len(obj.reply.List); i++ {
		v := &(obj.reply.List[i])
		if v.State == TDX_SECURITY_TRADE_STATE_DELISTING {
			securityCode := proto.GetMarketFlag(v.Market) + v.Code
			if _, ok := obj.mapCode[securityCode]; ok {
				// 代码正常
				delete(obj.mapCode, securityCode)
			} else {
				for key, value := range obj.mapCode {
					if value.Market == v.Market {
						securityCode = key
						v.Code = value.Code
						break
					}
				}
				delete(obj.mapCode, securityCode)
			}
		}
	}
	return nil
}

func (obj *SecurityQuotesPackage) Reply() interface{} {
	return obj.reply
}

func (obj *SecurityQuotesPackage) getPrice(price int, diff int) float64 {
	return float64(price+diff) / 100.0
}
