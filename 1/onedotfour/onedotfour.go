package onedotfour

import "strings"

// O(N)
func IsPalindromePermutation(s string) bool {
	s = strings.Replace(s, " ", "", -1)
	s = strings.ToLower(s)
	count := 0
	st := make(map[rune]int, len(s))
	for _, c := range []rune(s) {
		st[c]++
	}
	for _, v := range st {
		if v%2 != 0 {
			count++
			if count > 1 {
				return false
			}
		}
	}
	return true
}
