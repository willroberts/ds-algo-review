package bubblesort

// BubbleSort iteratively arranges values in array until sorted.
// Performance is worse than other algorithms at O(n-squared) worse case.
func BubbleSort(array []int) {
	sorted := false
	end := len(array) - 1
	for !sorted {
		sorted = true
		for i := 0; i < end; i++ {
			if array[i] > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i]
				sorted = false
			}
		}
		end--
	}
}
