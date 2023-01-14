package cstruct

func Unpack(buf []byte, obj IStruct) error {
	return NewBuffer(buf).Unmarshal(obj)
}
