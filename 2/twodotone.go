package twodotone

import "container/list"

// Time: O(N)
func RemoveDup(ll *list.List) {
	ft := buildFrenqucyTable(ll)
	newll := list.New()
	for e := ll.Front(); e.Next() != nil; e = e.Next() {
		if ft[e.Value] > 1 {
			continue
		}
		newll.PushBack(e.Value)
	}
	ll = newll
}

func buildFrenqucyTable(ll *list.List) map[interface{}]int {
	st := make(map[interface{}]int, ll.Len())
	for e := ll.Front(); e.Next() != nil; e = e.Next() {
		st[e.Value]++
	}
	return st
}
