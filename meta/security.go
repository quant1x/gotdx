package meta

type Security struct {
	Code         string  // 代码
	Name         string  // 名称
	VolUnit      uint16  // 每手股数
	DecimalPoint int8    // 价格单位指数,10的几次方
	PreClose     float64 // 昨日收盘
}
