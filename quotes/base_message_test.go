package quotes

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"testing"
)

func parseResponse(data []byte) (*StdResponseHeader, []byte, error) {
	var header StdResponseHeader
	//err := cstruct.Unpack(data, &header)
	headerBuf := bytes.NewReader(data)
	err := binary.Read(headerBuf, binary.LittleEndian, &header)
	if err != nil {
		return nil, nil, err
	}
	fmt.Println(headerBuf.Len(), headerBuf.Size())
	pos := int(headerBuf.Size()) - headerBuf.Len()
	return &header, data[pos:], err

}

func TestProcess(t *testing.T) {
	hexString := "b1cb74000c760028000004000a000a0000000000000000000000"
	data, err := hex.DecodeString(hexString)
	if err != nil {
		fmt.Println(err)
	}
	respHeader, respBody, err := parseResponse(data)
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
