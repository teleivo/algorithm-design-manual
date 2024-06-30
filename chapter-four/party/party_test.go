package party

import (
	"testing"

	"github.com/teleivo/algorithm-design-manual/assert"
)

func TestInterval(t *testing.T) {
	a := interval{2, 7}
	tests := []struct {
		b    interval
		want bool
	}{
		{
			b:    interval{0, 1},
			want: false,
		},
		{
			b:    interval{8, 13},
			want: false,
		},
		{
			b:    interval{0, 5},
			want: true,
		},
		{
			b:    interval{3, 5},
			want: true,
		},
		{
			b:    interval{6, 10},
			want: true,
		},
	}

	for _, test := range tests {
		got := a.Intersect(test.b)

		assert.Equals(t, got, test.want)
	}
}
