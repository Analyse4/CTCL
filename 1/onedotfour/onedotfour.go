package onedotfour

import "strings"

// O(N)
func IsPalindromePermutation(s string) bool {
	s = strings.Replace(s, " ", "", -1)
	s = strings.ToLower(s)
	foundOdd := false
	st := make(map[rune]int, len(s))
	for _, c := range []rune(s) {
		st[c]++
	}
	for _, v := range st {
		if v%2 != 0 {
			if foundOdd {
				return false
			}
			foundOdd = true
		}
	}
	return true
}
