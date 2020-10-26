package onedotfive

type pair struct {
	key      rune
	distance int
}

func IsOneAway(s1, s2 string) bool {
	switch len(s2) {
	case len(s1):
		st := buildCharFrequencyTable(s1, s2)
		oddCount := 0
		for _, v := range st {
			if v%2 != 0 {
				oddCount++
				if oddCount > 2 {
					return false
				}
			}
		}
		if !isOrdely(s1, s2) {
			return false
		}
		return true
	case len(s1) - 1, len(s1) + 1:
		st := buildCharFrequencyTable(s1, s2)
		foundOdd := false
		for _, v := range st {
			if v%2 != 0 {
				if foundOdd {
					return false
				}
				foundOdd = true
			}
		}
		if !isOrdely(s1, s2) {
			return false
		}
		return true
	default:
		return false
	}
}

func buildCharFrequencyTable(s1, s2 string) map[rune]int {
	st := make(map[rune]int, len(s1)+len(s2))
	for _, c := range s1 {
		st[c]++
	}
	for _, c := range s2 {
		st[c]++
	}
	return st
}

// func isOrdely(s1, s2 string) bool {
// 	o1 := make(map[byte][]int, len(s1))
// 	o2 := make(map[byte][]int, len(s2))
// 	for i := 0; i < len(s1); i++ {
// 		if o1[s1[i]] == nil {
// 			o1[s1[i]] = make([]int, 0)
// 		}
// 		o1[s1[i]] = append(o1[s1[i]], i)
// 	}
// 	for i := 0; i < len(s2); i++ {
// 		if o2[s2[i]] == nil {
// 			o2[s2[i]] = make([]int, 0)
// 		}
// 		o2[s2[i]] = append(o2[s2[i]], i)
// 	}
// 	for i, sc := range s1 {
// 		if o2[byte(sc)] == nil {
// 			continue
// 		}
// 		for _, c := range s1[i:] {
// 			for _, v := range o1[byte(sc)] {

// 			}
// 		}
// 	}
// }

func isOrdely(s1, s2 string) bool {
	return true
}
