package quotes

import (
	"fmt"
	"testing"
)

func TestOpenConfig(t *testing.T) {
	list := GetFastHost(HOST_HQ)
	fmt.Printf("%+v\n", list)
}
