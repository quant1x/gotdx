package functions

import (
	//"context"
	"fmt"
	"testing"
	"time"
)

func TestConsts(t *testing.T) {
	count := 10
	for i := 0; i < count; i++ {
		fmt.Println(FROMOPEN)
		time.Sleep(1 * time.Second)
	}
}
