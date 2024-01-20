package list

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		in   *node
		want *node
	}{
		{
			in:   nil,
			want: nil,
		},
		{
			in:   &node{Value: 10},
			want: &node{Value: 10},
		},
		{
			in:   &node{Value: 10, Next: &node{Value: 9, Next: &node{Value: 8}}},
			want: &node{Value: 8, Next: &node{Value: 9, Next: &node{Value: 10}}},
		},
	}

	for _, tc := range tests {
		got := reverse(tc.in)

		if diff := cmp.Diff(tc.want, got); diff != "" {
			t.Errorf("%s(%v) mismatch (-want +got):\n%s", "reverse", tc.in, diff)
		}
	}
}
