package treedepth

import (
	"testing"

	"github.com/teleivo/assertive/assert"
)

func TestDepth(t *testing.T) {
	tests := []struct {
		in   *node
		want int
	}{
		{
			in:   nil,
			want: 0,
		},
		{
			in:   &node{},
			want: 1,
		},
		{
			in: &node{
				left: &node{},
			},
			want: 2,
		},
		{
			in: &node{
				left:  &node{},
				right: &node{},
			},
			want: 2,
		},
		{
			in: &node{
				left: &node{
					right: &node{
						left: &node{},
					},
				},
				right: &node{},
			},
			want: 4,
		},
		{
			in: &node{
				left: &node{},
				right: &node{
					right: &node{
						left: &node{},
					},
				},
			},
			want: 4,
		},
	}

	t.Run("depth", func(t *testing.T) {
		for _, tc := range tests {

			got := depth(tc.in)

			assert.Equals(t, got, tc.want)
		}
	})

	t.Run("depthIter", func(t *testing.T) {
		for _, tc := range tests {

			got := depthIter(tc.in)

			assert.Equals(t, got, tc.want)
		}
	})
}
