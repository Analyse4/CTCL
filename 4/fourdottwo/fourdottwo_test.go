package fourdottwo

import (
	"reflect"
	"testing"
)

func TestMinimalHeightTree(t *testing.T) {
	type args struct {
		l []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"height-3", args{[]int{1, 2, 3, 4, 5}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := MinimalHeightTree(tt.args.l)
			if got := tr.Height(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinimalHeightTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
