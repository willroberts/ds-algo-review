package binary

// BinarySearch indicates the presence of a specific element of a sorted slice in O(log n) time.
func BinarySearch(data []int, value int) bool {
	left := 1
	right := len(data) - 1
	for left <= right {
		mid := (left + right) / 2
		if data[mid] == value {
			return true
		} else if value < data[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

// BinarySearchWithRecursion uses recursion to indicate the presence of a specific element of a sorted slice in O(log n) time.
func BinarySearchWithRecursion(data []int, value int) bool {
	return binarySearchWithRecursion(data, value, 0, len(data)-1)
}

func binarySearchWithRecursion(data []int, value int, left int, right int) bool {
	if left > right {
		return false
	}
	mid := (left + right) / 2
	if data[mid] == value {
		return true
	}
	if value < data[mid] {
		return binarySearchWithRecursion(data, value, left, mid-1)
	}
	return binarySearchWithRecursion(data, value, mid+1, right)
}
