package util

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
	"io"
	"strings"
)

// Utf8ToGbk utf8 转gbk
func Utf8ToGbk(text []byte) string {
	r := bytes.NewReader(text)
	decoder := transform.NewReader(r, simplifiedchinese.GBK.NewDecoder()) //GB18030
	content, _ := io.ReadAll(decoder)
	return strings.ReplaceAll(string(content), string([]byte{0x00}), "")
}

// DecodeGBK convert GBK to UTF-8
func DecodeGBK(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, simplifiedchinese.GBK.NewDecoder())
	d, e := io.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}

// EncodeGBK convert UTF-8 to GBK
func EncodeGBK(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, simplifiedchinese.GBK.NewEncoder())
	d, e := io.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}

// DecodeBig5 convert BIG5 to UTF-8
func DecodeBig5(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, traditionalchinese.Big5.NewDecoder())
	d, e := io.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}

// EncodeBig5 convert UTF-8 to BIG5
func EncodeBig5(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, traditionalchinese.Big5.NewEncoder())
	d, e := io.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}
