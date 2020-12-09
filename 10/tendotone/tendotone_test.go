package tendotone

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	type args struct {
		a      []int
		b      []int
		length int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1-test", args{[]int{1, 3, 5, 7, 9, 0, 0, 0, 0, 0}, []int{2, 4, 6, 8, 10}, 5}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{"2-test", args{[]int{2, 4, 6, 8, 10, 0, 0, 0, 0, 0}, []int{1, 3, 5, 7, 9}, 5}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge(tt.args.a, tt.args.b, tt.args.length); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
