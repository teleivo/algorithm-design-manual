package treedepth

type node struct {
	left  *node
	right *node
}

// depth solves exercise 3.10 3-12.
// Uses a recursive traversal and runs in O(N) time.
func depth(n *node) int {
	return maxDepth(n, 0)
}

func maxDepth(n *node, depth int) int {
	if n == nil {
		return 0
	}

	if n.left != nil && n.right != nil {
		return max(maxDepth(n.left, depth+1), maxDepth(n.right, depth+1))
	}
	if n.left != nil {
		return maxDepth(n.left, depth+1)
	}
	if n.right != nil {
		return maxDepth(n.right, depth+1)
	}

	// either the initial call to maxDepth(n, 1) or I add here
	// feels a bit cleaner here as the int zero value is 0 and
	// if n is nil maxDepth(nil, 1) is odd
	return depth + 1
}

// depth solves exercise 3.10 3-12.
// Uses an iterative algorithm that runs in O(N) time.
func depthIter(n *node) int {
	var maxDepth int
	nodes := []*node{n} // explicit stack

	for len(nodes) != 0 {
		n := nodes[0]
		if n == nil {
			return maxDepth
		}
		if len(nodes) > 1 {
			nodes = nodes[1:]
		} else {
			nodes = nil
		}

		if n.left != nil && n.right != nil {
			nodes = append(nodes, n.left)
			nodes = append(nodes, n.right)
			maxDepth++
		} else if n.left != nil {
			maxDepth++
			nodes = append(nodes, n.left)
		} else if n.right != nil {
			maxDepth++
			nodes = append(nodes, n.right)
		}
	}

	return maxDepth + 1
}
