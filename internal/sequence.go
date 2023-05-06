package internal

import "sync/atomic"

// 局部变量
var (
	// 序列号
	_seqId uint32
)

// SeqID 生成序列号
func SeqID() uint32 {
	atomic.AddUint32(&_seqId, 1)
	return _seqId
}
