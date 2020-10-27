package twodottwo

import (
	"container/list"
	"testing"
)

func TestKthToLast(t *testing.T) {
	ra := 888
	ll := list.New()
	for i := 0; i < 100; i++ {
		if i == 97 {
			ll.PushBack(ra)
			continue
		}
		ll.PushBack(i)
	}
	ma := KthToLast(ll, 3)
	if ra != ma {
		t.Errorf("wrong, want: %v, get: %v\n", ra, ma)
	}
}

func TestKthToLastRecursion(t *testing.T) {
	ra := 888
	ll := list.New()
	for i := 0; i < 100; i++ {
		if i == 97 {
			ll.PushBack(ra)
			continue
		}
		ll.PushBack(i)
	}
	ma := KthToLastRecursion(ll, 3)
	if ra != ma {
		t.Errorf("wrong, want: %v, get: %v\n", ra, ma)
	}
}
