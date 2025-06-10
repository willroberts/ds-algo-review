package quicksort

// QuickSort sorts the given array by rotating values around a pivot point, with O(n log n) performance.
func QuickSort(array []int) {
	quickSortRange(array, 0, len(array)-1)
}

func quickSortRange(array []int, left int, right int) {
	if left >= right {
		return
	}
	pivot := array[(left+right)/2]
	index := partition(array, left, right, pivot)
	quickSortRange(array, left, index-1)
	quickSortRange(array, index, right)
}

func partition(array []int, left int, right int, pivot int) int {
	for left <= right {
		for array[left] < pivot {
			left++
		}
		for array[right] > pivot {
			right--
		}
		if left <= right {
			array[left], array[right] = array[right], array[left]
			left++
			right--
		}
	}
	// Elements are now in pivot-partitioned order. Partition point is stored at left index.
	return left
}
