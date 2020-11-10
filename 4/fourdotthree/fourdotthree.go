package fourdotthree

import (
	"container/list"

	"github.com/Analyse4/mew/adt/tree"
)

// O(N)
func ListOfDepths(tr *tree.Tree, ll *list.List, dep int) {
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

// O(N)
func ListOfDepthV1(n *tree.Tree) [][]*tree.Tree {
	if n == nil {
		return nil
	}
	sl := make([][]*tree.Tree, 0)
	current := make([]*tree.Tree, 0)
	current = append(current, n)
	for len(current) != 0 {
		sl = append(sl, current)
		parent := current
		current = make([]*tree.Tree, 0)
		for _, v := range parent {
			if v.Children[0] != nil {
				current = append(current, v.Children[0])
			}
			if v.Children[1] != nil {
				current = append(current, v.Children[1])
			}
		}
	}
	return sl
}
