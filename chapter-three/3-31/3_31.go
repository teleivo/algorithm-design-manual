// Package problem_331 solves exercise 3.10 3-31.
package problem_331

import (
	"fmt"
)

// numbers allows inserting, searching and deleting integers in O(1) time. numbers does not handle
// inserting duplicates.
// TODO I could not come up with a better name.
type numbers struct {
	cnt int // cnt holds the count of inserted numbers
	a   []int
	b   []int
}

// New returns numbers which can hold m integers in the range of 1 <= x <= n.
func New(n, m int) *numbers {
	return &numbers{
		a: make([]int, n),
		b: make([]int, m),
	}
}

// Insert inserts x if x was not already inserted.
// Time: O(1)
func (n *numbers) Insert(x int) {
	if n.Search(x) {
		return
	}
	if n.cnt == len(n.b) {
		panic(fmt.Errorf("cannot insert as numbers is at capacity %d", len(n.b)))
	}

	n.cnt++
	n.a[x] = n.cnt
	n.b[n.cnt] = x
}

// Search returns true if x was inserted and false otherwise.
// Time: O(1)
func (n *numbers) Search(x int) bool {
	n.validateUpperBound(x)

	return n.a[x] <= n.cnt && n.b[n.a[x]] == x
}

// Delete deletes x if it has been inserted before.
// Delete works by swapping the deleted element with the last inserted element.
// Time: O(1)
func (n *numbers) Delete(x int) {
	if !n.Search(x) {
		return
	}

	n.a[n.b[n.cnt]] = n.a[x]
	n.b[n.a[x]] = n.b[n.cnt]
	n.cnt--
}

func (n *numbers) validateUpperBound(x int) {
	if x >= len(n.a) {
		panic(fmt.Errorf("x is greater than the upper bound of %d", len(n.a)))
	}
}
