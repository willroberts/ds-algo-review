package heap

// MinHeap is a complete binary tree where each node is smaller than its children.
// Can be used to implement a priority queue.
// For this implementation, our internal data structure is a slice, instead of using
// custom data structures with pointers to nodes. This is possible since the tree is
// complete, and it simplifies append and swap operations.
// For the standard library implementation, import 'container/heap'.
type MinHeap struct {
	items []int
}

// GetParent returns the parent index of the given node index.
func (h *MinHeap) GetParent(i int) (int, error) {
	if i == 0 {
		return 0, ErrNoRootParent
	}
	// Go rounds down on integer division. Add 1 for odd inputs.
	if i%2 == 1 {
		i += 1
	}
	return (i - 2) / 2, nil
}

// HasLeftChild returns 'true' when the given node index has a left child in the MaxHeap.
func (h *MinHeap) HasLeftChild(i int) bool {
	return h.GetLeftChild(i) < len(h.items)
}

// GetLeftChild returns the left child index of the given node index.
func (h *MinHeap) GetLeftChild(i int) int {
	return i*2 + 1
}

// HasRightChild returns 'true' when the given node index has a right child in the MaxHeap.
func (h *MinHeap) HasRightChild(i int) bool {
	return h.GetRightChild(i) < len(h.items)
}

// GetRightChild returns the right child index of the given node index.
func (h *MinHeap) GetRightChild(i int) int {
	return i*2 + 2
}

// Peek returns the first element in the heap.
func (h *MinHeap) Peek() (int, error) {
	if len(h.items) == 0 {
		return 0, ErrHeapIsEmpty
	}
	return h.items[0], nil
}

// Value returns the contained data of the given node index.
func (h *MinHeap) Value(i int) int {
	return h.items[i]
}

// Insert places the given element at the bottom of the tree before re-heapifying.
func (h *MinHeap) Insert(i int) {
	h.items = append(h.items, i)
	h.heapifyUp()
}

// ExtractMin removes the minimum element from the heap before re-heapifying.
func (h *MinHeap) ExtractMin() (int, error) {
	if len(h.items) == 0 {
		return 0, ErrHeapIsEmpty
	}
	v := h.items[0]
	// Move the last element to the root and shrink the heap.
	h.items[0] = h.items[len(h.items)-1]
	h.items = h.items[:len(h.items)-1]
	h.heapifyDown()
	return v, nil
}

func (h *MinHeap) heapifyUp() {
	if len(h.items) <= 1 {
		return
	}
	i := len(h.items) - 1
	for {
		parent, _ := h.GetParent(i)
		if i > 0 && h.items[parent] > h.items[i] {
			h.items[parent], h.items[i] = h.items[i], h.items[parent]
			i = parent
			continue
		}
		break
	}
}

func (h *MinHeap) heapifyDown() {
	if len(h.items) <= 1 {
		return
	}
	i := 0
	for {
		if !h.HasLeftChild(i) {
			break
		}
		smaller := h.GetLeftChild(i)
		if h.HasRightChild(i) && h.items[h.GetRightChild(i)] < h.items[h.GetLeftChild(i)] {
			smaller = h.GetRightChild(i)
		}
		if h.items[i] < h.items[smaller] {
			break
		}
		h.items[i], h.items[smaller] = h.items[smaller], h.items[i]
		i = smaller
	}
}
