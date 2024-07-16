package hindex

import (
	"testing"

	"github.com/teleivo/assertive/assert"
)

func TestHindex(t *testing.T) {
	tests := []struct {
		in   []int
		want int
	}{
		{
			in:   []int{2, 1, 3, 4, 5, 7, 6, 10, 9, 8},
			want: 5,
		},
		{
			in:   []int{2, 100, 2, 100, 2, 2, 2, 2},
			want: 2,
		},
		{
			in:   []int{9, 3, 8, 2, 1, 2, 100, 0, 1, 100},
			want: 4,
		},
	}

	for _, tc := range tests {
		got := hindex(tc.in)

		assert.Equals(t, got, tc.want)
	}
}
