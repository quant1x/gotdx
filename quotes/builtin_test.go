package quotes

import (
	"fmt"
	"testing"
)

func Test__format_time(t *testing.T) {
	fmt.Println(time_from_str("073382"))
	fmt.Println(time_from_str("14989631"))

	fmt.Println(time_from_int(73382))
	fmt.Println(time_from_int(14989631))

}
