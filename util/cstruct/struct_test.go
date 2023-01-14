package cstruct

import (
	"fmt"
	"testing"
	"unsafe"
)

type StructA struct {
	A1 uint8
	A2 uint32
	A3 [5]uint8
}

type StructB struct {
	B1 uint8
	B2 StructA
	B3 uint16
	B4 float32
	B5 [3]StructA
}

func TestStruct(t *testing.T) {
	b := StructB{}
	b.B1 = 127
	b.B2.A1 = 56
	b.B2.A2 = 999
	b.B2.A3[0] = 0
	b.B2.A3[1] = 1
	b.B2.A3[2] = 2
	b.B2.A3[3] = 3
	b.B2.A3[4] = 4
	b.B3 = 8888
	b.B4 = 88.8
	b.B5[0] = b.B2
	b.B5[1] = b.B2
	b.B5[2] = b.B2

	data, _ := Pack(&b)

	fmt.Println("len(b) =", unsafe.Sizeof(b))
	fmt.Println("struct data len = ", len(data))
	fmt.Println("struct data is:")
	for i := 0; i < len(data); i++ {
		fmt.Printf("%d ", data[i])
	}
}
