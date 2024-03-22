// Package treemerge solves exercise 3.10 3-14. I implemented it recursively and iteratively.
package treemerge

import (
	"fmt"
)

type node struct {
	Value int
	Left  *node
	Right *node
}

func New(values ...int) *node {
	var root *node

	for _, v := range values {
		root = insert(root, v)
	}

	return root
}

func deleteNode(n *node, value int) *node {
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

func insert(n *node, value int) *node {
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

// TODO solve recursively without mutation?
// merge solves exercise 3.10 3-14.
// Uses a recursive algorithm that runs in O(N+M) time and O(N+M) space.
func merge(n1, n2 *node, result *list) {
	if n1 == nil && n2 == nil {
		return
	}

	n1Next, n2Next := mergeRecur(n1, n2, result)
	if n1Next == nil {
		collect(n2Next, result)
	}
	if n2Next == nil {
		collect(n1Next, result)
	}
}

func mergeRecur(n1, n2 *node, result *list) (*node, *node) {
	return nil, nil
}

func collectUntil(n *node, upper int, result *list) *node {
	if n == nil {
		return n
	}

	out := collectUntil(n.Left, upper, result)
	if out != nil {
		return out
	}

	if n.Value >= upper {
		return n
	}

	result.Insert(n.Value)
	return collectUntil(n.Right, upper, result)
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

// mergeRecurMutate solves exercise 3.10 3-14.
// Uses a recursive algorithm that mutates the inputs! This is not a good API and also an inefficien
// implementation. It is easy to read though. The nodes themselves to keep track of the work that is
// left. I implemented it just to explore how I could implement a recursive algorithm without
// mutation.
func mergeRecurMutate(n1, n2 *node, result *list) {
	if n1 == nil && n2 == nil {
		return
	}

	if n1 != nil && n2 != nil {
		n1Min := findMin(n1)
		n2Min := findMin(n2)

		if n1Min.Value < n2Min.Value {
			result.Insert(n1Min.Value)
			n1 = deleteNode(n1, n1Min.Value)
		} else if n2Min.Value < n1Min.Value {
			result.Insert(n2Min.Value)
			n2 = deleteNode(n2, n2Min.Value)
		} else {
			result.Insert(n1Min.Value)
			n1 = deleteNode(n1, n1Min.Value)
			n2 = deleteNode(n2, n2Min.Value)
		}
	} else if n1 != nil {
		n1Min := findMin(n1)
		result.Insert(n1Min.Value)
		n1 = deleteNode(n1, n1Min.Value)
	} else {
		n2Min := findMin(n2)
		result.Insert(n2Min.Value)
		n2 = deleteNode(n2, n2Min.Value)
	}

	mergeRecurMutate(n1, n2, result)
}

// mergeIter solves exercise 3.10 3-14.
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

// mergeIterDfsInline solves exercise 3.10 3-14.
// Uses an iterative algorithm that runs in O(N+M) time and O(N+M) space.
// SUPER UGLY! Don't want to spend time refactoring though. I just wanted to explore doing an
// iterative dfs on both trees.
func mergeIterDfsInline(n1, n2 *node, result *list) {
	if n1 == nil && n2 == nil {
		return
	}

	var n1Stack []*node
	if n1 != nil {
		n1Stack = []*node{n1}
	}
	n1Visited := make(map[*node]struct{})
	var n2Stack []*node
	if n2 != nil {
		n2Stack = []*node{n2}
	}
	n2Visited := make(map[*node]struct{})

	for len(n1Stack) > 0 || len(n2Stack) > 0 {
		// collect n2 in order if n1 is empty
		if len(n1Stack) == 0 {
			c2 := n2Stack[len(n2Stack)-1]
			if _, visited := n2Visited[c2.Left]; c2.Left != nil && !visited {
				n2Stack = append(n2Stack, c2.Left)
				continue
			}

			// insert, mark and pop
			result.Insert(c2.Value)
			n2Visited[c2] = struct{}{}
			if len(n2Stack) > 1 {
				n2Stack = n2Stack[:len(n2Stack)-1]
			} else {
				n2Stack = nil
			}

			if _, visited := n2Visited[c2.Right]; c2.Right != nil && !visited {
				n2Stack = append(n2Stack, c2.Right)
				continue
			}

			continue
		}
		// collect n1 in order if n1 is empty
		if len(n2Stack) == 0 {
			c1 := n1Stack[len(n1Stack)-1]
			if _, visited := n1Visited[c1.Left]; c1.Left != nil && !visited {
				n1Stack = append(n1Stack, c1.Left)
				continue
			}

			// insert, mark and pop
			result.Insert(c1.Value)
			n1Visited[c1] = struct{}{}
			if len(n1Stack) > 1 {
				n1Stack = n1Stack[:len(n1Stack)-1]
			} else {
				n1Stack = nil
			}

			if _, visited := n1Visited[c1.Right]; c1.Right != nil && !visited {
				n1Stack = append(n1Stack, c1.Right)
				continue
			}

			continue
		}

		c1 := n1Stack[len(n1Stack)-1]
		c2 := n2Stack[len(n2Stack)-1]

		// go to the leftmost node in both trees
		if _, visited := n1Visited[c1.Left]; c1.Left != nil && !visited {
			n1Stack = append(n1Stack, c1.Left)
			if _, visited := n2Visited[c2.Left]; c2.Left != nil && !visited {
				n2Stack = append(n2Stack, c2.Left)
			}
			continue
		}
		if _, visited := n2Visited[c2.Left]; c2.Left != nil && !visited {
			n2Stack = append(n2Stack, c2.Left)
			if _, visited := n1Visited[c1.Left]; c1.Left != nil && !visited {
				n1Stack = append(n1Stack, c1.Left)
			}
			continue
		}

		if c1.Value < c2.Value {
			// insert, mark and pop
			result.Insert(c1.Value)
			n1Visited[c1] = struct{}{}
			if len(n1Stack) > 1 {
				n1Stack = n1Stack[:len(n1Stack)-1]
			} else {
				n1Stack = nil
			}
			if _, visited := n1Visited[c1.Right]; c1.Right != nil && !visited {
				n1Stack = append(n1Stack, c1.Right)
			}
		} else if c2.Value < c1.Value {
			// insert, mark and pop
			result.Insert(c2.Value)
			n2Visited[c2] = struct{}{}
			if len(n2Stack) > 1 {
				n2Stack = n2Stack[:len(n2Stack)-1]
			} else {
				n2Stack = nil
			}
			if _, visited := n2Visited[c2.Right]; c2.Right != nil && !visited {
				n2Stack = append(n2Stack, c2.Right)
			}
		} else {
			// insert once as they are equal, mark and pop both
			result.Insert(c1.Value)

			n1Visited[c1] = struct{}{}
			if len(n1Stack) > 1 {
				n1Stack = n1Stack[:len(n1Stack)-1]
			} else {
				n1Stack = nil
			}
			if _, visited := n1Visited[c1.Right]; c1.Right != nil && !visited {
				n1Stack = append(n1Stack, c1.Right)
			}

			n2Visited[c2] = struct{}{}
			if len(n2Stack) > 1 {
				n2Stack = n2Stack[:len(n2Stack)-1]
			} else {
				n2Stack = nil
			}
			if _, visited := n2Visited[c2.Right]; c2.Right != nil && !visited {
				n2Stack = append(n2Stack, c2.Right)
			}
		}
	}
}

func printStack(name string, stack []*node) {
	var values []int
	for _, v := range stack {
		values = append(values, v.Value)
	}
	fmt.Printf("%q %v\n", name, values)
}
