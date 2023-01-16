package quotes

import (
	"fmt"
	"testing"
)

func TestOpenConfig(t *testing.T) {
	as := OpenConfig()
	fmt.Printf("%+v\n", as)
}
