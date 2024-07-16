package nsum

import (
	"testing"

	"github.com/teleivo/assertive/assert"
)

func TestExists(t *testing.T) {
	in := []int{1, 3, 4, 7, 9, 10, 13, 14}

	tests := []struct {
		k    int
		sum  int
		want bool
	}{
		{
			k:    1,
			sum:  14,
			want: true,
		},
		{
			k:    1,
			sum:  11,
			want: false,
		},
		{
			k:    2,
			sum:  5,
			want: true,
		},
		{
			k:    2,
			sum:  8,
			want: true,
		},
		{
			k:    2,
			sum:  14,
			want: true,
		},
		{
			k:    2,
			sum:  15,
			want: true,
		},
		{
			k:    2,
			sum:  17,
			want: true,
		},
		{
			k:    2,
			sum:  24,
			want: true,
		},
		{
			k:    2,
			sum:  27,
			want: true,
		},
		{
			k:    2,
			sum:  6,
			want: false,
		},
		{
			k:    2,
			sum:  25,
			want: false,
		},
		{
			k:    3,
			sum:  8,
			want: true,
		},
		{
			k:    3,
			sum:  14,
			want: true,
		},
		{
			k:    3,
			sum:  20,
			want: true,
		},
		{
			k:    3,
			sum:  7,
			want: false,
		},
		{
			k:    4,
			sum:  17,
			want: true,
		},
		{
			k:    4,
			sum:  18,
			want: true,
		},
		{
			k:    4,
			sum:  20,
			want: true,
		},
	}

	for _, tc := range tests {
		t.Logf("exists(%d, %v, %d)\n", tc.k, in, tc.sum)
		got := exists(tc.k, in, tc.sum)

		assert.Equals(t, got, tc.want)
	}
}

func FuzzExists(f *testing.F) {
	n := 30
	in := make([]int, n)
	for i := range n {
		in[i] = i
	}

	f.Add(0, 1, 2, 3)
	f.Fuzz(func(t *testing.T, i, j, k, l int) {
		if duplicates(i, j, k, l) { // only distinct indices are allowed
			return
		}
		sum := in[i] + in[j] + in[k] + in[l]

		got := exists(4, in, sum)

		assert.True(t, got)
	})
}

func duplicates(values ...int) bool {
	set := make(map[int]struct{})
	for _, v := range values {
		if _, ok := set[v]; ok {
			return true
		}
		set[v] = struct{}{}
	}
	return false
}
