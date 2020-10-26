package onedotfive

import "testing"

func TestIsOneAway(t *testing.T) {
	e := [][]string{[]string{"pale", "ple"},
		[]string{"pales", "pale"},
		[]string{"pale", "bale"},
		[]string{"pale", "bake"}}
	ra := []bool{true, true, true, false}

	for i, s := range e {
		if ra[i] != IsOneAway(s[0], s[1]) {
			t.Errorf("Wrong, s1: %v, s2: %v\n", s[0], s[1])
		}
	}
}
