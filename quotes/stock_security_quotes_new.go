package quotes

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"gitee.com/quant1x/exchange"
	"gitee.com/quant1x/gotdx/internal"
	"gitee.com/quant1x/gotdx/proto"
	"gitee.com/quant1x/gox/logger"
)

const (
	TDX_SECURITY_QUOTES_MAX_V2 = 80 // 单次最大获取80条实时数据
)

// V2SecurityQuotesPackage 盘口五档报价
type V2SecurityQuotesPackage struct {
	reqHeader  *StdRequestHeader
	respHeader *StdResponseHeader
	request    *V2SecurityQuotesRequest
	reply      *V2SecurityQuotesReply

	contentHex string
}

type V2Stock struct {
	Market uint8
	Code   string
}

type V2SecurityQuotesRequest struct {
	Count     uint16
	StockList []V2Stock
}

type V2SecurityQuotesReply struct {
	Count uint16
	List  []V2SecurityQuote
}

type V2SecurityQuote struct {
	Market         uint8   // 市场
	Code           string  // 代码
	Active1        uint16  // 活跃度
	Price          float64 // 现价
	LastClose      float64 // 昨收
	Open           float64 // 开盘
	High           float64 // 最高
	Low            float64 // 最低
	ServerTime     string  // 时间
	ReversedBytes0 int     // 保留(时间 ServerTime)
	ReversedBytes1 int     // 保留
	Vol            int     // 总量
	CurVol         int     // 现量
	Amount         float64 // 总金额
	SVol           int     // 内盘
	BVol           int     // 外盘
	ReversedBytes2 int     // 保留
	ReversedBytes3 int     // 保留
	Bid1           float64
	Ask1           float64
	BidVol1        int
	AskVol1        int
	//Bid2           float64
	//Ask2           float64
	//BidVol2        int
	//AskVol2        int
	//Bid3           float64
	//Ask3           float64
	//BidVol3        int
	//AskVol3        int
	//Bid4           float64
	//Ask4           float64
	//BidVol4        int
	//AskVol4        int
	//Bid5           float64
	//Ask5           float64
	//BidVol5        int
	//AskVol5        int
	ReversedBytes4 uint16  // 保留
	ReversedBytes5 int     // 保留
	ReversedBytes6 int     // 保留
	ReversedBytes7 int     // 保留
	ReversedBytes8 int     // 保留
	Rate           float64 // 涨速
	Active2        uint16  // 活跃度
}

type V2Level struct {
	Price float64
	Vol   int
}

func NewV2SecurityQuotesPackage() *V2SecurityQuotesPackage {
	obj := new(V2SecurityQuotesPackage)
	obj.reqHeader = new(StdRequestHeader)
	obj.respHeader = new(StdResponseHeader)
	obj.request = new(V2SecurityQuotesRequest)
	obj.reply = new(V2SecurityQuotesReply)

	obj.reqHeader.ZipFlag = proto.FlagNotZipped
	obj.reqHeader.SeqID = internal.SequenceId()
	obj.reqHeader.PacketType = 0x01
	obj.reqHeader.Method = proto.STD_MSG_SECURITY_QUOTES_new
	obj.contentHex = "0500000000000000" // 1.3.5以前的版本
	return obj
}

func (obj *V2SecurityQuotesPackage) SetParams(req *V2SecurityQuotesRequest) {
	req.Count = uint16(len(req.StockList))
	obj.request = req
}

