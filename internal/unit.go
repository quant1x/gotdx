package internal

// BaseUnit 交易单位
func BaseUnit(code string) float64 {
	c := code[:2]
	switch c {
	case "60", "68", "00", "30", "39":
		return 100.0
	}
	c = code[:3]
	switch c {
	case "510":
		return 1000.0
	}
	return 100.00
}
