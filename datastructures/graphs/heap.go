package graphs

// MinHeap is a complete binary tree where each node is smaller than its children.
// Can be used to implement a priority queue.
// For the standard library implementation, import 'container/heap'.
type MinHeap struct {
	root *MinHeapNode
}

type MinHeapNode struct {
	value int
	left  *MinHeapNode
	right *MinHeapNode
}

// Insert places the given element at the bottom of the tree before calling Fix().
func (h *MinHeap) Insert(v int) {

	h.Fix()
}

// Fix ensures correct ordering by bubbling up the element to its correct position.
func (h *MinHeap) Fix() {

}

func (h *MinHeap) ExtractMin() {

}

// MaxHeap is a complete binary tree where each node is larger than its children.
type MaxHeap struct {
}
