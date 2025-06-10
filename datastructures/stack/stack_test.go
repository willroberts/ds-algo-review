package stack

import "testing"

func TestStack(t *testing.T) {
	s := &Stack{}
	if !s.IsEmpty() {
		t.Log("expected Stack to be empty")
		t.Fail()
	}
	s.Push(123)
	s.Push(456)
	v, err := s.Peek()
	if err != nil {
		t.Log("unexpected error in Stack.Peek")
		t.Log(err)
		t.Fail()
	}
	if v != 456 {
		t.Log("expected top item in Stack to be 456")
		t.Fail()
	}
	v, err = s.Pop()
	if err != nil {
		t.Log("unexpected error in Stack.Pop")
		t.Log(err)
		t.Fail()
	}
	if v != 456 {
		t.Log("expected first item in Stack to be 456")
		t.Fail()
	}
	v, err = s.Peek()
	if err != nil {
		t.Log("unexpected error in Stack.Peek")
		t.Log(err)
		t.Fail()
	}
	if v != 123 {
		t.Log("expected first item in Stack to be 123")
		t.Fail()
	}
}
