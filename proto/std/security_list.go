package std

// 获取股票列表
import (
	"gitee.com/quant1x/exchange"
	"gitee.com/quant1x/gotdx/internal"
)

// GetSecurityListRequest 请求包结构
type GetSecurityListRequest struct {
	// struc不允许slice解析，只允许包含长度的array，该长度可根据hex字符串计算
	Unknown1 []byte `struc:"[12]byte"`
	// pytdx中使用struct.Pack进行反序列化
	// 其中<H等价于这里的struc:"uint16,little"
	// <I等价于struc:"uint32,little"
	Market exchange.MarketType `struc:"uint16,little" json:"market"`
	Start  int                 `struc:"uint16,little" json:"start"`
}

// 请求包序列化输出
func (req *GetSecurityListRequest) Marshal() ([]byte, error) {
	return DefaultMarshal(req)
}

// 响应包结构
type getSecurityListResponseRaw struct {
	Count     int        `struc:"uint16,little,sizeof=StocksRaw" json:"count"`
	StocksRaw []stockRaw `struc:"[29]byte, little" json:"stocks_raw"`
}
type stockRaw struct {
	Code           string `struc:"[6]byte,little" json:"code"`
	VolUnit        int    `struc:"uint16,little" json:"vol_unit"`
	Name           []byte `struc:"[8]byte,little" json:"name"`
	ReversedBytes1 []byte `struc:"[4]byte,little" json:"reversed_bytes_1"`
	DecimalPoint   int    `struc:"byte,little" json:"decimal_point"`
	PreCloseRaw    int    `struc:"uint32,little" json:"pre_close_raw"`
	ReversedBytes2 []byte `struc:"[4]byte,little" json:"reversed_bytes_2"`
}

func (resp *getSecurityListResponseRaw) Stocks() ([]Stock, error) {
	var (
		stocks []Stock
	)
	// 后续处理
	for idx := range resp.StocksRaw {
		name, err := internal.DecodeGBK(resp.StocksRaw[idx].Name) // .rstrip("\x00")
		if err != nil {
			return nil, err
		}
		//log.Println(resp.StocksRaw[idx].PreCloseRaw)
		stocks = append(stocks, Stock{
			Code:         resp.StocksRaw[idx].Code,
			VolUnit:      resp.StocksRaw[idx].VolUnit,
			DecimalPoint: resp.StocksRaw[idx].DecimalPoint,
			Name:         string(name),
			PreClose:     internal.IntToFloat64(resp.StocksRaw[idx].PreCloseRaw),
		})
	}
	return stocks, nil
}
func (resp *getSecurityListResponseRaw) Unmarshal(data []byte) error {
	return DefaultUnmarshal(data, &resp)
}

type Stock struct {
	Code         string  `json:"code"`
	VolUnit      int     `json:"vol_unit"`
	DecimalPoint int     `json:"decimal_point"`
	Name         string  `json:"name"`
	PreClose     float64 `json:"pre_close"`
}

// GetSecurityListResponse 响应包结构
type GetSecurityListResponse struct {
	Count  int     `struc:"uint16,little,sizeof=Stocks" json:"count"`
	Stocks []Stock `struc:"[29]byte, little" json:"stocks"`
}

// Unmarshal 内部套用原始结构解析，外部为经过解析之后的响应信息
func (resp *GetSecurityListResponse) Unmarshal(data []byte) error {
	var raw getSecurityListResponseRaw
	err := raw.Unmarshal(data)
	if err != nil {
		return err
	}
	stocks, err := raw.Stocks()
	if err != nil {
		return err
	}
	resp.Stocks = stocks
	return nil
}

// todo: 检测market是否为合法值
func NewGetSecurityListRequest(market exchange.MarketType, start int) (*GetSecurityListRequest, error) {
	request := &GetSecurityListRequest{
		Unknown1: internal.HexString2Bytes("0c 01 18 64 01 01 06 00 06 00 50 04"),
		Market:   market,
		Start:    start,
	}
	return request, nil
}

func NewGetSecurityList(market exchange.MarketType, start int) (*GetSecurityListRequest, *GetSecurityListResponse, error) {
	var response GetSecurityListResponse
	var request, err = NewGetSecurityListRequest(market, start)
	return request, &response, err
}
