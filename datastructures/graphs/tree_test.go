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

func TestBinaryTree(t *testing.T) {
	_ = BinaryTree{
		root: &BinaryNode{
			value: 0,
			left:  nil,
			right: nil,
		},
	}
}

func TestTraverseInOrder(t *testing.T) {
	tree := BinaryTree{
		root: &BinaryNode{
			value: 8,
			left: &BinaryNode{
				value: 4,
				left:  &BinaryNode{value: 2},
				right: &BinaryNode{value: 6},
			},
			right: &BinaryNode{
				value: 12,
				left:  &BinaryNode{value: 10},
				right: &BinaryNode{value: 14},
			},
		},
	}
	valuesCh := make(chan any, 7)
	TraverseInOrder(tree.root, valuesCh)
	for i := 1; i <= 7; i++ {
		v := <-valuesCh
		if v != i*2 {
			t.Fail()
		}
	}
}
