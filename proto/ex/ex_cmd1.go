package ex

import (
	"gitee.com/quant1x/gotdx/proto/std"
	"gitee.com/quant1x/gotdx/util"
)

// ExCmd1Request 请求包结构
type ExCmd1Request struct {
	Cmd []byte `struc:"[92]byte" json:"cmd"`
}

// Marshal 请求包序列化输出
func (req *ExCmd1Request) Marshal() ([]byte, error) {
	return std.DefaultMarshal(req)
}

// ExCmd1Response 响应包结构
type ExCmd1Response struct {
	Unknown []byte `json:"unknown"`
	Reply   string `json:"reply"`
}

func (resp *ExCmd1Response) Unmarshal(data []byte) error {
	//resp.Unknown = data
	resp.Reply = util.Utf8ToGbk(data[3:53])
	return nil
}

// NewExCmd1Request 创建ExCmd1请求包
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
