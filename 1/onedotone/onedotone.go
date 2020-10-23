package onedotone

import (
	"github.com/Analyse4/mew/algs/binarysearch"
	"github.com/Analyse4/mew/algs/quicksort"
)

// O(N^2)
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

// O(Nlog(N))
func IsStringAllDifferenceV2(s []rune) bool {
	sm := make([]int, len(s))
	for i := range s {
		sm[i] = int(s[i])
	}
	quicksort.QuickSort(sm)
	for _, v := range sm {
		if binarysearch.BinarySearch(v, sm) >= 0 {
			return false
		}
	}
	return true
}
