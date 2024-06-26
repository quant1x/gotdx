package std

import (
	"gitee.com/quant1x/gotdx/internal"
)

// SetupCmd2Request 请求包结构
type SetupCmd2Request struct {
	Cmd []byte `struc:"[13]byte" json:"cmd"`
}

// Marshal 请求包序列化输出
func (req *SetupCmd2Request) Marshal() ([]byte, error) {
	return DefaultMarshal(req)
}

// SetupCmd2Response 响应包结构
type SetupCmd2Response struct {
	Unknown []byte `json:"unknown"`
}

func (resp *SetupCmd2Response) Unmarshal(data []byte) error {
	resp.Unknown = data
	return nil
}

// NewSetupCmd2Request 创建SetupCmd2请求包
func NewSetupCmd2Request() (*SetupCmd2Request, error) {
	request := &SetupCmd2Request{
		Cmd: internal.HexString2Bytes("0c 02 18 94 00 01 03 00 03 00 0d 00 02"),
	}
	return request, nil
}

func NewSetupCmd2() (*SetupCmd2Request, *SetupCmd2Response, error) {
	var response SetupCmd2Response
	var request, err = NewSetupCmd2Request()
	return request, &response, err
}
