package internal

import "sync/atomic"

// 局部变量
var (
	// 序列号
	_seqId atomic.Uint32
)

// SequenceId 生成序列号
func SequenceId() uint32 {
	return _seqId.Add(1)
}
