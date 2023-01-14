package cstruct

import "errors"

var (
	ErrNil = errors.New("cstruct: Pack called with nil")
)

type IStruct interface {
}

// slice 元素类型为指针时，是否忽略nil
var OptionSliceIgnoreNil = false
