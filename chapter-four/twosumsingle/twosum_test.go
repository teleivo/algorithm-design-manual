package twosum

import (
	"testing"

	"github.com/teleivo/algorithm-design-manual/assert"
)

func TestExists(t *testing.T) {
	tests := []struct {
		in   []int
		in2  []int
		sum  int
		want bool
	}{
		{
			in:   []int{11, 4, 7, 1},
			sum:  5,
			want: true,
		},
		{
			in:   []int{11, 4, 7, 1},
			sum:  12,
			want: true,
		},
		{
			in:   []int{11, 4, 7, 1},
			sum:  21,
			want: false,
		},
	}

	t.Run("UsingMap", func(t *testing.T) {
		for _, tc := range tests {
			got := existsUsingMap(tc.in, tc.sum)

			assert.Equals(t, got, tc.want)
		}
	})

	t.Run("UsingSort", func(t *testing.T) {
		for _, tc := range tests {
			got := existsUsingSort(tc.in, tc.sum)

			assert.Equals(t, got, tc.want)
		}
	})
}
