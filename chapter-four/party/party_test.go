package party

import (
	"slices"
	"strconv"
	"testing"

	"github.com/teleivo/algorithm-design-manual/assert"
)

func TestIntersect(t *testing.T) {
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

		got = test.b.Intersect(a)

		assert.Equals(t, got, test.want)
	}
}

func TestContains(t *testing.T) {
	a := interval{2, 7}
	tests := []struct {
		in   int
		want bool
	}{
		{
			in:   0,
			want: false,
		},
		{
			in:   8,
			want: false,
		},
		{
			in:   5,
			want: true,
		},
		{
			in:   2,
			want: true,
		},
		{
			in:   3,
			want: true,
		},
		{
			in:   6,
			want: true,
		},
		{
			in:   7,
			want: true,
		},
	}

	for _, test := range tests {
		got := a.Contains(test.in)

		assert.Equals(t, got, test.want)
	}
}

func TestAllIntersectingIntervals(t *testing.T) {
	tests := []struct {
		a    interval
		want []interval
	}{
		{
			a: interval{0, 6},
			want: []interval{
				{1, 3},
				{2, 5},
				{4, 8},
			},
		},
		{
			a: interval{1, 3},
			want: []interval{
				{0, 6},
				{2, 5},
			},
		},
		{
			a: interval{2, 5},
			want: []interval{
				{0, 6},
				{1, 3},
				{4, 8},
			},
		},
		{
			a: interval{4, 8},
			want: []interval{
				{0, 6},
				{2, 5},
			},
		},
		{
			a: interval{9, 20},
			want: []interval{
				{10, 19},
				{11, 18},
				{12, 17},
				{13, 16},
				{14, 15},
			},
		},
		{
			a: interval{10, 19},
			want: []interval{
				{9, 20},
				{11, 18},
				{12, 17},
				{13, 16},
				{14, 15},
			},
		},
		{
			a: interval{11, 18},
			want: []interval{
				{9, 20},
				{10, 19},
				{12, 17},
				{13, 16},
				{14, 15},
			},
		},
		{
			a: interval{12, 17},
			want: []interval{
				{9, 20},
				{10, 19},
				{11, 18},
				{13, 16},
				{14, 15},
			},
		},
		{
			a: interval{13, 16},
			want: []interval{
				{9, 20},
				{10, 19},
				{11, 18},
				{12, 17},
				{14, 15},
			},
		},
		{
			a: interval{14, 15},
			want: []interval{
				{9, 20},
				{10, 19},
				{11, 18},
				{12, 17},
				{13, 16},
			},
		},
	}

	var root *node
	for _, test := range tests {
		root = Insert(root, test.a)
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := AllIntersectingIntervals(root, test.a)

			assert.EqualValues(t, got, test.want)
		})
	}
}

func TestAllIntersections(t *testing.T) {
	a := []interval{
		{11, 18},
		{1, 3},
		{0, 6},
		{4, 8},
		{14, 15},
		{9, 20},
		{10, 19},
		{13, 16},
		{2, 5},
		{12, 17},
	}
	tests := []struct {
		in   int
		want []interval
	}{
		{
			in: 0,
			want: []interval{
				{0, 6},
			},
		},
		{
			in: 1,
			want: []interval{
				{0, 6},
				{1, 3},
			},
		},
		{
			in: 2,
			want: []interval{
				{0, 6},
				{1, 3},
				{2, 5},
			},
		},
		{
			in: 3,
			want: []interval{
				{0, 6},
				{1, 3},
				{2, 5},
			},
		},
		{
			in: 4,
			want: []interval{
				{0, 6},
				{2, 5},
				{4, 8},
			},
		},
		{
			in: 8,
			want: []interval{
				{4, 8},
			},
		},
		{
			in: 13,
			want: []interval{
				{9, 20},
				{10, 19},
				{11, 18},
				{12, 17},
				{13, 16},
			},
		},
	}

	var root *node
	for _, in := range a {
		root = Insert(root, in)
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := AllIntersections(root, test.in)

			slices.SortFunc(got, func(a, b interval) int {
				if a[0] < b[0] {
					return -1
				} else if a[0] > b[0] {
					return 1
				}
				return 0
			})
			assert.EqualValues(t, got, test.want)
		})
	}
}

func TestMaxIntersections(t *testing.T) {
	tests := []struct {
		a    []interval
		want int
	}{
		{
			a: []interval{
				{2, 5},
				{0, 6},
				{4, 8},
				{1, 3},
			},
			want: 3,
		},
		{
			a: []interval{
				{9, 20},
				{10, 19},
				{11, 18},
				{12, 17},
				{13, 16},
				{14, 15},
			},
			want: 5,
		},
		{
			a: []interval{
				{0, 6},
				{11, 18},
				{14, 15},
				{4, 8},
				{1, 3},
				{9, 20},
				{10, 19},
				{2, 5},
				{13, 16},
				{12, 17},
			},
			want: 5,
		},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := MaxIntersections(test.a)

			assert.Equals(t, got, test.want)
		})
	}
}

func TestPeekAttendanceTime(t *testing.T) {
	a := []interval{
		{11, 18},
		{1, 3},
		{0, 6},
		{4, 8},
		{14, 15},
		{9, 20},
		{10, 19},
		{13, 16},
		{2, 5},
		{12, 17},
	}

	got := PeekAttendanceTime(a)

	assert.Equals(t, got, 14)
}
