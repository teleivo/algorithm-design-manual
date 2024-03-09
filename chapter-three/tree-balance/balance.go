// Package treebalance solves exercise 3.10 3-15.
package treebalance

// node is a binary search tree.
type node struct {
	Value int
	Left  *node
	Right *node
}

// New creates a node of given values inserting them in order.
func New(values ...int) *node {
	var root *node
	for _, v := range values {
		root = Insert(root, v)
	}
	return root
}

// Insert value into BST node.
func Insert(n *node, value int) *node {
	if n == nil {
		return &node{Value: value}
	}

	if value < n.Value {
		n.Left = Insert(n.Left, value)
	}
	if value > n.Value {
		n.Right = Insert(n.Right, value)
	}
	return n
}

// IsBalanced returns true if the height of left and right subtrees is never more than 1.
func IsBalanced(n *node) bool {
	if n == nil {
		return true
	}

	if n.Left == nil && n.Right == nil {
		return true
	}
	if n.Left != nil && n.Right == nil {
		if n.Left.Left == nil && n.Left.Right == nil {
			return true
		}
		return false
	}
	if n.Left == nil && n.Right != nil {
		if n.Right.Left == nil && n.Right.Right == nil {
			return true
		}
		return false
	}

	return IsBalanced(n.Left) && IsBalanced(n.Right)
}

// Height computes the max height of the given node.
func Height(n *node) int {
	return heightRecur(n, 0)
}

func heightRecur(n *node, height int) int {
	if n == nil {
		return height
	}

	hLeft := heightRecur(n.Left, height+1)
	hRight := heightRecur(n.Right, height+1)
	return max(hLeft, hRight)
}

// Balance solves exercise 3.10 3-15.
// Time: O(N)
// Space: O(N) in case of an unbalanced tree due to collecting values in a slice.
func Balance(n *node) *node {
	if IsBalanced(n) {
		return n
	}

	var ordered []int
	ordered = collectInOrder(n, ordered)
	return insertBalanced(nil, ordered)
}

func collectInOrder(n *node, result []int) []int {
	if n == nil {
		return result
	}

	if n.Left != nil {
		result = collectInOrder(n.Left, result)
	}
	result = append(result, n.Value)
	if n.Right != nil {
		result = collectInOrder(n.Right, result)
	}

	return result
}

func insertBalanced(n *node, values []int) *node {
	if len(values) == 0 {
		return n
	}

	i := len(values) / 2
	n = Insert(n, values[i])
	if upper := i - 1; i > 0 && len(values) > upper {
		n = insertBalanced(n, values[0:upper])
	}
	if lower := i + 1; len(values) > lower {
		n = insertBalanced(n, values[lower:])
	}
	return n
}
