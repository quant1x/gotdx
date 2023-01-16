package quotes

import "errors"

type TdxMarket int

const (
	DefaultRetryTimes  = 3 // 重试次数
	MessageHeaderBytes = 0x10
	MessageMaxBytes    = 1 << 15
)

var (
	ErrBadData = errors.New("more than 8M data")
)
