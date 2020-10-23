package onedottwo

// O(N)
func CheckPermutation(s1, s2 string) bool {
	if len(s1) < len(s2) {
		return false
	}
	s12r := []rune(s1)
	s22r := []rune(s2)
	st := make(map[rune]int, 2*len(s12r))
	for i := 0; i < len(s12r); i++ {
		st[s12r[i]]++
		st[s22r[i]]++
	}
	for _, v := range st {
		if v < 2 {
			return false
		}
	}
	return true
}
