package tree

// Binary Search Tree: Smaller values always to the left, greater always to the right.
// - Exception: Duplicate values: duplicates can be either left or right aligned, must be specified.
// Unbalanced Tree: No enforcement of branch lengths.
// Balanced Tree: Balanced enough to ensure O(log n) runtime for insert and find operations.
// - Examples: Red-Black Tree, AVL Tree.
// Complete Binary Tree: All values are filled. Last level may be unfilled as long as values are filled left-to-right.
// Full Binary Tree: All nodes have 0 or 2 children, not 1.
// Perfect Tree: Both Full and Complete. Exact same number of nodes in each branch.

// Tree contains a root node, which can contain pointers to further nodes.
type Tree struct {
	root *Node
}

// Node contains a value, and zero or more child nodes.
type Node struct {
	value    any
	children []*Node
}

// IsLeaf determines if a node has no child nodes.
func (n *Node) IsLeaf() bool {
	return len(n.children) == 0
}

// BinaryTree contains a root node, which can contain pointers to further nodes.
type BinaryTree struct {
	root *BinaryNode
}

// BinaryNode contains a value, as well as up to two child nodes.
type BinaryNode struct {
	value any
	left  *BinaryNode
	right *BinaryNode
}

func (n *BinaryNode) IsLeaf() bool {
	return n.left == nil && n.right == nil
}

// TraverseInOrder yields values from left nodes before current and right nodes.
func (n *BinaryNode) TraverseInOrder(valuesCh chan any) {
	if n == nil {
		return
	}
	TraverseInOrder(n.left, valuesCh)
	valuesCh <- n.value
	TraverseInOrder(n.right, valuesCh)
}

// TraversePreOrder yields values from current node before left and right nodes.
func (n *BinaryNode) TraversePreOrder(valuesCh chan any) {
	if n == nil {
		return
	}
	valuesCh <- n.value
	TraverseInOrder(n.left, valuesCh)
	TraverseInOrder(n.right, valuesCh)
}

// TraversePostOrder yields values from left and right nodes before current node.
func (n *BinaryNode) TraversePostOrder(valuesCh chan any) {
	if n == nil {
		return
	}
	TraverseInOrder(n.left, valuesCh)
	TraverseInOrder(n.right, valuesCh)
	valuesCh <- n.value
}
