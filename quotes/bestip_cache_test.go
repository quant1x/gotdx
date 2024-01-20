package quotes

import (
	"fmt"
	"testing"
)

func TestOpenConfig(t *testing.T) {
	list := GetFastHost(TDX_HOST_HQ)
	fmt.Printf("%+v\n", list)
}
