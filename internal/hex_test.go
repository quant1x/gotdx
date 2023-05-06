package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBytes2HexString(t *testing.T) {
	str := Bytes2HexString([]byte{0, 1, 1, 5, 4, 16, 255})
	t.Log(str)
	assert.Equal(t, "00 01 01 05 04 10 ff", str)
}

func TestHexString2Bytes(t *testing.T) {
	hexStr := "0c 02 18 93 00 01 03 00 03 00 0d 00 01"
	b := HexString2Bytes(hexStr)
	t.Log(b)
}
