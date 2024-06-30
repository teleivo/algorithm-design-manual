// Package set solves exercise 4.9 4-13.
// As described in the exercise all entry and exit times are distinct i.e. no interval has equal
// start and end values and no 2 intervals have the same start or end value.
package party

import (
	"math"
)

type interval [2]int

func (a interval) Intersect(b interval) bool {
	return (a[0] < b[0] && b[0] < a[1]) || (a[0] < b[1] && b[1] < a[1]) || (b[0] < a[0] && a[0] < b[1]) || (b[0] < a[1] && a[1] < b[1])
}

func (a interval) Contains(b int) bool {
	return (a[0] <= b && b <= a[1])
}

type node struct {
	left, right *node
	value       interval
	// maxEndpoint is the maximum value of the intervals in the nodes left subtree
	maxEndpoint int
}

func Insert(n *node, a interval) *node {
	if n == nil {
		return &node{value: a, maxEndpoint: math.MinInt}
	}

	if a[0] < n.value[0] {
		n.maxEndpoint = max(n.maxEndpoint, a[1])
		n.left = Insert(n.left, a)
	} else if a[0] > n.value[0] {
		n.right = Insert(n.right, a)
	}

	return n
}

func AllIntersectingIntervals(n *node, a interval) []interval {
	var result []interval
	nodes := []*node{n}

	for len(nodes) > 0 {
		cur := nodes[len(nodes)-1]
		nodes = nodes[:len(nodes)-1]

		if cur.value.Intersect(a) {
			result = append(result, cur.value)
		}

		if cur.left != nil && cur.maxEndpoint >= a[0] {
			nodes = append(nodes, cur.left)
		}
		if cur.right != nil {
			nodes = append(nodes, cur.right)
		}
	}

	return result
}

func AllIntersections(n *node, a int) []interval {
	var result []interval
	nodes := []*node{n}

	for len(nodes) > 0 {
		cur := nodes[len(nodes)-1]
		nodes = nodes[:len(nodes)-1]

		if cur.value.Contains(a) {
			result = append(result, cur.value)
		}

		if cur.left != nil && cur.maxEndpoint >= a {
			nodes = append(nodes, cur.left)
		}
		if cur.right != nil {
			nodes = append(nodes, cur.right)
		}
	}

	return result
}

func MaxIntersections(intervals []interval) int {
	var root *node
	for _, a := range intervals {
		root = Insert(root, a)
	}

	maxIntervals := math.MinInt
	for _, a := range intervals {
		maxIntervals = max(maxIntervals, len(AllIntersectingIntervals(root, a)))
	}
	return maxIntervals
}

// PeekAttendanceTime solves exercise 4.9 4-13.
// Time: O(N log N) - this is only guaranteed with a balanced BST which I did not use but could.
// Space: O(N)
func PeekAttendanceTime(intervals []interval) int {
	var root *node
	for _, a := range intervals {
		root = Insert(root, a)
	}

	maxAttendanceTime := -1
	maxIntervals := math.MinInt
	for _, a := range intervals {
		val := len(AllIntersections(root, a[0]))
		if val > maxIntervals {
			maxIntervals = val
			maxAttendanceTime = a[0]
		}
	}
	return maxAttendanceTime
}
