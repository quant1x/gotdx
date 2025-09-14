package quotes

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"testing"

	"gitee.com/quant1x/gox/logger"
)

func parseResponseHeader(data []byte) (*StdResponseHeader, []byte, error) {
	var header StdResponseHeader
	//err := cstruct.Unpack(data, &header)
	headerBuf := bytes.NewReader(data)
	err := binary.Read(headerBuf, binary.LittleEndian, &header)
	if err != nil {
		return nil, nil, err
	}
	fmt.Println(headerBuf.Len(), headerBuf.Size())
	pos := int(headerBuf.Size()) - headerBuf.Len()
	if header.ZipSize > MessageMaxBytes {
		logger.Debugf("msgData has bytes(%d) beyond max %d\n", header.ZipSize, MessageMaxBytes)
		return &header, nil, ErrBadData
	}
	var out bytes.Buffer
	var body []byte
	if header.ZipSize != header.UnZipSize {
		b := bytes.NewReader(data[pos:])
		r, _ := zlib.NewReader(b)
		_, _ = io.Copy(&out, r)
		body = out.Bytes()
		_ = r.Close()
	} else {
		body = data[pos:]
	}
	return &header, body, err

}

func TestProcess(t *testing.T) {
	hexString := "b1cb74000c760028000004000a000a0000000000000000000000"
	data, err := hex.DecodeString(hexString)
	if err != nil {
		fmt.Println(err)
	}
	respHeader, respBody, err := parseResponseHeader(data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", respHeader)
	fmt.Printf("%+v\n", respBody)
	bodyBuff := bytes.NewReader(respBody)
	var resp HeartBeatReply
	err = binary.Read(bodyBuff, binary.LittleEndian, &resp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", resp)
}
