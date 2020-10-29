package threedottwo

import "testing"

func TestMin(t *testing.T) {
	e := []int{5, 4, 3, 2, 1, 6, 7, 8, 9, 0}
	s := New()
	for _, v := range e {
		s.Push(v)
	}
	ra := 0

	if ra != s.Min() {
		t.Fatalf("min is wrong, want: %v, get: %v\n", ra, s.Min())
	}
}

func TestRace(t *testing.T) {
	ch := make(chan bool)
	s := New()
	go func() {
		s.Min()
		s.Push(1)
		ch <- true
	}()
	s.Push(2)
	s.Pop()
	<-ch
}
