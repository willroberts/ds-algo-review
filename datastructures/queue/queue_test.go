package queue

import "testing"

func TestQueue(t *testing.T) {
	q := &Queue{}
	if !q.IsEmpty() {
		t.Log("expected queue to be empty")
		t.Fail()
	}
	q.Add(123)
	q.Add(456)
	v, err := q.Peek()
	if err != nil {
		t.Log("unexpected error in Queue.Peek")
		t.Log(err)
		t.Fail()
	}
	if v != 123 {
		t.Log("expected first item in queue to be 123")
		t.Fail()
	}
	v, err = q.Remove()
	if err != nil {
		t.Log("unexpected error in Queue.Remove")
		t.Log(err)
		t.Fail()
	}
	if v != 123 {
		t.Log("expected first item in queue to be 123")
		t.Fail()
	}
	v, err = q.Peek()
	if err != nil {
		t.Log("unexpected error in Queue.Peek")
		t.Log(err)
		t.Fail()
	}
	if v != 456 {
		t.Log("expected first item in queue to be 456")
		t.Fail()
	}
}
