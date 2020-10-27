package twodottwo

import "container/list"

// O(N)
func KthToLast(ll *list.List, k int) interface{} {
	e := ll.Front()
	for i := 0; i < ll.Len()-k; i++ {
		e = e.Next()
	}
	return e.Value
}

// O(N^2)
func KthToLastRecursion(ll *list.List, k int) interface{} {
	if ll.Len() == k {
		return ll.Front().Value
	}
	ll.Remove(ll.Front())
	return KthToLastRecursion(ll, k)
}
