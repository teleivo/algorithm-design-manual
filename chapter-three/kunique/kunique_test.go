package kunique

import (
	"fmt"
	"testing"

	"github.com/teleivo/assertive/assert"
)

func TestIsKunique(t *testing.T) {
	tests := []struct {
		in   []int
		k    int
		want bool
	}{
		{
			in:   []int{1, 2, 1},
			k:    1,
			want: true,
		},
		{
			in:   []int{1, 2, 1},
			k:    2,
			want: false,
		},
		{
			in:   []int{1, 2, 1},
			k:    3,
			want: false,
		},
		{
			in:   []int{4, 5, 1, 2, 4, 1, 5, 7, 5},
			k:    2,
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%d/%v", test.k, test.in), func(t *testing.T) {
			assert.Equals(t, IsKUnique(test.k, test.in), test.want)
		})
	}
}
