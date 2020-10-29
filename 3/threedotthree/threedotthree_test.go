package threedotthree

import "testing"

func TestThreeDotThree(t *testing.T) {
	e := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	ss := New()
	for _, v := range e {
		ss.Push(v)
	}
	// size
	ra1 := 3
	if ra1 != ss.Size() {
		t.Fatalf("wrong size, want: %v, get: %v\n", ra1, ss.Size())
	}
	ss.Pop()
	//size
	ra2 := 2
	if ra2 != ss.Size() {
		t.Fatalf("wrong size, want: %v, get: %v\n", ra2, ss.Size())
	}
}
