package threedottwo

import "sync"

type Stack struct {
	buffer []int
	len    int
	min    int
	mu     *sync.Mutex
}

func New() *Stack {
	return &Stack{make([]int, 0), 0, 0, new(sync.Mutex)}
}

// O(1)
func (s *Stack) Push(v int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.buffer = append(s.buffer, v)
	if s.len == 0 {
		s.min = v
	} else {
		if v < s.min {
			s.min = v
		}
	}
	s.len++
}

// O(1)
func (s *Stack) Pop() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.buffer[s.len-1]
}

// O(1)
func (s *Stack) Min() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.min
}
