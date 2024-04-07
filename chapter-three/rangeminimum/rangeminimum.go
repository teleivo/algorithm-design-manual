// Package rangeminimum solves exercise 3.10 3-26. b) is solved using two Fenwick trees as discusses
// in the paper "Efficient Range Minimum Queries using Binary Indexed Trees" Mircea DIMA, Rodica
// CETERCHI https://ioinformatics.org/journal/v9_2015_39_44.pdf
package rangeminimum

import (
	"slices"
)

// naive solves a)
type naive struct {
	minimum [][]int
}

// Space: O(N^2)
// Time: O(N^2)
func NewNaive(values []int) *naive {
	minimum := make([][]int, len(values))
	for i := range values {
		minimum[i] = make([]int, len(values))
		for j := i; j < len(values); j++ {
			minimum[i][j] = slices.Min(values[i : j+1])
		}
	}

	return &naive{minimum: minimum}
}

// Min returns the minimum value of the corresponding input values within the inclusive range of i
// to j.
// Time: O(1)
func (n *naive) Min(i, j int) int {
	return n.minimum[i][j]
}
