package internal

import (
	"fmt"
	"testing"
)

func TestSeqID(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(SequenceId())
	}
}
