package std

import (
	"gitee.com/quant1x/gotdx/internal"
)

// SetupCmd1Request 请求包结构
type SetupCmd1Request struct {
	Cmd []byte `struc:"[13]byte" json:"cmd"`
}

// Marshal 请求包序列化输出
func (req *SetupCmd1Request) Marshal() ([]byte, error) {
	return DefaultMarshal(req)
}

// SetupCmd1Response 响应包结构
// serverInfo := Utf8ToGbk(data[68:])
type SetupCmd1Response struct {
	Unknown []byte `json:"unknown"`
	Reply   string `json:"reply"`
}

func (resp *SetupCmd1Response) Unmarshal(data []byte) error {
	//resp.Unknown = data
	resp.Reply = internal.Utf8ToGbk(data[68:])
	return nil
}

// 创建SetupCmd1请求包
func NewSetupCmd1Request() (*SetupCmd1Request, error) {
	request := &SetupCmd1Request{
		Cmd: internal.HexString2Bytes("0c 02 18 93 00 01 03 00 03 00 0d 00 01"),
	}
	return request, nil
}

func NewSetupCmd1() (*SetupCmd1Request, *SetupCmd1Response, error) {
	var response SetupCmd1Response
	var request, err = NewSetupCmd1Request()
	return request, &response, err
}
