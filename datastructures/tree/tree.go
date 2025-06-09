package tree

// Binary Search Tree: Smaller values always to the left, greater always to the right.
// - Exception: Duplicate values: duplicates can be either left or right aligned, must be specified.
// Unbalanced Tree: No enforcement of branch lengths.
// Balanced Tree: Balanced enough to ensure O(log n) runtime for insert and find operations.
// - Examples: Red-Black Tree, AVL Tree.
// Complete Binary Tree: All values are filled. Last level may be unfilled as long as values are filled left-to-right.
// Full Binary Tree: All nodes have 0 or 2 children, not 1.
// Perfect Tree: Both Full and Complete. Exact same number of nodes in each branch.

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

//

type BinaryTree struct {
	root *binaryNode
}

type BinaryNode interface {
	Value() any
	Left() *binaryNode
	Right() *binaryNode
}

type binaryNode struct {
	value any
	left  *binaryNode
	right *binaryNode
}

func (n *binaryNode) Value() any {
	return n.value
}

func (n *binaryNode) Left() *binaryNode {
	return n.left
}

func (n *binaryNode) Right() *binaryNode {
	return n.right
}

func (n *binaryNode) IsLeaf() bool {
	return n.left == nil && n.right == nil
}

//

// TraverseInOrder yields values from left nodes before current and right nodes.
func TraverseInOrder(n *binaryNode, valuesCh chan any) {
	if n == nil {
		return
	}
	TraverseInOrder(n.Left(), valuesCh)
	valuesCh <- n.Value()
	TraverseInOrder(n.Right(), valuesCh)
}

// TraversePreOrder yields values from current node before left and right nodes.
func TraversePreOrder(n *binaryNode, valuesCh chan any) {
	if n == nil {
		return
	}
	valuesCh <- n.Value()
	TraverseInOrder(n.Left(), valuesCh)
	TraverseInOrder(n.Right(), valuesCh)
}

// TraversePostOrder yields values from left and right nodes before current node.
func TraversePostOrder(n *binaryNode, valuesCh chan any) {
	if n == nil {
		return
	}
	TraverseInOrder(n.Left(), valuesCh)
	TraverseInOrder(n.Right(), valuesCh)
	valuesCh <- n.Value()
}
