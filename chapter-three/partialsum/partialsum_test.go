package partialsum

import (
	"strconv"
	"testing"

	"github.com/teleivo/algorithm-design-manual/assert"
	"github.com/teleivo/algorithm-design-manual/require"
)

func TestPartial(t *testing.T) {
	vals := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}

	p := New(vals...)

	// test that each partial sum is valid
	for i := range vals {
		t.Run("Sum/"+strconv.Itoa(i), func(t *testing.T) {
			got := p.Sum(i)

			want := gaussSum(i)
			assert.Equals(t, got, want)
		})
	}

	// test that each partial sum is valid after an update
	for i := range vals {
		p.Add(i, i)
		vals[i] += i

		for i := range vals {
			t.Run("SumAfterAdd/"+strconv.Itoa(i), func(t *testing.T) {
				got := p.Sum(i)

				want := partialSum(i, vals)
				require.Equals(t, got, want)
			})
		}
	}
}

func gaussSum(i int) int {
	return (i * (i + 1)) / 2
}

func partialSum(i int, vals []int) int {
	var result int

	for j, v := range vals {
		if j > i {
			break
		}
		result += v
	}

	return result
}
