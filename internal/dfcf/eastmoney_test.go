package dfcf

import (
	"fmt"
	"testing"
)

func Test_stock_hist(t *testing.T) {
	ks, err := A("sh000001")
	if err != nil {
		_ = fmt.Errorf("error :%v", err.Error())
	}
	fmt.Println(ks)

}
