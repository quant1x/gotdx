package internal

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeFromInt(t *testing.T) {
	nTime := 14986367
	nTime = 14986967
	nTime = 11026532
	nTime = 11295421
	nTime = 10100682
	nTime = 8836243
	nTime = 150006364
	t1 := time.UnixMilli(int64(nTime))
	fmt.Println(t1)
	//nTime = 8
	s := TimeFromInt(nTime)
	fmt.Println(s)

}
