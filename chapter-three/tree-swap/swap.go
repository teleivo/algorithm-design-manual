package treeswap

import (
	"fmt"
)

type node struct {
	Value  int
	Parent *node
	Left   *node
	Right  *node
}

type check struct {
	n      *node
	bounds nodeRange
}

type nodeRange struct {
	lower, upper *node
}

func New(values ...int) *node {
	var root *node

	for _, v := range values {
		root = insert(root, nil, v)
	}

	return root
}

func insert(n, parent *node, value int) *node {
	if n == nil {
		return &node{Value: value, Parent: parent}
	}

	if n.Value > value {
		n.Left = insert(n.Left, n, value)
	} else if n.Value < value {
		n.Right = insert(n.Right, n, value)
	}
	return n
}

func search(n *node, value int) *node {
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

func swap(a, b *node) {
	a.Value, b.Value = b.Value, a.Value
}

// swappedIter solves exercise 3.10 3-13.
// Uses an iterative algorithm that runs in O(N) time.
func swappedIter(n *node) (*node, *node) {
	if n == nil {
		return nil, nil
	}

	var err1, err2 *check
	nodes := []check{{n: n}} // explicit stack

	for len(nodes) != 0 {
		current := nodes[0]
		if len(nodes) > 1 {
			nodes = nodes[1:]
		} else {
			nodes = nil
		}

		if current.n.Left != nil {
			bounds := nodeRange{lower: current.bounds.lower, upper: current.n}
			fmt.Printf("check left side of %d = %d lower %v, upper %v\n", current.n.Value, current.n.Left.Value, bounds.lower, bounds.upper)

			if validBounds(bounds) {
				if isOutOfRange(current.n.Left, bounds) {
					if err1 == nil {
						err1 = &check{n: current.n.Left, bounds: bounds}
					} else {
						err2 = &check{n: current.n.Left, bounds: bounds}
					}
				}
				nodes = append(nodes, check{n: current.n.Left, bounds: bounds})
			}
		}
		if current.n.Right != nil {
			bounds := nodeRange{lower: current.n, upper: current.bounds.upper}
			fmt.Printf("check right side of %d = %d lower %v, upper %v\n", current.n.Value, current.n.Right.Value, bounds.lower, bounds.upper)

			if validBounds(bounds) {
				if isOutOfRange(current.n.Right, bounds) {
					if err1 == nil {
						err1 = &check{n: current.n.Right, bounds: bounds}
					} else {
						err2 = &check{n: current.n.Right, bounds: bounds}
					}
				}
				nodes = append(nodes, check{n: current.n.Right, bounds: bounds})
			}
		}

		if err1 != nil && err2 != nil {
			fmt.Println("found 2 errors")
			// nodes in different subtrees have been swapped
			if isInRange(err1.n, err2.bounds) && isInRange(err2.n, err1.bounds) {
				fmt.Println("nodes from different subtrees")
				return err1.n, err2.n
			}

			// TODO clean up
			// nodes in the same subtree have been swapped
			if isLeftChild(err1.n) {
				fmt.Println("left child")
				// try making err2.n child node the parent of err1.n
				newBounds := nodeRange{lower: err1.bounds.lower, upper: err2.n}
				if isInRange(err1.n, newBounds) && isInRange(err1.n.Parent, err2.bounds) {
					return err1.n.Parent, err2.n
				}

				return err2.n.Parent, err1.n
			}
			// try making err2.n child node the parent of err1.n
			newBounds := nodeRange{lower: err2.n, upper: err1.bounds.upper}
			if isInRange(err1.n, newBounds) && isInRange(err1.n.Parent, err2.bounds) {
				return err1.n.Parent, err2.n
			}

			return err2.n.Parent, err1.n
		}
	}

	// if we only found one node that is out of place it has been swapped with the bound it violates
	if err1 != nil && err2 == nil {
		fmt.Printf("err1 %v\n", err1)
		if isSmallerThanLower(err1.n, err1.bounds.lower) {
			return err1.n, err1.bounds.lower
		}
		return err1.n, err1.bounds.upper
	}

	return err1.n.Parent, err2.n
}

func isOutOfRange(n *node, bounds nodeRange) bool {
	if isSmallerThanLower(n, bounds.lower) {
		fmt.Printf("n %d is smaller than lower %d\n", n.Value, bounds.lower.Value)
		return true
	}
	if isGreaterThanUpper(n, bounds.upper) {
		fmt.Printf("n %d is larger than upper %d\n", n.Value, bounds.upper.Value)
		return true
	}

	return false
}

func isInRange(n *node, bounds nodeRange) bool {
	return !isOutOfRange(n, bounds)
}

func isSmallerThanLower(n, lower *node) bool {
	if lower != nil && n.Value < lower.Value {
		return true
	}

	return false
}

func isGreaterThanUpper(n, upper *node) bool {
	if upper != nil && n.Value > upper.Value {
		return true
	}

	return false
}

func validBounds(bounds nodeRange) bool {
	if bounds.lower == nil || bounds.upper == nil {
		return true
	}

	return bounds.upper.Value > bounds.lower.Value
}

func isLeftChild(n *node) bool {
	if n.Parent.Left == nil {
		return false
	}
	return n.Parent.Left.Value == n.Value
}
