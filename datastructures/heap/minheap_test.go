package heap

import (
	"reflect"
	"testing"
)

func TestMinHeap_GetParent(t *testing.T) {
	h := &MinHeap{items: make([]int, 0)}
	if _, err := h.GetParent(0); err == nil {
		t.Log("got nil, expected err")
		t.Fail()
	}
	if v, err := h.GetParent(1); err != nil || v != 0 {
		t.Log("got", v, "expected", 0)
		t.Fail()
	}
	if v, err := h.GetParent(2); err != nil || v != 0 {
		t.Log("got", v, "expected", 0)
		t.Fail()
	}
	if v, err := h.GetParent(3); err != nil || v != 1 {
		t.Log("got", v, "expected", 1)
		t.Fail()
	}
	if v, err := h.GetParent(4); err != nil || v != 1 {
		t.Log("got", v, "expected", 1)
		t.Fail()
	}
	if v, err := h.GetParent(5); err != nil || v != 2 {
		t.Log("got", v, "expected", 2)
		t.Fail()
	}
	if v, err := h.GetParent(6); err != nil || v != 2 {
		t.Log("got", v, "expected", 2)
		t.Fail()
	}
}

func TestMinHeap_Insert(t *testing.T) {
	h := &MinHeap{items: make([]int, 0)}
	h.Insert(30)
	h.Insert(20)
	h.Insert(10)
	expected := []int{10, 30, 20}
	if !reflect.DeepEqual(h.items, expected) {
		t.Log("expected", expected, "got", h.items)
		t.Fail()
	}
}

func TestMinHeap_ExtractMin(t *testing.T) {
	h := &MinHeap{items: make([]int, 0)}
	h.Insert(10)
	h.Insert(20)
	h.Insert(30)
	h.Insert(40)
	v, err := h.ExtractMin()
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if v != 10 {
		t.Log("expected 10, got", v)
		t.Fail()
	}
	expected := []int{20, 40, 30}
	if !reflect.DeepEqual(h.items, expected) {
		t.Log("expected", expected, "got", h.items)
		t.Fail()
	}
}
