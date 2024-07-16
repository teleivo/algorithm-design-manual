package telephone

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/teleivo/assertive/assert"
	"github.com/teleivo/assertive/require"
)

func TestWords(t *testing.T) {
	tests := []struct {
		inDigits []int
		inDict   *list
		want     []string
	}{
		{
			inDigits: []int{2, 6, 9},
			inDict:   New("any"),
			want:     []string{"any"},
		},
		{
			inDigits: []int{2, 6, 9},
			inDict:   New("ada", "any", "anyhow", "after", "zoo"),
			want:     []string{"any"},
		},
		{
			inDigits: []int{2, 6, 9},
			inDict:   New("ada", "any", "after", "boy", "zoo", "box", "cow"),
			want:     []string{"any", "box", "boy", "cow"},
		},
		{
			inDigits: []int{2, 3, 3},
			inDict:   New("an", "apple", "ada", "bee", "joy", "pow"),
			want:     []string{"bee"},
		},
		{
			inDigits: []int{2, 3, 3},
			inDict:   New("an", "apple", "ada", "bee", "joyce", "pow"),
			want:     []string{"bee"},
		},
	}

	for _, tc := range tests {
		got := words(tc.inDigits, tc.inDict)

		assert.EqualValues(t, got, tc.want)
	}
}

func TestDict(t *testing.T) {
	t.Run("InsertAndSearch", func(t *testing.T) {
		l := New()

		got := l.search("cow")

		var n *node
		EqualValues(t, "search", l, got, n)

		n = &node{Value: "cow"}
		l.insert(n)

		want := &list{Head: &node{Next: &node{Value: "cow", Next: &node{}}}, Tail: &node{}}
		EqualValues(t, "insert", n, l, want)

		got = l.search("cow")

		EqualValues(t, "search", l, got, &node{Value: "cow", Next: &node{}})

		n = &node{Value: "any"}
		l.insert(n)

		want = &list{Head: &node{Next: &node{Value: "any", Next: &node{Value: "cow", Next: &node{}}}}, Tail: &node{}}
		EqualValues(t, "insert", n, l, want)

		got = l.search("cow")

		EqualValues(t, "search", l, got, &node{Value: "cow", Next: &node{}})
	})

	t.Run("SuccessorAndPredecessor", func(t *testing.T) {
		l := New()

		l.insert(&node{Value: "box"})
		l.insert(&node{Value: "zoo"})
		l.insert(&node{Value: "any"})
		l.insert(&node{Value: "boy"})
		l.insert(&node{Value: "cow"})

		n := l.search("boy")

		got := l.predecessor(n)

		require.Equals(t, got.Value, "box")

		got = l.successor(n)

		require.Equals(t, got.Value, "cow")

		n = l.search("cow")

		got = l.predecessor(n)

		require.Equals(t, got.Value, "boy")

		got = l.successor(n)

		require.Equals(t, got.Value, "zoo")

		n = l.search("any")

		got = l.predecessor(n)

		require.Equals(t, "", got.Value)

		got = l.successor(n)

		require.Equals(t, got.Value, "box")
	})

	t.Run("MinAndMax", func(t *testing.T) {
		l := New()

		l.insert(&node{Value: "box"})

		gotMin := l.minimum()
		gotMax := l.maximum()

		require.Equals(t, gotMin.Value, "box")
		require.Equals(t, gotMax.Value, "box")

		l.insert(&node{Value: "zoo"})

		gotMin = l.minimum()
		gotMax = l.maximum()

		require.Equals(t, gotMin.Value, "box")
		require.Equals(t, gotMax.Value, "zoo")

		l.insert(&node{Value: "any"})

		gotMin = l.minimum()
		gotMax = l.maximum()

		require.Equals(t, gotMin.Value, "any")
		require.Equals(t, gotMax.Value, "zoo")

		l.delete(l.search("any"))

		gotMin = l.minimum()
		gotMax = l.maximum()

		require.Equals(t, gotMin.Value, "box")
		require.Equals(t, gotMax.Value, "zoo")

		l.delete(l.search("zoo"))

		gotMin = l.minimum()
		gotMax = l.maximum()

		require.Equals(t, gotMin.Value, "box")
		require.Equals(t, gotMax.Value, "box")
	})
}

func EqualValues(t *testing.T, method string, in, got, want any) {
	t.Helper()

	if diff := cmp.Diff(want, got, cmpopts.IgnoreUnexported(node{}, list{})); diff != "" {
		t.Fatalf("%s(%v) mismatch (-want +got):\n%s", method, in, diff)
	}
}
