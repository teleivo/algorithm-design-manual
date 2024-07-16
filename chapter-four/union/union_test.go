package union

import (
	"slices"
	"testing"

	"github.com/teleivo/assertive/assert"
)

func TestUnion(t *testing.T) {
	tests := []struct {
		a    []int
		b    []int
		want []int
	}{
		{
			a:    []int{1, 4, 17, 2, -1, 3},
			b:    []int{3, 17, 1},
			want: []int{-1, 1, 2, 3, 4, 17},
		},
	}

	for _, tc := range tests {
		got1 := unionUnsorted(tc.a, tc.b)

		slices.Sort(tc.a)
		slices.Sort(tc.b)
		got2 := unionSorted(tc.a, tc.b)

		assert.EqualValues(t, got1, tc.want)
		assert.EqualValues(t, got2, tc.want)
	}
}
