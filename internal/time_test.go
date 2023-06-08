package internal

import (
	"fmt"
	"testing"
)

func TestTimeFromInt(t *testing.T) {
	nTime := 14986367
	nTime = 14986967
	s := TimeFromInt(nTime)
	fmt.Println(s)

}
