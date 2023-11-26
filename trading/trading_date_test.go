package trading

import (
	"fmt"
	"testing"
)

func TestLastNDate(t *testing.T) {
	dates := LastNDate(Today(), 5)
	fmt.Println(dates)
}

func TestNextTradeDate(t *testing.T) {
	date := NextTradeDate("20230403")
	fmt.Println(date)
}

func TestTradeRange(t *testing.T) {
	start := "2014-08-18"
	end := "2014-08-24"
	dates := TradeRange(start, end)
	fmt.Println(len(dates))
	fmt.Println(dates)
}

func TestGetLastDayForUpdate(t *testing.T) {
	fmt.Println(GetLastDayForUpdate())
}

func TestGetFrontTradeDay(t *testing.T) {
	fmt.Println(GetFrontTradeDay())
}

func TestGetCurrentDate(t *testing.T) {
	date := "20230721"
	v := GetCurrentDate(date)
	fmt.Println(v)
}

func TestTimeRange(t *testing.T) {
	getTimeRanges()
	tr := DateTimeRange{Begin: trAMBegin, End: trAMEnd}
	fmt.Println(tr.Minutes())
}
