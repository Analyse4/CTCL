package onedotfour

import "testing"

func TestIsPalindromePermutation(t *testing.T) {
	e := "Tact Coa"
	ra := true
	if ra != IsPalindromePermutation(e) {
		t.Errorf("Wrong\n")
	}
}
