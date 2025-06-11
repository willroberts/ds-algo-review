package stack

import "errors"

var ErrStackIsEmpty = errors.New("stack is empty")

// Stack implements a LIFO linear data structure.
type Stack struct {
	top *StackItem
}

// IsEmpty returns 'true' when there are no items in the Stack.
func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

// Peek returns the first item in the Queue, or an error if the Stack is empty.
func (s *Stack) Peek() (any, error) {
	if s.IsEmpty() {
		return nil, ErrStackIsEmpty
	}
	return s.top.value, nil
}

// Push adds the given item to the top of the Stack.
func (s *Stack) Push(v any) {
	item := NewStackItem(v)
	item.next = s.top
	s.top = item
}

// Pop returns the first item in the Stack, or an error if the Stack is empty.
func (s *Stack) Pop() (any, error) {
	if s.IsEmpty() {
		return nil, ErrStackIsEmpty
	}
	v := s.top.value
	s.top = s.top.next
	return v, nil
}

// StackItem represents an item in the Stack. It contains a value, as well as a reference to the next item.
type StackItem struct {
	value any
	next  *StackItem
}

// NewStackItem stores the given value in a StackItem before returning its reference.
func NewStackItem(v any) *StackItem {
	return &StackItem{value: v}
}
