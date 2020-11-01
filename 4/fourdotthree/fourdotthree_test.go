package fourdotthree

import (
	"container/list"
	"testing"

	"github.com/Analyse4/mew/adt/tree"
)

func TestListOfDepths(t *testing.T) {
	tr1 := tree.New()
	tr1.ChangeRoot(tree.GenerateMinimalHeightTree([]int{1, 2, 3, 4, 5}))
	tr2 := tree.New()
	tr2.ChangeRoot(tree.GenerateMinimalHeightTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}))
	type args struct {
		tr  *tree.Node
		ll  *list.List
		dep int
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{"depth-2", args{tr1.GetRoot(), list.New(), 2}, []interface{}{1, 4}},
		{"depth-3", args{tr1.GetRoot(), list.New(), 3}, []interface{}{1, 3, 5, 7, 9, 11, 13, 15}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ListOfDepths(tt.args.tr, tt.args.ll, tt.args.dep)
			if tt.args.ll.Len() != len(tt.want) {
				t.Fatalf("wrong linked list length, got: %v, want: %v\n", tt.args.ll.Len(), len(tt.want))
			}
			for _, v := range tt.want {
				if v != tt.args.ll.Front().Value.(tree.Node).Value {
					t.Fatalf("wrong element, got: %v, want: %v\n", tt.args.ll.Front().Value.(tree.Node).Value, v)
				}
			}
		})
	}
}
