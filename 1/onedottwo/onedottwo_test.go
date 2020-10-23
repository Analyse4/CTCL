package onedottwo

import "testing"

func TestCheckPermutation(t *testing.T) {
	e1 := "abcde"
	e2 := "edcba"
	e3 := "aaaaa"
	ra1 := true
	ra2 := false

	if ra1 != CheckPermutation(e1, e2) {
		t.Errorf("check permutation is wrong\n")
	}
	if ra2 != CheckPermutation(e1, e3) {
		t.Errorf("check permutation is wrong\n")
	}
}
