package v1

import "errors"

var (
	ErrBadData = errors.New("more than 8M data")
)

const (
	DefaultRetryTimes  = 3 // 重试次数
	MessageHeaderBytes = 0x10
	MessageMaxBytes    = 1 << 15
)

// 命令字
const (
	KMSG_CMD1                   = 0x000d // 建立链接
	KMSG_CMD2                   = 0x0fdb // 建立链接
	KMSG_PING                   = 0x0015 // 测试连接
	KMSG_HEARTBEAT              = 0xFFFF // 心跳(自定义)
	KMSG_HEARTBEAT2             = 0x0523 // 心跳
	KMSG_SECURITYCOUNT          = 0x044e // 证券数量
	KMSG_BLOCKINFOMETA          = 0x02c5 // 板块文件信息
	KMSG_BLOCKINFO              = 0x06b9 // 板块文件
	KMSG_COMPANYCATEGORY        = 0x02cf // 公司信息文件信息
	KMSG_COMPANYCONTENT         = 0x02d0 // 公司信息描述
	KMSG_FINANCEINFO            = 0x0010 // 财务信息
	KMSG_HISTORYMINUTETIMEDATE  = 0x0fb4 // 历史分时信息
	KMSG_HISTORYTRANSACTIONDATA = 0x0fb5 // 历史分笔成交信息
	KMSG_INDEXBARS              = 0x052d // 指数K线
	KMSG_SECURITYBARS           = 0x052d // 股票K线
	KMSG_MINUTETIMEDATA         = 0x0537 // 分时数据
	KMSG_SECURITYLIST           = 0x0450 // 证券列表
	KMSG_SECURITYQUOTES         = 0x053e // 行情信息
	KMSG_TRANSACTIONDATA        = 0x0fc5 // 分笔成交信息
	KMSG_XDXRINFO               = 0x000f // 除权除息信息

)

// 市场
const (
	MARKET_SZ = 0 // 深圳交易所
	MARKET_SH = 1 // 上海交易所
	MARKET_BJ = 2 // 北京交易所
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

/*
0c 02000000 00 1c00 1c00 2d05 0100363030303030080001000000140000000000000000000000
0c 02189300 01 0300 0300 0d00 01
0c 00000000 00 0200 0200 1500
*/
type RequestHeader struct {
	Zip        uint8  // ZipFlag
	SeqID      uint32 // 请求编号
	PacketType uint8
	PkgLen1    uint16
	PkgLen2    uint16
	Method     uint16 // method 请求方法
}

type ResponseHeader struct {
	I1        uint32
	I2        uint8
	SeqID     uint32 // 请求编号
	I3        uint8
	Method    uint16 // method
	ZipSize   uint16 // 长度
	UnZipSize uint16 // 未压缩长度
}
