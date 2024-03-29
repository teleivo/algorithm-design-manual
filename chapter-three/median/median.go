// Package set solves exercise 3.10 3-22. It is based on the solution of 3.10 3-20.
package median

// node is a binary search tree.
type node struct {
	Value         int
	Left          *node
	LeftChildren  uint // children smaller than this node
	Right         *node
	RightChildren uint // children larger than this node
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
	if IsMember(n, value) {
		// to not increment child counters if we do not actually insert
		return n
	}

	if value < n.Value {
		n.LeftChildren++
		n.Left = Insert(n.Left, value)
		return n
	}

	n.RightChildren++
	n.Right = Insert(n.Right, value)
	return n
}

// IsMember returns true if given value is found in node.
func IsMember(n *node, value int) bool {
	if n == nil {
		return false
	}
	if n.Value == value {
		return true
	}

	if n.Value > value {
		return IsMember(n.Left, value)
	}
	return IsMember(n.Right, value)
}

// FindMinK finds the k smallest node. k is one indexed.
// Partially solves exercise 3.10 3-20 as it asks for delete(n, k) but I did not want to implement
// delete again.
// Time: O(log N)
func FindMinK(n *node, k uint) *node {
	if n == nil {
		return n
	}
	if k > n.LeftChildren+n.RightChildren+1 {
		return nil
	}
	curRank := n.LeftChildren + 1
	if k == curRank {
		return n
	}

	if k < curRank {
		return FindMinK(n.Left, k)
	}

	// reduce k by the amount of nodes we remove by going right
	return FindMinK(n.Right, k-n.LeftChildren-1)
}

// Median returns the median value in given node.
// Median solves exercise 3.10 3-22.
// Time: O(log N). Uses either one or two FindMinK operations which run in O(log N) time.
func Median(n *node) int {
	if n == nil {
		// not sure if 0 is ok as median in case of an empty data set
		// adding an error to the result would be an option. keeping it simple for testing
		return 0
	}

	total := n.LeftChildren + n.RightChildren + 1
	if total%2 != 0 {
		k := total/2 + 1
		return FindMinK(n, k).Value
	}
	return (FindMinK(n, total/2).Value + FindMinK(n, total/2+1).Value) / 2
}
