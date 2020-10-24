package onedotthree

// O(N)
func URLify(s []rune, length int) []rune {
	targetSignI := []rune(" ")
	targetSignR := targetSignI[0]
	tmps := make([]rune, len(s))
	j := 0
	for i := 0; i < length; i++ {
		if s[i] == targetSignR {
			tmps[j] = rune(37)
			j++
			tmps[j] = rune(50)
			j++
			tmps[j] = rune(48)
			j++
			continue
		}
		tmps[j] = s[i]
		j++
	}
	return tmps
}
