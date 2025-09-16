package quotes

import "gitee.com/quant1x/exchange"

type ExchangeState int8

const (
	EXCHANGE_STATE_DELISTING ExchangeState = iota - 1 // 终止上市
	EXCHANGE_STATE_CLOSING                            // 收盘
	EXCHANGE_STATE_NORMAL                             // 正常交易
	EXCHANGE_STATE_PAUSE                              // 暂停交易
)

// Snapshot L1 行情快照
type Snapshot struct {
	Date            string        // 交易日期
	SecurityCode    string        // 证券代码
	ExchangeState   ExchangeState // 交易状态
	State           TradeState    // 上市公司状态
	Market          uint8         // 市场
	Code            string        // 代码
	Active          uint16        // 活跃度
	Price           float64       // 现价
	LastClose       float64       // 昨收
	Open            float64       // 开盘
	High            float64       // 最高
	Low             float64       // 最低
	ServerTime      string        // 时间
	ReversedBytes0  int           // 保留(时间 ServerTime)
	ReversedBytes1  int           // 保留
	Vol             int           // 总量
	CurVol          int           // 个股-现成交量,板块指数-现成交额
	Amount          float64       // 总金额
	SVol            int           // 个股有效-内盘
	BVol            int           // 个股有效-外盘
	IndexOpenAmount int           // 指数有效-集合竞价成交金额=开盘成交金额
	StockOpenAmount int           // 个股有效-集合竞价成交金额=开盘成交金额
	OpenVolume      int           // 集合竞价-开盘量, 单位是股
	CloseVolume     int           // 集合竞价-收盘量, 单位是股
	IndexUp         int           // 指数有效-上涨数
	IndexUpLimit    int           // 指数有效-涨停数
	IndexDown       int           // 指数有效-下跌数
	IndexDownLimit  int           // 指数有效-跌停数
	Bid1            float64       // 个股-委买价1
	Ask1            float64       // 个股-委卖价1
	BidVol1         int           // 个股-委买量1 板块-上涨数
	AskVol1         int           // 个股-委卖量1 板块-下跌数
	Bid2            float64       // 个股-委买价2
	Ask2            float64       // 个股-委卖价2
	BidVol2         int           // 个股-委买量2 板块-涨停数
	AskVol2         int           // 个股-委卖量2 板块-跌停数
	Bid3            float64       // 个股-委买价3
	Ask3            float64       // 个股-委卖价3
	BidVol3         int           // 个股-委买量3
	AskVol3         int           // 个股-委卖量3
	Bid4            float64       // 个股-委买价4
	Ask4            float64       // 个股-委卖价4
	BidVol4         int           // 个股-委买量4
	AskVol4         int           // 个股-委卖量4
	Bid5            float64       // 个股-委买价5
	Ask5            float64       // 个股-委卖价5
	BidVol5         int           // 个股-委买量5
	AskVol5         int           // 个股-委卖量5
	ReversedBytes4  uint16        // 保留
	ReversedBytes5  int           // 保留
	ReversedBytes6  int           // 保留
	ReversedBytes7  int           // 保留
	ReversedBytes8  int           // 保留
	Rate            float64       // 涨速
	Active2         uint16        // 活跃度, 如果是指数则为0, 个股同Active1
	TimeStamp       string        // 本地当前时间戳
}

// CheckDirection 检测当前交易方向
//
//	todo: 只能检测即时行情数据, 对于历史数据无效
func (this *Snapshot) CheckDirection() (biddingDirection, volumeDirection int) {
	if this.Price == this.Bid1 {
		biddingDirection = -1
	} else if this.Price == this.Ask1 {
		biddingDirection = 1
	}
	bidVol := this.BidVol1 + this.BidVol2 + this.BidVol3 + this.BidVol4 + this.BidVol5
	askVol := this.AskVol1 + this.AskVol2 + this.AskVol3 + this.AskVol4 + this.AskVol5
	volumeDirection = bidVol - askVol
	return
}

// AverageBiddingVolume 平均竞量
func (this *Snapshot) AverageBiddingVolume() int {
	bidVol := this.BidVol1 + this.BidVol2 + this.BidVol3 + this.BidVol4 + this.BidVol5
	askVol := this.AskVol1 + this.AskVol2 + this.AskVol3 + this.AskVol4 + this.AskVol5
	return (bidVol + askVol) / 10
}

// DetectBiddingPhase 检测竞价阶段
// 如果5档行情
func (this *Snapshot) DetectBiddingPhase() (head, tail bool) {
	head = false
	tail = false
	kind := exchange.AssertCode(this.SecurityCode)
	switch kind {
	case exchange.STOCK, exchange.ETF:
		// 个股竞价阶段, 竞价3-5的数据都是0
		bidPrice := int(this.Bid3 + this.Bid4 + this.Bid5)
		bidVol := this.BidVol3 + this.BidVol4 + this.BidVol5
		if bidPrice+bidVol == 0 {
			// 早盘竞价时开盘等于0
			if this.Open == 0 {
				head = true
			} else {
				tail = true
			}
		}
	case exchange.INDEX:
		// 指数
		head = this.Active == 0
		tail = this.Active > 0
	case exchange.BLOCK:
		// 板块
		head = this.Active == 0
		tail = this.Active > 0
	}

	return
}
