package set

import (
	"fmt"
	"slices"
	"strconv"
	"testing"

	"github.com/teleivo/assertive/assert"
	"github.com/teleivo/assertive/require"
)

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
