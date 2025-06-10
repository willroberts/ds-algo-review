package bubblesort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	input := []int{15, 3, 2, 1, 9, 5, 7, 8, 6}
	expected := []int{1, 2, 3, 5, 6, 7, 8, 9, 15}
	BubbleSort(input)
	if !reflect.DeepEqual(input, expected) {
		t.Log("got", input, "expected", expected)
		t.Fail()
	}
}
