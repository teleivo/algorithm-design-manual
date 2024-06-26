package median

import (
	"fmt"
	"slices"
	"strconv"
	"testing"

	"github.com/teleivo/algorithm-design-manual/assert"
	"github.com/teleivo/algorithm-design-manual/require"
)

func FuzzMedianOdd(f *testing.F) {
	f.Add(1, 2, 3, 4, 5, 6, 7)

	f.Fuzz(func(t *testing.T, a int, b int, c int, d int, e int, f int, g int) {
		nums := []int{a, b, c, d, e, f, g}
		n := New(nums...)
		slices.Sort(nums)
		nums = slices.Compact(nums)
		t.Logf("input %v", nums)

		var want int
		k := (len(nums) - 1) / 2
		if len(nums)%2 == 0 {
			want = (nums[k] + nums[k+1]) / 2
		} else {
			want = nums[k]
		}

		assert.Equals(t, Median(n), want)
	})
}

func TestMedian(t *testing.T) {
	t.Run("Odd", func(t *testing.T) {
		n := New(1, 2, 3, 4, 5, 6, 7)
		want := 4

		assert.Equals(t, Median(n), want)
	})
	t.Run("Even", func(t *testing.T) {
		n := New(1, 2, 3, 5, 6, 7)
		want := (3 + 5) / 2

		assert.Equals(t, Median(n), want)
	})
}

func TestIsMember(t *testing.T) {
	in := []int{8, 3, 20, 1, 4, 7, 15, 25, 18, 24, 28, 16, 19}

	n := New(in...)

	for _, v := range in {
		t.Run(strconv.Itoa(v), func(t *testing.T) {
			assert.True(t, IsMember(n, v))
		})
	}
}

func TestFindMinK(t *testing.T) {
	in := []int{8, 3, 20, 1, 4, 7, 15, 25, 18, 24, 28, 16, 19}
	want := slices.Clone(in)
	slices.Sort(want)

	n := New(in...)

	for i, v := range want {
		k := uint(i) + 1
		t.Run(fmt.Sprintf("K=%d", k), func(t *testing.T) {
			got := FindMinK(n, k)

			require.NotNil(t, got)
			assert.Equals(t, got.Value, v)
		})
	}

	t.Run("KLargerThanTotalNodes", func(t *testing.T) {
		in := []int{8, 3, 20, 1, 4, 7, 15, 25, 18, 24, 28, 16, 19}
		n := New(in...)

		got := FindMinK(n, uint(len(in)+1))

		if got != nil {
			t.Fatalf(fmt.Sprintf("got %v want nil instead", got))
		}
	})
}
