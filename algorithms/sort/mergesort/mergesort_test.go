package mergesort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	input := []int{10, 5, 2, 7, 4, 9, 12, 1, 8, 6, 11, 3}
	MergeSort(input)
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	if !reflect.DeepEqual(input, expected) {
		t.Log("expected:", expected, "got:", input)
		t.Fail()
	}
}
