package mode

import (
	"testing"

	"github.com/teleivo/algorithm-design-manual/assert"
)

func TestMode(t *testing.T) {
	tests := []struct {
		in   []int
		want int
	}{
		{
			in:   []int{4, 6, 2, 4, 3, 1},
			want: 4,
		},
	}

	t.Run("UsingMap", func(t *testing.T) {
		for _, tc := range tests {
			got := modeUsingMap(tc.in)

			assert.Equals(t, got, tc.want)
		}
	})

	t.Run("UsingSort", func(t *testing.T) {
		for _, tc := range tests {
			got := modeUsingSort(tc.in)

			assert.Equals(t, got, tc.want)
		}
	})
}
