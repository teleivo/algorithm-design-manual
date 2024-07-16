package rangeminimum

import (
	"fmt"
	"slices"
	"testing"

	"github.com/teleivo/assertive/assert"
)

func TestNaive(t *testing.T) {
	N := 16
	in := make([]int, N)
	for i := range in {
		in[i] = N - 1 - i
	}

	n := NewNaive(in)

	for i := range in {
		for j := i; j < N; j++ {
			t.Run(fmt.Sprintf("MinFrom%dTo%d", i, j), func(t *testing.T) {
				got := n.Min(i, j)
				want := slices.Min(in[i : j+1])

				assert.Equals(t, got, want)
			})
		}
	}
}

func TestBinaryIndexTree(t *testing.T) {
	N := 16
	in := make([]int, N)
	for i := range in {
		in[i] = N - 1 - i
	}

	b := NewBit(in)

	for i := range in {
		for j := i; j < N; j++ {
			t.Run(fmt.Sprintf("MinFrom%dTo%d", i, j), func(t *testing.T) {
				got := b.Min(i, j)
				want := slices.Min(in[i : j+1])

				assert.Equals(t, got, want)
			})
		}
	}
}
