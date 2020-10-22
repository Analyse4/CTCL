package onedotone

import "testing"

func TestIsStringAllDifferenceV1(t *testing.T) {
	const elen = 50
	er := make([]rune, 0)
	eur := make([]rune, 0)
	for i := 0; i < elen; i++ {
		er = append(er, rune(65))
		er = append(er, rune(65+i))
	}
	rar := false
	raur := true

	if rar != IsStringAllDifferenceV1(er) {
		t.Errorf("string: %v, want: %v, get: %v\n", er, rar, !rar)
	}
	if raur != IsStringAllDifferenceV1(eur) {
		t.Errorf("string: %v, want: %v, get: %v\n", eur, raur, !rar)
	}
}
