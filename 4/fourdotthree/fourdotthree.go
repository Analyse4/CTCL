package fourdotthree

import (
	"container/list"

	"github.com/Analyse4/mew/adt/tree"
)

// O(N)
func ListOfDepths(tr *tree.Node, ll *list.List, dep int) {
	if tr == nil {
		return
	}
	if dep == 0 {
		ll.PushBack(tr)
		return
	}
	dep--
	ListOfDepths(tr.Children[0], ll, dep)
	ListOfDepths(tr.Children[1], ll, dep)
	dep++
}
