package eightdottwelve

import (
	"reflect"
	"testing"
)

// func TestEightQueen(t *testing.T) {
// 	type args struct {
// 		n        int
// 		allForms [][]*Position
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want int
// 	}{
// 		{"8-test", args{8, make([][]*Position, 0)}, 92},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := EightQueen(tt.args.n, tt.args.allForms); !reflect.DeepEqual(len(got), tt.want) {
// 				t.Errorf("EightQueen() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestTslice(t *testing.T) {
// 	l := make([]int, 0)
// 	Tslice(l)
// 	if len(l) != 1 {
// 		t.Fatalf("error, got: %d\n", len(l))
// 	}
// }

func TestEightQueenV2(t *testing.T) {
	type args struct {
		n       int
		poss    []int
		allPoss [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"8-test", args{0, make([]int, 8), make([][]int, 0)}, 92},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EightQueenV2(tt.args.n, tt.args.poss, tt.args.allPoss); !reflect.DeepEqual(len(got), tt.want) {
				t.Errorf("EightQueenV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
