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
	start := "2023-04-12"
	end := "2023-04-13"
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
