package twodotone

import (
	"container/list"
	"testing"
)

func TestRemoveDup(t *testing.T) {
	l := list.New()
	for i := 0; i < 100; i++ {
		if i == 44 || i == 88 {
			l.PushBack(88)
			continue
		}
		if i == 4 || i == 8 {
			l.PushBack(8)
			continue
		}
		l.PushBack(i)
	}
	RemoveDup(l)
	if IsHaveDup(l) {
		t.Errorf("wrong, get: %v\n", l)
	}
}

func IsHaveDup(ll *list.List) bool {
	ft := buildFrenqucyTable(ll)
	for _, v := range ft {
		if v > 1 {
			return false
		}
	}
	return true
}
