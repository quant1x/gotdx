package proto

import (
	"fmt"
	"testing"
)

func TestGetSecurityCode(t *testing.T) {
	fmt.Println(GetSecurityCode(1, "600600"))
	fmt.Println(GetSecurityCode(0, "399001"))
	fmt.Println(GetSecurityCode(2, "399001"))
}
