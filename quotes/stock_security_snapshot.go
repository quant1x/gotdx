package quotes

// Snapshot L1 行情快照
type Snapshot struct {
	Date            string     // 交易日期
	State           TradeState // 交易状态
	Market          uint8      // 市场
	Code            string     // 代码
	Active          uint16     // 活跃度
	Price           float64    // 现价
	LastClose       float64    // 昨收
	Open            float64    // 开盘
	High            float64    // 最高
	Low             float64    // 最低
	ServerTime      string     // 时间
	ReversedBytes0  int        // 保留(时间 ServerTime)
	ReversedBytes1  int        // 保留
	Vol             int        // 总量
	CurVol          int        // 现量
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
