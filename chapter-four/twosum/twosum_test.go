package twosum

import (
	"testing"

	"github.com/teleivo/algorithm-design-manual/assert"
)

func TestExists(t *testing.T) {
	tests := []struct {
		in1  []int
		in2  []int
		sum  int
		want bool
	}{
		{
			in1:  []int{11, 4, 7, 1},
			in2:  []int{9, 3, 5, 8},
			sum:  9,
			want: true,
		},
		{
			in1:  []int{11, 4, 7, 1},
			in2:  []int{9, 3, 5, 8},
			sum:  4,
			want: true,
		},
		{
			in1:  []int{11, 4, 7, 1},
			in2:  []int{9, 3, 5, 8},
			sum:  21,
			want: false,
		},
	}

	t.Run("UsingMap", func(t *testing.T) {
		for _, tc := range tests {
			got := existsUsingMap(tc.in1, tc.in2, tc.sum)

			assert.Equals(t, got, tc.want)
		}
	})

	t.Run("UsingSort", func(t *testing.T) {
		for _, tc := range tests {
			got := existsUsingSort(tc.in1, tc.in2, tc.sum)

			assert.Equals(t, got, tc.want)
		}
	})
}
