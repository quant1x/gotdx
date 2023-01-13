package parse

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
	"io"
)

// convert GBK to UTF-8
func DecodeGBK(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, simplifiedchinese.GBK.NewDecoder())
	d, e := io.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}

// convert UTF-8 to GBK
func EncodeGBK(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, simplifiedchinese.GBK.NewEncoder())
	d, e := io.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}

// convert BIG5 to UTF-8
func DecodeBig5(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, traditionalchinese.Big5.NewDecoder())
	d, e := io.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}

// convert UTF-8 to BIG5
func EncodeBig5(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, traditionalchinese.Big5.NewEncoder())
	d, e := io.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}
