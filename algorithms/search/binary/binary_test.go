package binary

import "testing"

func TestBinarySearch(t *testing.T) {
	data := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	if !BinarySearch(data, 3) {
		t.Fail()
	}
	if BinarySearch(data, 4) {
		t.Fail()
	}
}

func TestBinarySearchWithRecursion(t *testing.T) {
	data := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	if !BinarySearchWithRecursion(data, 3) {
		t.Fail()
	}
	if BinarySearchWithRecursion(data, 4) {
		t.Fail()
	}
}
