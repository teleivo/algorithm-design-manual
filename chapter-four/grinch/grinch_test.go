package grinch

import (
	"testing"

	"github.com/teleivo/assertive/assert"
)

func TestPartition(t *testing.T) {
	tests := map[string]struct {
		in    []int
		want1 []int
		want2 []int
	}{
		"EvenNumberOfPlayers": {
			in:    []int{10, 3, 1, 5, 3, 4, 7, 9},
			want1: []int{5, 7, 9, 10},
			want2: []int{1, 3, 3, 4},
		},
		"OddNumberOfPlayers": {
			in:    []int{10, 1, 5, 3, 4, 7, 9},
			want1: []int{5, 7, 9, 10},
			want2: []int{1, 3, 4},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got1, got2 := partition(tc.in)

			assert.EqualValues(t, got1, tc.want1)
			assert.EqualValues(t, got2, tc.want2)
		})
	}
}
