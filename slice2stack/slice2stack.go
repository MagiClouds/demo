package slice2stack

import "sync"

type stack struct {
	sync.Mutex

	l []int
}

func NewStack() *stack {
	return &stack{
		l: make([]int, 0),
	}
}

func (s *stack) Pop() int {
	s.Lock()
	defer s.Unlock()

	if len(s.l) == 0 {
		return 0
	}

	res := s.l[len(s.l)-1]
	s.l = s.l[:len(s.l)-1]
	return res
}

func (s *stack) Put(i int)  {
	s.Lock()
	defer s.Unlock()

	s.l = append(s.l, i)
}
