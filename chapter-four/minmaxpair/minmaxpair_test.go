package minmaxpair

import (
	"testing"

	"github.com/teleivo/assertive/assert"
)

func TestPartition(t *testing.T) {
	tests := []struct {
		in   []int
		want [][2]int
	}{
		{
			in:   []int{9, 5, 1, 3},
			want: [][2]int{{1, 9}, {3, 5}},
		},
	}

	for _, tc := range tests {
		got := partition(tc.in)

		assert.EqualValues(t, got, tc.want)
	}
}
