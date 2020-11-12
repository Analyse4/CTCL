package fivedotthree

import "testing"

func TestFlipBitToWin(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1775-8-test", args{1775}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlipBitToWin(tt.args.num); got != tt.want {
				t.Errorf("FlipBitToWin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlipBitToWinV2(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1775-8-test", args{1775}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlipBitToWinV2(tt.args.num); got != tt.want {
				t.Errorf("FlipBitToWin() = %v, want %v", got, tt.want)
			}
		})
	}
}
