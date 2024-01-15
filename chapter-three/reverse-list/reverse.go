// Package reverse reverses a singly linked list.
package reverse

type node struct {
	Value int
	Next  *node
}

// reverse solves exercise 3.10 3-3.
func reverse(list *node) *node {
	var prev *node
	for cur := list; cur != nil; cur = cur.Next {
		if prev == nil {
			prev = &node{Value: cur.Value}
			continue
		}

		prev = &node{Value: cur.Value, Next: prev}
	}

	return prev
}
