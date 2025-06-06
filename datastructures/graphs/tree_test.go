package graphs

import (
	"testing"
)

func TestTree(t *testing.T) {
	n := &Node{
		value:    0,
		children: make([]*Node, 0),
	}
	_ = Tree{
		root: n,
	}
}

func TestTraverseInOrder(t *testing.T) {
	tree := BinaryTree{
		root: &BinaryNode{
			value: 8,
			left:  &BinaryNode{value: 4, left: &BinaryNode{value: 2}, right: &BinaryNode{value: 6}},
			right: &BinaryNode{value: 12, left: &BinaryNode{value: 10}, right: &BinaryNode{value: 14}},
		},
	}
	valuesCh := make(chan any, 7)
	TraverseInOrder(tree.root, valuesCh)
	for i := 1; i <= 7; i++ {
		v := <-valuesCh
		if v != i*2 {
			t.Log("expected", i*2, "got", v)
			t.Fail()
		}
	}
}

func TestTraversePreOrder(t *testing.T) {
	tree := BinaryTree{
		root: &BinaryNode{
			value: 2,
			left:  &BinaryNode{value: 6, left: &BinaryNode{value: 4}, right: &BinaryNode{value: 8}},
			right: &BinaryNode{value: 12, left: &BinaryNode{value: 10}, right: &BinaryNode{value: 14}},
		},
	}
	valuesCh := make(chan any, 7)
	TraversePreOrder(tree.root, valuesCh)
	for i := 1; i <= 7; i++ {
		v := <-valuesCh
		if v != i*2 {
			t.Log("expected", i*2, "got", v)
			t.Fail()
		}
	}
}

func TestTraversePostOrder(t *testing.T) {
	tree := BinaryTree{
		root: &BinaryNode{
			value: 14,
			left:  &BinaryNode{value: 4, left: &BinaryNode{value: 2}, right: &BinaryNode{value: 6}},
			right: &BinaryNode{value: 10, left: &BinaryNode{value: 8}, right: &BinaryNode{value: 12}},
		},
	}
	valuesCh := make(chan any, 7)
	TraversePostOrder(tree.root, valuesCh)
	for i := 1; i <= 7; i++ {
		v := <-valuesCh
		if v != i*2 {
			t.Log("expected", i*2, "got", v)
			t.Fail()
		}
	}
}
