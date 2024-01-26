package list

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/teleivo/algorithm-design-manual/require"
)

func TestList(t *testing.T) {
	t.Run("InsertAndSearch", func(t *testing.T) {
		l := New()

		got := l.search(10)

		var n *node
		EqualValues(t, "search", l, got, n)

		n = &node{Value: 10}
		l.insert(n)

		want := &list{Head: &node{Next: &node{Value: 10, Next: &node{}}}, Tail: &node{}}
		EqualValues(t, "insert", n, l, want)

		got = l.search(10)

		EqualValues(t, "search", l, got, &node{Value: 10, Next: &node{}})

		n = &node{Value: 11}
		l.insert(n)

		want = &list{Head: &node{Next: &node{Value: 11, Next: &node{Value: 10, Next: &node{}}}}, Tail: &node{}}
		EqualValues(t, "insert", n, l, want)

		got = l.search(10)

		EqualValues(t, "search", l, got, &node{Value: 10, Next: &node{}})
	})

	t.Run("DeleteFromTail", func(t *testing.T) {
		l := New()

		l.insert(&node{Value: 10})
		l.insert(&node{Value: 11})

		got := l.search(10)

		EqualValues(t, "search", l, got, &node{Value: 10, Next: &node{}})

		l.delete(got)

		want := &list{Head: &node{Next: &node{Value: 11, Next: &node{}}}, Tail: &node{}}
		EqualValues(t, "delete", got, l, want)

		got = l.search(11)

		EqualValues(t, "search", l, got, &node{Value: 11, Next: &node{}})

		l.delete(got)

		want = &list{Head: &node{Next: &node{}}, Tail: &node{}}
		EqualValues(t, "delete", got, l, want)
	})

	t.Run("DeleteFromHead", func(t *testing.T) {
		l := New()

		l.insert(&node{Value: 11})
		l.insert(&node{Value: 10})

		got := l.search(10)

		wantNode := &node{Value: 10, Next: &node{Value: 11, Next: &node{}}}
		EqualValues(t, "search", l, got, wantNode)

		l.delete(got)

		want := &list{Head: &node{Next: &node{Value: 11, Next: &node{}}}, Tail: &node{}}
		EqualValues(t, "delete", got, l, want)
	})

	t.Run("SuccessorAndPredecessor", func(t *testing.T) {
		l := New()

		l.insert(&node{Value: 2})
		l.insert(&node{Value: 8})
		l.insert(&node{Value: 1})
		l.insert(&node{Value: 3})
		l.insert(&node{Value: 7})

		n := l.search(3)

		got := l.predecessor(n)

		require.Equals(t, got.Value, 2)

		got = l.successor(n)

		require.Equals(t, got.Value, 7)

		n = l.search(7)

		got = l.predecessor(n)

		require.Equals(t, got.Value, 3)

		got = l.successor(n)

		require.Equals(t, got.Value, 8)

		n = l.search(1)

		got = l.predecessor(n)

		require.Equals(t, 0, got.Value)

		got = l.successor(n)

		require.Equals(t, got.Value, 2)
	})

	t.Run("MinAndMax", func(t *testing.T) {
		l := New()

		l.insert(&node{Value: 2})

		gotMin := l.minimum()
		gotMax := l.maximum()

		require.Equals(t, gotMin.Value, 2)
		require.Equals(t, gotMax.Value, 2)

		l.insert(&node{Value: 8})

		gotMin = l.minimum()
		gotMax = l.maximum()

		require.Equals(t, gotMin.Value, 2)
		require.Equals(t, gotMax.Value, 8)

		l.insert(&node{Value: 1})

		gotMin = l.minimum()
		gotMax = l.maximum()

		require.Equals(t, gotMin.Value, 1)
		require.Equals(t, gotMax.Value, 8)

		l.delete(l.search(1))

		gotMin = l.minimum()
		gotMax = l.maximum()

		require.Equals(t, gotMin.Value, 2)
		require.Equals(t, gotMax.Value, 8)

		l.delete(l.search(8))

		gotMin = l.minimum()
		gotMax = l.maximum()

		require.Equals(t, gotMin.Value, 2)
		require.Equals(t, gotMax.Value, 2)
	})
}

func EqualValues(t *testing.T, method string, in, got, want any) {
	t.Helper()

	if diff := cmp.Diff(want, got, cmpopts.IgnoreUnexported(node{}, list{})); diff != "" {
		t.Fatalf("%s(%v) mismatch (-want +got):\n%s", method, in, diff)
	}
}