func (obj *V2SecurityQuotesPackage) Serialize() ([]byte, error) {
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

func (obj *V2SecurityQuotesPackage) UnSerialize(header interface{}, data []byte) error {
	obj.respHeader = header.(*StdResponseHeader)

	if logger.IsDebug() {
		logger.Debugf(hex.EncodeToString(data))
	}
	pos := 0

	pos += 2 // 跳过两个字节
	_ = binary.Read(bytes.NewBuffer(data[pos:pos+2]), binary.LittleEndian, &obj.reply.Count)
	pos += 2
	for index := uint16(0); index < obj.reply.Count; index++ {
		ele := V2SecurityQuote{}
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
			ele.Code = exchange.StockDelisting
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

		ele.ReversedBytes2 = internal.DecodeVarint(data, &pos)
		ele.ReversedBytes3 = internal.DecodeVarint(data, &pos)
		//fmt.Printf("pos: %d\n", pos)
		//fmt.Println(hex.EncodeToString(data[:pos]))

		var bidLevels []V2Level
		var askLevels []V2Level
		//baNum := 5
		baNum := 1
		for i := 0; i < baNum; i++ {
			bidele := V2Level{Price: obj.getPrice(internal.DecodeVarint(data, &pos), price)}
			offerele := V2Level{Price: obj.getPrice(internal.DecodeVarint(data, &pos), price)}
			bidele.Vol = internal.DecodeVarint(data, &pos)
			offerele.Vol = internal.DecodeVarint(data, &pos)
			bidLevels = append(bidLevels, bidele)
			askLevels = append(askLevels, offerele)
		}
		ele.Bid1 = bidLevels[0].Price
		//ele.Bid2 = bidLevels[1].Price
		//ele.Bid3 = bidLevels[2].Price
		//ele.Bid4 = bidLevels[3].Price
		//ele.Bid5 = bidLevels[4].Price
		ele.Ask1 = askLevels[0].Price
		//ele.Ask2 = askLevels[1].Price
		//ele.Ask3 = askLevels[2].Price
		//ele.Ask4 = askLevels[3].Price
		//ele.Ask5 = askLevels[4].Price

		ele.BidVol1 = bidLevels[0].Vol
		//ele.BidVol2 = bidLevels[1].Vol
		//ele.BidVol3 = bidLevels[2].Vol
		//ele.BidVol4 = bidLevels[3].Vol
		//ele.BidVol5 = bidLevels[4].Vol

		ele.AskVol1 = askLevels[0].Vol
		//ele.AskVol2 = askLevels[1].Vol
		//ele.AskVol3 = askLevels[2].Vol
		//ele.AskVol4 = askLevels[3].Vol
		//ele.AskVol5 = askLevels[4].Vol
		//fmt.Printf("pos: %d\n", pos)
		//fmt.Println(hex.EncodeToString(data[:pos]))

		_ = binary.Read(bytes.NewBuffer(data[pos:pos+2]), binary.LittleEndian, &ele.ReversedBytes4)
		pos += 2
		//ele.ReversedBytes5 = getPrice(data, &pos)
		//ele.ReversedBytes6 = getPrice(data, &pos)
		//ele.ReversedBytes7 = getPrice(data, &pos)
		//ele.ReversedBytes8 = getPrice(data, &pos)

		var reversedbytes9 int16
		_ = binary.Read(bytes.NewBuffer(data[pos:pos+2]), binary.LittleEndian, &reversedbytes9)
		pos += 2
		ele.Rate = float64(reversedbytes9) / 100.0

		// 保留 2个字节
		//_r1 := data[pos : pos+2]
		//_pos := 0
		//_price1 := getPrice(_r1, &_pos)
		pos += 2
		//pos += _pos

		// 保留 12x4字节
		_lenth := 12 * 4
		_r2 := data[pos : pos+_lenth]
		_pos2 := 0
		for {
			_p2 := obj.getPrice(internal.DecodeVarint(_r2, &_pos2), price)
			//_p2 := getPrice(_r2, &_pos2)
			if logger.IsDebug() {
				logger.Debug(_p2)
			}
			if _pos2 >= _lenth {
				break
			}
		}
		//_ = _price
		pos += _lenth

		_ = binary.Read(bytes.NewBuffer(data[pos:pos+2]), binary.LittleEndian, &ele.Active2)
		pos += 2

		obj.reply.List = append(obj.reply.List, ele)
	}
	return nil
}

func (obj *V2SecurityQuotesPackage) Reply() interface{} {
	return obj.reply
}

func (obj *V2SecurityQuotesPackage) getPrice(price int, diff int) float64 {
	return float64(price+diff) / 100.0
}
