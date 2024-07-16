package binpacking

import (
	"fmt"
	"testing"

	"github.com/teleivo/assertive/assert"
	"github.com/teleivo/assertive/require"
)

func TestFindBestFit(t *testing.T) {
	objects := []uint{1, 5, 7, 3, 4}

	assert.Equals(t, FindBestFit(objects, 10), uint(2))
}

func FuzzFindBestFit(f *testing.F) {
	f.Add(uint(1), uint(5), uint(7), uint(3), uint(4))

	f.Fuzz(func(t *testing.T, a uint, b uint, c uint, d uint, e uint) {
		objects := []uint{a, b, c, d, e}
		var capacity uint = 10
		total := a + b + c + d + e
		want := total / capacity

		got := FindBestFit(objects, capacity)

		assert.Equals(t, got, want)
	})
}

func TestFindSmallestLargerThan(t *testing.T) {
	n := New(2, 1, 6, 4, 3, 5, 8, 7)

	t.Run("Found", func(t *testing.T) {
		got := findSmallestLargerThan(n, 6)
		assert.Equals(t, got.Value, uint(6))

		got = findSmallestLargerThan(n, 5)
		assert.Equals(t, got.Value, uint(5))

		got = findSmallestLargerThan(n, 4)
		assert.Equals(t, got.Value, uint(4))

		got = findSmallestLargerThan(n, 3)
		assert.Equals(t, got.Value, uint(3))
	})

	got := findSmallestLargerThan(n, 9)
	if got != nil {
		t.Fatalf(fmt.Sprintf("got %v want nil instead", got))
	}
}

func TestDelete(t *testing.T) {
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

func TestFindMin(t *testing.T) {
	t.Run("FullTree", func(t *testing.T) {
		n1 := New(5, 3, 20, 15, 25, 18, 16, 7, 19, 1, 4, 24, 0, 28)

		got := findMin(n1)

		require.NotNil(t, got)
		assert.Equals(t, got.Value, uint(0))
	})
	t.Run("SingleNode", func(t *testing.T) {
		n1 := New(5)

		got := findMin(n1)

		require.NotNil(t, got)
		assert.Equals(t, got.Value, uint(5))
	})
	t.Run("Nil", func(t *testing.T) {
		t.Skip("fix asserting nil")
		n1 := New()

		got := findMin(n1)

		require.Nil(t, got)
	})
}
