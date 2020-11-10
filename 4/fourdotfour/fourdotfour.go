package fourdotfour

import (
	"math"

	"github.com/Analyse4/mew/adt/tree"
)

// (n-1)+(n-3)+(n-7)+...	n items = O(NlogN)
func IsBalance(tr *tree.Tree) bool {
	if tr == nil {
		return true
	}
	hl := 0
	hr := 0
	if tr.Children[0] == nil {
		hl = 0
	} else {
		hl = tr.Children[0].Height()
	}
	if tr.Children[1] == nil {
		hr = 0
	} else {
		hr = tr.Children[1].Height()
	}
	if math.Abs(float64(hl-hr)) > 1 {
		return false
	}
	if !IsBalance(tr.Children[0]) {
		return false
	}
	return IsBalance(tr.Children[1])
}

// time: O(N)
// space: O(N) Number of function stacks called at the same time
func IsBalanceAndHeight(tr *tree.Tree) (bool, int) {
	if tr == nil {
		return true, 0
	}
	hl := 0
	hr := 0
	bl := true
	br := true
	if tr.Children[0] == nil {
		hl = 0
	} else {
		bl, hl = IsBalanceAndHeight(tr.Children[0])
	}
	if tr.Children[1] == nil {
		hr = 0
	} else {
		br, hr = IsBalanceAndHeight(tr.Children[1])
	}
	h := 0
	if hl >= hr {
		h = hl + 1
	} else {
		h = hr + 1
	}
	if !bl {
		return false, h
	}
	if !br {
		return false, h
	}
	if (tr.Children[0] == nil || tr.Children[1] == nil) && math.Abs(float64(hl-hr)) > 1 {
		return false, h
	}
	return true, h
}
