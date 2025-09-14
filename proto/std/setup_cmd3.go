package std

import (
	"github.com/quant1x/gotdx/internal"
)

// SetupCmd3Request 请求包结构
type SetupCmd3Request struct {
	Cmd []byte `struc:"[42]byte" json:"cmd"`
}

// Marshal 请求包序列化输出
func (req *SetupCmd3Request) Marshal() ([]byte, error) {
	return DefaultMarshal(req)
}

// SetupCmd3Response 响应包结构
type SetupCmd3Response struct {
	Unknown []byte `json:"unknown"`
}

func (resp *SetupCmd3Response) Unmarshal(data []byte) error {
	resp.Unknown = data
	return nil
}

// NewSetupCmd3Request 创建SetupCmd3请求包
func NewSetupCmd3Request() (*SetupCmd3Request, error) {
	request := &SetupCmd3Request{
		Cmd: internal.HexString2Bytes("0c 03 18 99 00 01 20 00 20 00 db 0f d5" +
			"d0 c9 cc d6 a4 a8 af 00 00 00 8f c2 25" +
			"40 13 00 00 d5 00 c9 cc bd f0 d7 ea 00" +
			"00 00 02"),
	}
	return request, nil
}

func NewSetupCmd3() (*SetupCmd3Request, *SetupCmd3Response, error) {
	var response SetupCmd3Response
	var request, err = NewSetupCmd3Request()
	return request, &response, err
}
