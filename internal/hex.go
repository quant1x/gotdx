package internal

import (
	"encoding/hex"
	"fmt"
	"log"
	"strings"
)

// HexString2Bytes 16进制字符串转bytes
func HexString2Bytes(hexStr string) []byte {
	hexStr = strings.Replace(hexStr, " ", "", -1)
	data, err := hex.DecodeString(hexStr)
	if err != nil {
		// handle error
		log.Println(err.Error())
		return nil
	}
	return data
}

// Bytes2HexString bytes转16进制字符串
func Bytes2HexString(b []byte) string {
	// with "%x" format byte array into hex string
	return fmt.Sprintf("% x", b)
}
