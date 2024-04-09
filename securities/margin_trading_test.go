package securities

import (
	"fmt"
	"testing"
)

func TestMarginTradingList(t *testing.T) {
	v1 := MarginTradingList()
	fmt.Println(v1)
	v2 := MarginTradingList()
	fmt.Println(v2)
}

func TestIsMarginTradingTarget(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "600178",
			args: args{code: "600178"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsMarginTradingTarget(tt.args.code); got != tt.want {
				t.Errorf("IsMarginTradingTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
