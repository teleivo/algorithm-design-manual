// Package implements a singly linked list.
package list

// list solves exercise 3.10 3-7.
type list struct {
	Head *node
	Tail *node
}

type node struct {
	Value int
	Next  *node
}

func New() *list {
	return &list{
		Head: &node{},
		Tail: &node{},
	}
}

func (l *list) search(x int) *node {
	if l == nil {
		return nil
	}

	for n := l.Head.Next; n != nil; n = n.Next {
		if n.Value == x {
			return n
		}
	}

	return nil
}

func (l *list) insert(x *node) {
	x.Next = l.Head.Next
	l.Head.Next = x
	if x.Next == nil {
		x.Next = l.Tail
	}
}

// delete node x from list in constant time.
func (l *list) delete(x *node) {
	if x.Next == l.Tail {
		l.Tail = x
	}
	x.Value = x.Next.Value
	x.Next = x.Next.Next
}

// predecessor finds the logical predecessor of given node in O(N) time.
func (l *list) predecessor(x *node) *node {
	var result *node

	for n := l.Head.Next; n != nil; n = n.Next {
		if n.Value < x.Value && (result == nil || n.Value > result.Value) {
			result = n
		}
	}

	return result
}

// successor finds the logical successor of given node in O(N) time.
func (l *list) successor(x *node) *node {
	var result *node

	for n := l.Head.Next; n != nil; n = n.Next {
		if n.Value > x.Value && (result == nil || n.Value < result.Value) {
			result = n
		}
	}

	return result
}

// minimum finds the smallest node in O(N) time.
// O(N) as deletion has to be done in O(1) time so I cannot hide the cost of searching in the
// unsorted list in the deletion operation.
func (l *list) minimum() *node {
	var result *node

	for n := l.Head.Next; n != nil && n != l.Tail; n = n.Next {
		if result == nil || n.Value < result.Value {
			result = n
		}
	}

	return result
}

// maximum finds the largest node in O(N) time.
// O(N) as deletion has to be done in O(1) time so I cannot hide the cost of searching in the
// unsorted list in the deletion operation.
func (l *list) maximum() *node {
	var result *node

	for n := l.Head.Next; n != nil; n = n.Next {
		if result == nil || n.Value > result.Value {
			result = n
		}
	}

	return result
}
