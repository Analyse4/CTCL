package fivedottwo

import (
	"testing"
)

func TestBinaryToString(t *testing.T) {
	type args struct {
		bin float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"0.72-test", args{0.625}, "0.101"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinaryToString(tt.args.bin); got != tt.want {
				t.Errorf("BinaryToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
