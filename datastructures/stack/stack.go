package stack

import "errors"

// Stack implements a LIFO linear data structure.
type Stack struct {
	top *StackNode
}

func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

func (s *Stack) Peek() (any, error) {
	if s.top == nil {
		return nil, errors.New("stack is empty; cannot peek")
	}
	return s.top.value, nil
}

func (s *Stack) Push(v any) {
	node := NewStackNode(v)
	node.next = s.top
	s.top = node
}

func (s *Stack) Pop() (any, error) {
	if s.top == nil {
		return nil, errors.New("stack is empty; cannot pop")
	}
	v := s.top.value
	s.top = s.top.next
	return v, nil
}

type StackNode struct {
	value any
	next  *StackNode
}

func NewStackNode(v any) *StackNode {
	return &StackNode{value: v}
}
