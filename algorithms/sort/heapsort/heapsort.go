package heapsort

import (
	"github.com/willroberts/ds-algo-review/datastructures/heap"
)

// HeapSort sorts the given array by constructing a MaxHeap and extracting its heapified values.
func HeapSort(array []int) error {
	h := &heap.MaxHeap{}
	for _, s := range array {
		h.Insert(s)
	}
	for h.Size() > 1 {
		v, err := h.ExtractMax()
		if err != nil {
			return err
		}
		array[h.Size()] = v
	}
	v, err := h.ExtractMax()
	if err != nil {
		return err
	}
	array[0] = v
	return nil
}
