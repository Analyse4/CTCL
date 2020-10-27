package onedotthree

import "github.com/Analyse4/mew/adt/linkedlist"

// O(N)
func DeleteMiddleNode(n *linkedlist.Node) {
	for {
		if n.Next.Next == nil {
			n.Next = nil
			return
		}
		n.Key = n.Next.Key
		n.Value = n.Next.Value
		n = n.Next
	}
}

// O(1)
func DeleteMiddleNodeV2(n *linkedlist.Node) bool {
	if n == nil || n.Next == nil {
		return false
	}
	next := n.Next
	n.Key = next.Key
	n.Value = next.Value
	n.Next = next.Next
	return true
}
