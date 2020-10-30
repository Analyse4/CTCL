package fourdottwo

import "github.com/Analyse4/mew/adt/tree"

func MinimalHeightTree(l []int) *tree.Node {
	if len(l) == 0 {
		return nil
	}
	mid := len(l) / 2
	n := &tree.Node{l[mid], make([]*tree.Node, 2)}
	if len(l) != 1 {
		n.Children[0] = MinimalHeightTree(l[:mid])
		if mid+1 > len(l)-1 {
			n.Children = nil
		}
		n.Children[1] = MinimalHeightTree(l[mid+1:])
	}
	return n
}
