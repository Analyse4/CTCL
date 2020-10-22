package onedotone

import "sort"

//"github.com/Analyse4/mew/algs/quicksort"

// O(n^2)
func IsStringAllDifferenceV1(s []rune) bool {
	for i, v1 := range s {
		for _, v2 := range s[i+1:] {
			if v1 == v2 {
				return false
			}
		}
	}
	return true
}

func IsStringAllDifferenceV2(s []rune) bool {
	sm := make([]int, len(s))
	for i := range s {
		sm[i] = int(s[i])
	}
	sort.Ints(sm)
}
