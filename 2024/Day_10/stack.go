package main

type Stack struct {
	top    *node
	length int
}
type node struct {
	value any
	prev  *node
}

func NewStack() *Stack {
	return &Stack{nil, 0}
}

func Length(s *Stack) int {
	return s.length
}

func (s *Stack) Push(v any) {
	n := &node{v, s.top}
	s.top = n
	s.length++
}

func (s *Stack) Pop() any {
	if s.length == 0 {
		return nil
	}

	n := s.top
	s.top = n.prev
	s.length--
	return n.value
}

func (s *Stack) Peek() any {
	if s.length == 0 {
		return nil
	}
	return s.top.value
}