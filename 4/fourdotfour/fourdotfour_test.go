package fourdotfour

import (
	"testing"

	"github.com/Analyse4/mew/adt/tree"
)

type wt struct {
	ib bool
	h  int
}

func TestIsBalance(t *testing.T) {
	type args struct {
		tr *tree.Tree
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"5-b-tree", args{tree.GenerateMinimalHeightTree([]int{1, 2, 3, 4, 5})}, true},
		{"5-ub-tree", args{tree.GenerateBaddestTree([]int{1, 2, 3, 4, 5})}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBalance(tt.args.tr); got != tt.want {
				t.Errorf("IsBalance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsBalanceAndHeight(t *testing.T) {
	type args struct {
		tr *tree.Tree
	}
	tests := []struct {
		name string
		args args
		want wt
	}{
		{"3-b-tree", args{tree.GenerateMinimalHeightTree([]int{1, 2, 3, 4, 5})}, wt{true, 3}},
		{"5-ub-tree", args{tree.GenerateBaddestTree([]int{1, 2, 3, 4, 5})}, wt{false, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIB, gotH := IsBalanceAndHeight(tt.args.tr); gotIB != tt.want.ib || gotH != tt.want.h {
				t.Errorf("IsBalanceAndHeight() = %v, %v, want %v, %v", gotIB, gotH, tt.want.ib, tt.want.h)
			}
		})
	}
}
