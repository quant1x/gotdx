package gotdx

import "errors"

const (
	// 市场
	MARKET_SZ = 0 // 深圳交易所
	MARKET_SH = 1 // 上海交易所
	MARKET_BJ = 2 // 北京交易所
)

const (
	// K线种类

	KLINE_TYPE_5MIN      = 0  //  5 分钟 K线
	KLINE_TYPE_15MIN     = 1  // 15 分钟 K线
	KLINE_TYPE_30MIN     = 2  // 30 分钟 K线
	KLINE_TYPE_1HOUR     = 3  //  1 小时 K线
	KLINE_TYPE_DAILY     = 4  //      日 K线
	KLINE_TYPE_WEEKLY    = 5  // 周 K线
	KLINE_TYPE_MONTHLY   = 6  // 月 K线
	KLINE_TYPE_EXHQ_1MIN = 7  // 1分钟
	KLINE_TYPE_1MIN      = 8  // 1 分钟 K线
	KLINE_TYPE_RI_K      = 9  // 日 K线
	KLINE_TYPE_3MONTH    = 10 // 季 K线
	KLINE_TYPE_YEARLY    = 11 // 年 K线
)

const (
	DefaultRetryTimes = 3 // 重试次数
)

var (
	ErrBadData = errors.New("more than 8M data")
)
