package std

import (
	"gitee.com/quant1x/gotdx/internal"
	"gitee.com/quant1x/gotdx/proto"
)

// SecurityCountRequest 请求包结构
type SecurityCountRequest struct {
	Unknown1 []byte           `struc:"[12]byte"`
	Market   proto.MarketType `struc:"uint16,little" json:"market"`
	Unknown2 []byte           `struc:"[4]byte"`
}

// Marshal 请求包序列化输出
func (req *SecurityCountRequest) Marshal() ([]byte, error) {
	return DefaultMarshal(req)
}

// SecurityCountResponse 响应包结构
type SecurityCountResponse struct {
	Count uint `struc:"uint16,little" json:"count"`
}

func (resp *SecurityCountResponse) Unmarshal(data []byte) error {
	return DefaultUnmarshal(data, resp)
}

// todo: 检测market是否为合法值
func NewSecurityCountRequest(market proto.MarketType) (*SecurityCountRequest, error) {
	request := &SecurityCountRequest{
		Unknown1: internal.HexString2Bytes("0c 0c 18 6c 00 01 08 00 08 00 4e 04"),
		Market:   market,
		Unknown2: internal.HexString2Bytes("75 c7 33 01"),
	}
	return request, nil
}

func NewSecurityCount(market proto.MarketType) (*SecurityCountRequest, *SecurityCountResponse, error) {
	var response SecurityCountResponse
	var request, err = NewSecurityCountRequest(market)
	return request, &response, err
}
