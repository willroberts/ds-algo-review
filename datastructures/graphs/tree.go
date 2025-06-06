package graphs

type Tree struct {
	root *Node
}

type Node struct {
	value    any
	children []*Node
}

func (n *Node) IsLeaf() bool {
	return len(n.children) == 0
}

type BinaryTree struct {
	root *BinaryNode
}

type BinaryNode struct {
	value any
	left  *BinaryNode
	right *BinaryNode
}

func (n *BinaryNode) IsLeaf() bool {
	return n.left == nil && n.right == nil
}

func TraverseInOrder(n *BinaryNode, valuesCh chan any) {
	if n != nil {
		TraverseInOrder(n.left, valuesCh)
		valuesCh <- n.value
		TraverseInOrder(n.right, valuesCh)
	}
}

// Binary Search Tree: Smaller values always to the left, greater always to the right.
// - Exception: Duplicate values: duplicates can be either left or right aligned, must be specified.
// Unbalanced Tree: No enforcement of branch lengths.
// Balanced Tree: Balanced enough to ensure O(log n) runtime for insert and find operations.
// - Examples: Red-Black Tree, AVL Tree.
// Complete Binary Tree: All values are filled. Last level may be unfilled as long as values are filled left-to-right.
// Full Binary Tree: All nodes have 0 or 2 children, not 1.
// Perfect Tree: Both Full and Complete. Exact same number of nodes in each branch.
