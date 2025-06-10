package mergesort

// Sort a slice of integers with O(log n) runtime, at the expense of extra memory space.
func MergeSort(array []int) {
	mergeSortRange(array, make([]int, len(array)), 0, len(array)-1)
}

func mergeSortRange(array []int, tmp []int, leftStart int, rightEnd int) {
	if leftStart >= rightEnd {
		return
	}
	mid := (leftStart + rightEnd) / 2
	mergeSortRange(array, tmp, leftStart, mid)
	mergeSortRange(array, tmp, mid+1, rightEnd)
	mergeHalves(array, tmp, leftStart, rightEnd)
}

func mergeHalves(array []int, tmp []int, leftStart int, rightEnd int) {
	leftEnd := (leftStart + rightEnd) / 2
	rightStart := leftEnd + 1
	size := rightEnd - leftStart + 1 // Number of elements to be copied.

	left := leftStart
	right := rightStart
	i := leftStart

	for left <= leftEnd && right <= rightEnd {
		if array[left] <= array[right] {
			tmp[i] = array[left]
			left++
		} else {
			tmp[i] = array[right]
			right++
		}
		i++
	}

	copy(tmp[i:], array[left:leftEnd+1])                   // Copy left half of array into tmp.
	copy(tmp[i:], array[right:rightEnd+1])                 // Copy right half of array into tmp.
	copy(array[leftStart:], tmp[leftStart:leftStart+size]) // Copy selected range back into array.
}
