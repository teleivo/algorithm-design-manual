package treebalance

import (
	"testing"

	"github.com/teleivo/algorithm-design-manual/assert"
)

func TestBalance(t *testing.T) {
	tests := map[string]struct {
		in []int
	}{
		"BalancedGivenNil": {
			in: []int{},
		},
		"BalancedSingleNode": {
			in: []int{
				5,
			},
		},
		"Balanced": {
			in: []int{
				5, 1, 11, 8, 3,
			},
		},
		"BalancedWithLeftOneDeeperByOne": {
			in: []int{
				3, 2, 1, 4,
			},
		},
		"BalancedWithRightOneDeeperByOne": {
			in: []int{
				3, 1, 4, 5,
			},
		},
		"UnbalancedWithDeeperRightSubtree": {
			in: []int{
				5, 3, 20, 15, 25, 18, 24, 28, 16, 19,
			},
		},
		"UnbalancedWithDeeperLeftSubtree": {
			in: []int{
				3, 2, 1, 4, -1,
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			n := New(test.in...)

			got := Balance(n)

			assert.True(t, IsBalanced(got))
		})
	}
}

func FuzzBalance(f *testing.F) {
	// unbalanced seed corpus
	f.Add(5, 3, 20, 15, 25, 18, 24, 28, 16, 19)
	f.Fuzz(func(t *testing.T, a int, b int, c int, d int, e int, f int, g int, h int, i int, j int) {
		n := New(a, b, c, d, e, f, g, h, i, j)

		got := Balance(n)

		assert.True(t, IsBalanced(got))
	})
}

func TestIsBalanced(t *testing.T) {
	tests := map[string]struct {
		in   []int
		want bool
	}{
		"BalancedGivenNil": {
			in:   []int{},
			want: true,
		},
		"BalancedSingleNode": {
			in: []int{
				5,
			},
			want: true,
		},
		"Balanced": {
			in: []int{
				5, 1, 11, 8, 3,
			},
			want: true,
		},
		"BalancedWithLeftOneDeeperByOne": {
			in: []int{
				3, 2, 1, 4,
			},
			want: true,
		},
		"BalancedWithRightOneDeeperByOne": {
			in: []int{
				3, 1, 4, 5,
			},
			want: true,
		},
		"UnbalancedWithDeeperRightSubtree": {
			in: []int{
				5, 3, 20, 15, 25, 18, 24, 28, 16, 19,
			},
			want: false,
		},
		"UnbalancedWithDeeperLeftSubtree": {
			in: []int{
				3, 2, 1, 4, -1,
			},
			want: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			n := New(test.in...)

			got := IsBalanced(n)

			assert.Equals(t, got, test.want)
		})
	}
}

func TestHeight(t *testing.T) {
	tests := map[string]struct {
		in   []int
		want int
	}{
		"Nil": {
			in:   []int{},
			want: 0,
		},
		"ParentOnly": {
			in: []int{
				5,
			},
			want: 1,
		},
		"Balanced": {
			in: []int{
				5, 1, 11, 8, 3,
			},
			want: 3,
		},
		"DeeperRightSubtree": {
			in: []int{
				5, 3, 20, 15, 25, 18, 24, 28, 16, 19,
			},
			want: 5,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			n := New(test.in...)

			got := Height(n)

			assert.Equals(t, got, test.want)
		})
	}
}
