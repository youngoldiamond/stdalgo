package stack

import "errors"

type Stack struct {
	data []int
	top  int
}

func New(capacity int) (*Stack, error) {
	if capacity < 1 {
		return nil, errors.New("invalid capacity")
	}
	return &Stack{make([]int, capacity), -1}, nil
}

func (s *Stack) Push(key int) {
	s.top++
	if s.top == len(s.data) {
		tmp := make([]int, len(s.data)*2)
		copy(tmp, s.data)
		s.data = tmp
	}
	s.data[s.top] = key
}

func (s *Stack) Pop() int {
	if s.top == -1 {
		panic("Stack is empty")
	}
	s.top--
	return s.data[s.top+1]
}

func (s *Stack) IsEmpty() bool {
	return s.top == -1
}
