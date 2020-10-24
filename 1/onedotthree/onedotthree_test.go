package onedotthree

import (
	"strings"
	"testing"
)

func TestURLify(t *testing.T) {
	e := []rune("Mr John Smith                          ")
	ra := "Mr%20John%20Smith"
	e = URLify(e, 13)
	sm := strings.Trim(string(e), string(rune(0)))
	if ra != sm {
		t.Errorf("URLify is Wrong\nwant_string: %v\nwant_bytes: %v\nget_string: %v\nget_bytes: %v\n", ra, []rune(ra), sm, []rune(sm))
	}
}
