package v2

import (
	v1 "gitee.com/quant1x/gotdx/proto/v1"
	"gitee.com/quant1x/gotdx/util"
)

// 请求包结构
type ExCmd1Request struct {
	Cmd []byte `struc:"[92]byte" json:"cmd"`
}

// 请求包序列化输出
func (req *ExCmd1Request) Marshal() ([]byte, error) {
	return v1.DefaultMarshal(req)
}

// 响应包结构
// serverInfo := Utf8ToGbk(data[68:])
type ExCmd1Response struct {
	Unknown []byte `json:"unknown"`
	Reply   string `json:"reply"`
}

func (resp *ExCmd1Response) Unmarshal(data []byte) error {
	//resp.Unknown = data
	resp.Reply = v1.Utf8ToGbk(data[3:53])
	return nil
}

// 创建ExCmd1请求包
func NewExCmd1Request() (*ExCmd1Request, error) {
	request := &ExCmd1Request{
		Cmd: util.HexString2Bytes("01 01 48 65 00 01 52 00 52 00 54 24 1f 32 c6 e5 d5 3d fb 41 1f 32 c6 e5 d5 3d fb 41 1f 32 c6 e5 d5 3d fb 41 1f 32 c6 e5 d5 3d fb 41 1f 32 c6 e5 d5 3d fb 41 1f 32 c6 e5 d5 3d fb 41 1f 32 c6 e5 d5 3d fb 41 1f 32 c6 e5 d5 3d fb 41 cc e1 6d ff d5 ba 3f b8 cb c5 7a 05 4f 77 48 ea"),
	}
	return request, nil
}

func NewExCmd1() (*ExCmd1Request, *ExCmd1Response, error) {
	var response ExCmd1Response
	var request, err = NewExCmd1Request()
	return request, &response, err
}
