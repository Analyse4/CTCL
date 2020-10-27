package onedotthree

import (
	"log"
	"testing"

	"github.com/Analyse4/mew/adt/linkedlist"
)

func TestDeleteMiddleNode(t *testing.T) {
	ll := linkedlist.New()
	for i := 103; i > 96; i-- {
		ll.Insert(rune(i), rune(i))
	}
	var e *linkedlist.Node
	for e = ll.First; e != nil; e = e.Next {
		if e.Value.(rune) == rune(99) {
			break
		}
	}
	if e == nil {
		t.Fatal("example error")
	}
	DeleteMiddleNode(e)
	if ll.Get(rune(99)) != nil {
		for e := ll.First; e.Next != nil; e = e.Next {
			log.Println(e.Value)
		}
		t.Error("wrong")
	}
}
