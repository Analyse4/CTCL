package eightdotone

import (
	"testing"
)

// func TestThreeStep(t *testing.T) {
// 	type args struct {
// 		n int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want int
// 	}{
// 		{"3-test", args{5}, 8},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := ThreeStep(tt.args.n); got != tt.want {
// 				t.Errorf("ThreeStep() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestTHreeStepV2(t *testing.T) {
// 	type args struct {
// 		n int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want int
// 	}{
// 		{"3-test", args{5}, 8},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := ThreeStepV2(tt.args.n); got != tt.want {
// 				t.Errorf("THreeStepV2() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestThreeStepV3(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"4-test", args{4}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ThreeStepV3(tt.args.n); got != tt.want {
				t.Errorf("ThreeStepV3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestThreeStepV4(t *testing.T) {
	type args struct {
		n int
		l []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"4-test", args{4, make([]int, 5)}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ThreeStepV4(tt.args.n, tt.args.l); got != tt.want {
				t.Errorf("ThreeStepV4() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestThreeStepV5(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"4-test", args{4}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ThreeStepV5(tt.args.n); got != tt.want {
				t.Errorf("ThreeStepV5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkThreeStepV3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ThreeStepV3(i)
	}
}

func BenchmarkThreeStepV4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ThreeStepV4(i, make([]int, i+1))
	}
}
