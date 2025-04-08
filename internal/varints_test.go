package internal

import "testing"

func TestDecodeVarint(t *testing.T) {
	pos := 0
	type args struct {
		b   []byte
		pos *int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				b:   []byte{255, 136, 15},
				pos: &pos,
			},
			want: -123456,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecodeVarint(tt.args.b, tt.args.pos); got != tt.want {
				t.Errorf("DecodeVarint() = %v, want %v", got, tt.want)
			}
		})
	}
}
