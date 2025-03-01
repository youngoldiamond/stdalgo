package stack

import "errors"

type Stack[T any] struct {
	data []T
	top  int
}

func New[T any](capacity int) (*Stack[T], error) {
	if capacity < 1 {
		return nil, errors.New("invalid capacity")
	}
	return &Stack[T]{make([]T, capacity), -1}, nil
}

func (s *Stack[T]) Push(key T) {
	s.top++
	if s.top == len(s.data) {
		tmp := make([]T, len(s.data)*2)
		copy(tmp, s.data)
		s.data = tmp
	}
	s.data[s.top] = key
}

func (s *Stack[T]) Pop() T {
	if s.top == -1 {
		panic("Stack is empty")
	}
	s.top--
	return s.data[s.top+1]
}

func (s *Stack[T]) Size() int {
	return s.top + 1
}

func (s *Stack[T]) Empty() bool {
	return s.top == -1
}
