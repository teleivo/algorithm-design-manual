package treemerge

type node struct {
	Value int
	Left  *node
	Right *node
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
		return &node{Value: value}
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

type list struct {
	head *listNode
	tail *listNode
}

type listNode struct {
	Value int
	prev  *listNode
	next  *listNode
}

func (l *list) Insert(value int) {
	if l.head == nil {
		l.head = &listNode{Value: value}
		l.tail = l.head
		return
	}
	n := &listNode{Value: value, prev: l.tail}
	l.tail.next = n
	l.tail = n
}

// merge solves exercise 3.10 3-14.
// Uses a recursive algorithm that runs in O(N+M) time and O(N+M) space.
func merge(n1, n2 *node, result *list) {
	if n1 == nil && n2 == nil {
		return
	}

	if n1 == nil {
		collect(n2, result)
		return
	}
	if n2 == nil {
		collect(n1, result)
		return
	}

	if n1.Value == n2.Value {
		merge(n1.Left, n2.Left, result)
		result.Insert(n1.Value)
		merge(n1.Right, n2.Right, result)
	} else if n1.Value < n2.Value {
		collect(n1.Left, result)
		result.Insert(n1.Value)
		merge(n1.Right, n2, result)
	} else {
		collect(n2.Left, result)
		result.Insert(n2.Value)
		merge(n1, n2.Right, result)
	}
}

// collect node values in result using DFS in-order traversal.
func collect(n *node, result *list) {
	if n == nil {
		return
	}

	collect(n.Left, result)
	result.Insert(n.Value)
	collect(n.Right, result)
}

func dfs(n *node) {
	if n == nil {
		return
	}

	dfs(n.Left)
	// n.Value
	dfs(n.Right)
}

// merge solves exercise 3.10 3-14.
// Uses an iterative algorithm that runs in O(N+M) time and O(N+M) space.
func mergeIter(n1, n2 *node, result *list) {
	var c1 []int
	var c2 []int

	c1 = collectInOrderIter(n1, c1)
	c2 = collectInOrderIter(n2, c2)

	// use two cursors inserting and advancing the cursor of the smallest element until both trees'
	// values are exhausted
	for i, j := 0, 0; i < len(c1) || j < len(c2); {
		if len(c1) == i {
			result.Insert(c2[j])
			j++
		} else if len(c2) == j {
			result.Insert(c1[i])
			i++
		} else if c1[i] < c2[j] {
			result.Insert(c1[i])
			i++
		} else if c1[i] > c2[j] {
			result.Insert(c2[j])
			j++
		} else { // only insert value once if it exists in both trees
			result.Insert(c1[i])
			i++
			j++
		}
	}
}

func collectInOrderIter(n *node, vals []int) []int {
	if n == nil {
		return vals
	}

	stack := []*node{n} // explicit stack
	visited := make(map[*node]struct{})

	for len(stack) > 0 {
		current := stack[len(stack)-1] // peek
		if _, ok := visited[current.Left]; current.Left != nil && !ok {
			stack = append(stack, current.Left)
			continue
		}

		vals = append(vals, current.Value)
		visited[current] = struct{}{}
		// pop
		if len(stack) > 1 {
			stack = stack[:len(stack)-1]
		} else {
			stack = nil
		}

		if _, ok := visited[current.Right]; current.Right != nil && !ok {
			stack = append(stack, current.Right)
		}
	}

	return vals
}

func collectInOrder(n *node, vals []int) []int {
	if n == nil {
		return vals
	}

	vals = collectInOrder(n.Left, vals)
	vals = append(vals, n.Value)
	vals = collectInOrder(n.Right, vals)
	return vals
}
