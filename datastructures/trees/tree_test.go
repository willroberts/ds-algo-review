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
		root: &binaryNode{
			value: 8,
			left:  &binaryNode{value: 4, left: &binaryNode{value: 2}, right: &binaryNode{value: 6}},
			right: &binaryNode{value: 12, left: &binaryNode{value: 10}, right: &binaryNode{value: 14}},
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
		root: &binaryNode{
			value: 2,
			left:  &binaryNode{value: 6, left: &binaryNode{value: 4}, right: &binaryNode{value: 8}},
			right: &binaryNode{value: 12, left: &binaryNode{value: 10}, right: &binaryNode{value: 14}},
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
		root: &binaryNode{
			value: 14,
			left:  &binaryNode{value: 4, left: &binaryNode{value: 2}, right: &binaryNode{value: 6}},
			right: &binaryNode{value: 10, left: &binaryNode{value: 8}, right: &binaryNode{value: 12}},
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
