package treemerge

import (
	"testing"

	"github.com/teleivo/algorithm-design-manual/assert"
)

func TestMergeImplementations(t *testing.T) {
	implementations := map[string]func(*node, *node, *list){
		// "merge":              merge,
		"mergeIter":          mergeIter,
		"mergeIterDfsInline": mergeIterDfsInline,
	}

	for name, f := range implementations {
		t.Run(name+"/OnlyLeftTree", func(t *testing.T) {
			n1 := New(2, 1, 3)
			got := &list{}

			f(n1, nil, got)

			want := []int{
				1, 2, 3,
			}
			assertList(t, got, want)
		})
		t.Run(name+"/OnlyRightTree", func(t *testing.T) {
			n2 := New(2, 1, 3)
			got := &list{}

			f(nil, n2, got)

			want := []int{
				1, 2, 3,
			}
			assertList(t, got, want)
		})
		t.Run(name+"/LeftTreeSmallerThanRight", func(t *testing.T) {
			n1 := New(2, 1, 3)
			n2 := New(5, 4, 6)
			got := &list{}

			f(n1, n2, got)

			want := []int{
				1, 2, 3, 4, 5, 6,
			}
			assertList(t, got, want)
		})
		t.Run(name+"/RightTreeSmallerThanLeft", func(t *testing.T) {
			n1 := New(5, 4, 6)
			n2 := New(2, 1, 3)
			got := &list{}

			f(n1, n2, got)

			want := []int{
				1, 2, 3, 4, 5, 6,
			}
			assertList(t, got, want)
		})
		t.Run(name+"/EqualValues", func(t *testing.T) {
			n1 := New(2, 1, 4)
			n2 := New(2, 4, 6)
			got := &list{}

			f(n1, n2, got)

			want := []int{
				1, 2, 4, 6,
			}
			assertList(t, got, want)
		})
		t.Run(name+"/Mixed", func(t *testing.T) {
			n1 := New(11, 7, 4, 8, 24, 1, 6)
			n2 := New(3, 22, 2, 9, 5, 36)
			got := &list{}

			f(n1, n2, got)

			want := []int{
				1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 22, 24, 36,
			}
			assertList(t, got, want)
		})
	}
}

func TestDelete(t *testing.T) {
	// TODO write shell of merge func so I can implement that right away using delete
	// TODO min should return nil for this case. right now its 0 which can suggest
	// that there is a 0 in the tree
	// TODO write separate min tests
	t.Run("RootWithOnlyLeftChild", func(t *testing.T) {
		n1 := New(5, 2, 1)

		got := deleteNode(n1, 5)

		want := New(2, 1)
		assert.EqualValues(t, got, want)
	})
	t.Run("RootWithOnlyRightChild", func(t *testing.T) {
		n1 := New(5, 7, 9)

		got := deleteNode(n1, 5)

		want := New(7, 9)
		assert.EqualValues(t, got, want)
	})
	t.Run("RootWithLeftAndRightChild", func(t *testing.T) {
		n1 := New(5, 2, 10, 7, 8, 11)

		got := deleteNode(n1, 5)

		want := New(7, 2, 10, 8, 11)
		assert.EqualValues(t, got, want)
	})
	t.Run("LeftParentWithOnlyLeftChild", func(t *testing.T) {
		n1 := New(5, 2, 1)

		got := deleteNode(n1, 2)

		want := New(5, 1)
		assert.EqualValues(t, got, want)
	})
	t.Run("LeftParentWithOnlyRightChild", func(t *testing.T) {
		n1 := New(5, 2, 3)

		got := deleteNode(n1, 2)

		want := New(5, 3)
		assert.EqualValues(t, got, want)
	})
	t.Run("LeftParentWithLeftAndRightChild", func(t *testing.T) {
		n1 := New(10, 3, 12, 8, 6, 5, 7, 9)

		got := deleteNode(n1, 3)

		want := New(10, 12, 8, 6, 5, 7, 9)
		assert.EqualValues(t, got, want)
	})
	t.Run("LeftLeafNode", func(t *testing.T) {
		n1 := New(5, 2, 1)

		got := deleteNode(n1, 1)

		want := New(5, 2)
		assert.EqualValues(t, got, want)
	})
	t.Run("RightParentWithOnlyLeftChild", func(t *testing.T) {
		n1 := New(5, 7, 6)

		got := deleteNode(n1, 7)

		want := New(5, 6)
		assert.EqualValues(t, got, want)
	})
	t.Run("RightParentWithOnlyRightChild", func(t *testing.T) {
		n1 := New(5, 7, 9)

		got := deleteNode(n1, 7)

		want := New(5, 9)
		assert.EqualValues(t, got, want)
	})
	t.Run("RightParentWithLeftAndRightChild", func(t *testing.T) {
		n1 := New(5, 3, 20, 15, 25, 18, 16, 19, 24, 28)

		got := deleteNode(n1, 20)

		want := New(5, 3, 24, 15, 25, 18, 16, 19, 28)
		assert.EqualValues(t, got, want)
	})
	t.Run("RightLeafNode", func(t *testing.T) {
		n1 := New(5, 7, 9)

		got := deleteNode(n1, 9)

		want := New(5, 7)
		assert.EqualValues(t, got, want)
	})
	t.Run("NodeNotFound", func(t *testing.T) {
		n1 := New(5, 7, 9)

		got := deleteNode(n1, 13)

		want := New(5, 7, 9)
		assert.EqualValues(t, got, want)
	})
}

func assertList(t *testing.T, got *list, want []int) {
	var result []int
	for cur := got.head; cur != nil; cur = cur.next {
		result = append(result, cur.Value)
	}

	assert.EqualValues(t, result, want)
}
