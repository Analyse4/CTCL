package fourdottwo

import (
	"reflect"
	"testing"

	"github.com/Analyse4/mew/adt/tree"
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
		{"height-3", args{[]int{1, 2, 3, 5}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := tree.New()
			tr.ChangeRoot(MinimalHeightTree(tt.args.l))
			if got := tr.Height(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinimalHeightTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
