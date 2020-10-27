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
