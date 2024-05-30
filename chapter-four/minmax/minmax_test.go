package minmax

import (
	"slices"
	"testing"

	"github.com/teleivo/algorithm-design-manual/assert"
)

func TestMaximize(t *testing.T) {
	tests := []struct {
		in    []int
		want1 int
		want2 int
	}{
		{
			in:    []int{6, 13, 19, 3, 8},
			want1: 3,
			want2: 19,
		},
	}

	t.Run("Unsorted", func(t *testing.T) {
		for _, tc := range tests {
			got1, got2 := maximizeUnsorted(tc.in)

			assert.Equals(t, got1, tc.want1)
			assert.Equals(t, got2, tc.want2)
		}
	})

	t.Run("Sorted", func(t *testing.T) {
		for _, tc := range tests {
			in := make([]int, len(tc.in))
			copy(in, tc.in)
			slices.Sort(in)
			got1, got2 := maximizeSorted(in)

			assert.Equals(t, got1, tc.want1)
			assert.Equals(t, got2, tc.want2)
		}
	})
}

func TestMinimize(t *testing.T) {
	tests := []struct {
		in    []int
		want1 int
		want2 int
	}{
		{
			in:    []int{6, 13, 19, 3, 8},
			want1: 6,
			want2: 8,
		},
	}

	t.Run("Unsorted", func(t *testing.T) {
		for _, tc := range tests {
			got1, got2 := minimizeUnsorted(tc.in)

			assert.Equals(t, got1, tc.want1)
			assert.Equals(t, got2, tc.want2)
		}
	})

	t.Run("Sorted", func(t *testing.T) {
		for _, tc := range tests {
			in := make([]int, len(tc.in))
			copy(in, tc.in)
			slices.Sort(in)
			got1, got2 := minimizeSorted(in)

			assert.Equals(t, got1, tc.want1)
			assert.Equals(t, got2, tc.want2)
		}
	})
}
