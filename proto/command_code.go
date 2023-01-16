package proto

// 标准行情

// 标准行情-命令字
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
