package treeswap

import (
	"testing"

	"github.com/teleivo/algorithm-design-manual/assert"
)

func TestSwapped(t *testing.T) {
	// case A (1 error):
	// a node has been swapped with one of its parents
	// here we only find one node that is out of place
	// the one with which it was swapped is the node its bound it violated
	// case B (2 errors)
	// case B1:
	// the two nodes that have been swapped are found and fit into each
	// others bounds.
	// case B2:
	// one of the error nodes children has been swapped with the other
	// error nodes parents

	// case A: violating upper bound
	t.Run("RootWithImmediateLeftChild", func(t *testing.T) {
		in := New(4, 3, 5, 2)
		n1 := search(in, 3)
		n2 := search(in, 4)
		swap(n1, n2)

		got1, got2 := swappedIter(in)

		if got1 == nil {
			t.Fatal("got nil instead of != nil")
		}
		if got2 == nil {
			t.Fatal("got nil instead of != nil")
		}
		assert.Equals(t, got1.Value, 4)
		assert.Equals(t, got2.Value, 3)
	})
	// case A: violating lower bound
	t.Run("RootWithImmediateRightChild", func(t *testing.T) {
		in := New(3, 4, 2, 5)
		n1 := search(in, 3)
		n2 := search(in, 4)
		swap(n1, n2)

		got1, got2 := swappedIter(in)

		if got1 == nil {
			t.Fatal("got nil instead of != nil")
		}
		if got2 == nil {
			t.Fatal("got nil instead of != nil")
		}
		assert.Equals(t, got1.Value, 3)
		assert.Equals(t, got2.Value, 4)
	})
	// case A: violating upper bound
	t.Run("RootWithDistantRightChild", func(t *testing.T) {
		in := New(6, 3, 5, 4, 2, 1)
		n1 := search(in, 5)
		n2 := search(in, 6)
		swap(n1, n2)

		got1, got2 := swappedIter(in)

		if got1 == nil {
			t.Fatal("got nil instead of != nil")
		}
		if got2 == nil {
			t.Fatal("got nil instead of != nil")
		}
		assert.Equals(t, got1.Value, 6)
		assert.Equals(t, got2.Value, 5)
	})
	// case B1
	t.Run("Leaves", func(t *testing.T) {
		in := New(6, 3, 8, 5, 2, 1, 4)
		n1 := search(in, 1)
		n2 := search(in, 4)
		swap(n1, n2)

		got1, got2 := swappedIter(in)

		if got1 == nil {
			t.Fatal("got nil instead of != nil")
		}
		if got2 == nil {
			t.Fatal("got nil instead of != nil")
		}
		assert.Equals(t, got1.Value, 4)
		assert.Equals(t, got2.Value, 1)
	})
	// case B2
	t.Run("RootWithDistantLeftChild", func(t *testing.T) {
		in := New(6, 3, 8, 5, 2, 1, 4)
		n1 := search(in, 1)
		n2 := search(in, 6)
		swap(n1, n2)

		got1, got2 := swappedIter(in)

		if got1 == nil {
			t.Fatal("got nil instead of != nil")
		}
		if got2 == nil {
			t.Fatal("got nil instead of != nil")
		}
		assert.Equals(t, got1.Value, 1)
		assert.Equals(t, got2.Value, 6)
	})
	// case B2
	t.Run("ParentWithDistantLeftChild", func(t *testing.T) {
		in := New(6, 3, 8, 5, 2, 1, 4)
		n1 := search(in, 1)
		n2 := search(in, 3)
		swap(n1, n2)

		got1, got2 := swappedIter(in)

		if got1 == nil {
			t.Fatal("got nil instead of != nil")
		}
		if got2 == nil {
			t.Fatal("got nil instead of != nil")
		}
		assert.Equals(t, got1.Value, 1)
		assert.Equals(t, got2.Value, 3)
	})
	// case B2
	t.Run("ParentWithDistantRightChild", func(t *testing.T) {
		in := New(8, 3, 7, 1, 6, 4, 5)
		n1 := search(in, 7)
		n2 := search(in, 5)
		swap(n1, n2)

		got1, got2 := swappedIter(in)

		if got1 == nil {
			t.Fatal("got nil instead of != nil")
		}
		if got2 == nil {
			t.Fatal("got nil instead of != nil")
		}
		assert.Equals(t, got1.Value, 5)
		assert.Equals(t, got2.Value, 7)
	})
}
