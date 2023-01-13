package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBytes2HexString(t *testing.T) {
	str := Bytes2HexString([]byte{0, 1, 1, 5, 4, 16, 255})
	assert.Equal(t, "000101050410ff", str)
}

func TestHexString2Bytes(t *testing.T) {
	hexStr := "0c 02 18 93 00 01 03 00 03 00 0d 00 01"
	b := HexString2Bytes(hexStr)
	t.Log(b)
}
