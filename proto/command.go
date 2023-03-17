package proto

// 标准行情

// 标准行情-命令字
const (
	STD_MSG_HEARTBEAT                = 0x0004 // 心跳维持
	STD_MSG_LOGIN1                   = 0x000d // 第一次登录
	STD_MSG_LOGIN2                   = 0x0fdb // 第二次登录
	STD_MSG_XDXR_INFO                = 0x000f // 除权除息信息
	STD_MSG_FINANCE_INFO             = 0x0010 // 财务信息
	STD_MSG_PING                     = 0x0015 // 测试连接
	STD_MSG_COMPANY_CATEGORY         = 0x02cf // 公司信息文件信息
	STD_MSG_COMPANY_CONTENT          = 0x02d0 // 公司信息描述
	STD_MSG_SECURITY_COUNT           = 0x044e // 证券数量
	STD_MSG_SECURITY_LIST            = 0x0450 // 证券列表
	STD_MSG_INDEXBARS                = 0x052d // 指数K线
	STD_MSG_SECURITY_BARS            = 0x052d // 股票K线
	STD_MSG_SECURITY_QUOTES_old      = 0x053e // 行情信息
	STD_MSG_SECURITY_QUOTES_new      = 0x054c // 行情信息
	STD_MSG_MINUTETIME_DATA          = 0x051d // 分时数据
	STD_MSG_BLOCK_META               = 0x02c5 // 板块文件信息
	STD_MSG_BLOCK_DATA               = 0x06b9 // 板块文件数据
	STD_MSG_TRANSACTION_DATA         = 0x0fc5 // 分笔成交信息
	STD_MSG_HISTORY_MINUTETIME_DATA  = 0x0fb4 // 历史分时信息
	STD_MSG_HISTORY_TRANSACTION_DATA = 0x0fb5 // 历史分笔成交信息
)

// K线种类
const (
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
	Compressed    = uint8(0x10)                       // 压缩标志
	FlagNotZipped = uint8(0x0c)                       // zip未压缩
	FlagZipped    = uint8(Compressed | FlagNotZipped) // zip已压缩 消息头标志 0x789C
)
