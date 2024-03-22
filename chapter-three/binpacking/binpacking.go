// Package binpacking solves exercise 3.10 3-25.
package binpacking

type node struct {
	Value uint
	Left  *node
	Right *node
}

func New(values ...uint) *node {
	var root *node

	for _, v := range values {
		root = insert(root, v)
	}

	return root
}

// FindBestFit returns the number of bins needed to fit objects of given size using the best-fit
// heuristic.
// FindBestFit solves exercise 3.10 3-25.
// Time: O(N log N). For each of the objects a couple of O(log N) operations are needed.
// FindWorstFit can be implemented in the same way replacing findSmallesLargerThan with
// findLargestLarerThan. So instead of going left on a bin with enough space go right.
func FindBestFit(objects []uint, capacity uint) uint {
	var binCount uint
	bins := New()

	for _, object := range objects {
		bin := findSmallestLargerThan(bins, object)
		if bin == nil {
			bins = insert(bins, capacity-object)
			binCount++
		} else {
			bins = deleteNode(bins, bin.Value)
			bins = insert(bins, bin.Value-object)
		}
	}

	return binCount
}

func findSmallestLargerThan(n *node, value uint) *node {
	return findSmallestLargerThanRecur(n, value, nil)
}

func findSmallestLargerThanRecur(n *node, value uint, candidate *node) *node {
	if n == nil {
		return candidate
	}
	if n.Value >= value {
		// fits so we have a candidate but there might be a better-fit i.e. a smaller one that
		// fits as well
		return findSmallestLargerThanRecur(n.Left, value, n)
	}
	return findSmallestLargerThanRecur(n.Right, value, candidate)
}

func deleteNode(n *node, value uint) *node {
	if n == nil {
		return n
	}

	if n.Value == value {
		return findSuccessor(n)
	}

	if n.Left != nil && n.Right == nil && n.Left.Value == value {
		n.Left = findSuccessor(n.Left)
		return n
	} else if n.Left == nil && n.Right != nil && n.Right.Value == value {
		n.Right = findSuccessor(n.Right)
		return n
	}

	if value < n.Value {
		n.Left = deleteNode(n.Left, value)
	} else {
		n.Right = deleteNode(n.Right, value)
	}
	return n
}

func findSuccessor(n *node) *node {
	if n.Left == nil && n.Right == nil {
		return nil
	} else if n.Left != nil && n.Right == nil {
		successor := n.Left
		n.Left = nil
		return successor
	} else if n.Left == nil && n.Right != nil {
		successor := n.Right
		n.Right = nil
		return successor
	}

	minNode := findMin(n.Right)
	n.Right = deleteNode(n.Right, minNode.Value)
	successor := &node{Value: minNode.Value, Left: n.Left, Right: n.Right}
	n.Left, n.Right = nil, nil
	return successor
}

func findMin(n *node) *node {
	if n == nil {
		return nil
	}
	if n.Left == nil {
		return n
	}

	return findMin(n.Left)
}

func insert(n *node, value uint) *node {
	if n == nil {
		return &node{Value: value}
	}

	if n.Value > value {
		n.Left = insert(n.Left, value)
	} else if n.Value < value {
		n.Right = insert(n.Right, value)
	}
	return n
}

func search(n *node, value uint) *node {
	if n == nil {
		return nil
	}

	if n.Value == value {
		return n
	} else if n.Value > value {
		return search(n.Left, value)
	} else {
		return search(n.Right, value)
	}
}
