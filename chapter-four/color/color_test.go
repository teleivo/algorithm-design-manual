package color

import (
	"testing"

	"github.com/teleivo/assertive/assert"
)

func TestSort(t *testing.T) {
	tests := []struct {
		in   [][2]int
		want [][2]int
	}{
		{
			in:   [][2]int{{1, 2}, {3, 1}, {4, 2}, {6, 3}, {9, 1}},
			want: [][2]int{{3, 1}, {9, 1}, {1, 2}, {4, 2}, {6, 3}},
		},
	}

	for _, tc := range tests {
		got := sort(tc.in)

		assert.EqualValues(t, got, tc.want)
	}
}
