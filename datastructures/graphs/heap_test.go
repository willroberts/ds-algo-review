package graphs

import (
	"testing"
)

func TestHeapInsert(t *testing.T) {
	h := MinHeap{
		root: &MinHeapNode{value: 0, left: &MinHeapNode{value: 1}, right: &MinHeapNode{value: 2}},
	}
	h.Insert(2)
}
