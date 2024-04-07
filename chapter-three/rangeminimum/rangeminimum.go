// Package rangeminimum solves exercise 3.10 3-26. b) is solved using two Fenwick trees as discusses
// in the paper "Efficient Range Minimum Queries using Binary Indexed Trees" Mircea DIMA, Rodica
// CETERCHI https://ioinformatics.org/journal/v9_2015_39_44.pdf
package rangeminimum

import (
	"math"
	"slices"
)

// naive solves a).
type naive struct {
	minimum [][]int
}

// Space: O(N^2)
// Time: O(N^2)
func NewNaive(values []int) *naive {
	minimum := make([][]int, len(values))
	for i := range minimum {
		minimum[i] = make([]int, len(values))
		for j := i; j < len(values); j++ {
			minimum[i][j] = math.MaxInt
		}
	}

	for i := range values {
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

// bit solves b) using two binary index trees that are 1-indexed.
type bit struct {
	original []int // 1-indexed copy of the original values
	bit1     []int // 1-indexed binary index tree as used for prefix sum
	bit2     []int // 1-indexed binary index tree which which is a mirror of bit1
}

// Space: O(N)
// Time: O(N)
func NewBit(values []int) *bit {
	original := make([]int, len(values)+1)
	copy(original[1:], values)

	bit1 := make([]int, len(values)+1)
	for i := range bit1 {
		bit1[i] = math.MaxInt
	}

	// bottom-up initialization of the bit
	// move bit value to immediate parent leads to O(N)
	for i, v := range values {
		idx := i + 1
		bit1[idx] = min(bit1[idx], v)

		if r := idx + lsb(idx); r < len(bit1) {
			bit1[r] = min(bit1[r], bit1[idx])
		}
	}

	bit2 := make([]int, len(values)+1)
	for i := range bit2 {
		bit2[i] = math.MaxInt
	}

	// bottom-up initialization of the bit
	// move bit value to immediate parent leads to O(N)
	for i := len(values) - 1; i >= 0; i-- {
		v := values[i]
		idx := i + 1
		bit2[idx] = min(bit2[idx], v)

		if r := idx - lsb(idx); r >= 0 {
			bit2[r] = min(bit2[r], bit2[idx])
		}
	}

	return &bit{original: original, bit1: bit1, bit2: bit2}
}

// Min returns the minimum value of the corresponding input values within the inclusive range of i
// to j.
// Time: O(log N)
func (b *bit) Min(i, j int) int {
	result := math.MaxInt

	for k := i + 1; k <= j+1; k = k + lsb(k) {
		lowerBoundBit1 := k - lsb(k) + 1
		upperBoundBit2 := k + lsb(k) - 1
		// compare to value that is in range i-j which is either bit1, bit2 or original
		if lowerBoundBit1 < i {
			if upperBoundBit2 > j {
				result = min(result, b.original[k])
			} else {
				result = min(result, b.bit2[k])
			}
		} else {
			result = min(result, b.bit1[k])
		}
	}

	for k := j + 1; k >= i+1; k = k - lsb(k) {
		lowerBoundBit1 := k - lsb(k) + 1
		upperBoundBit2 := k + lsb(k) - 1
		// compare to value that is in range i-j which is either bit1, bit2 or original
		if lowerBoundBit1 < i {
			if upperBoundBit2 > j {
				result = min(result, b.original[k])
			} else {
				result = min(result, b.bit2[k])
			}
		} else {
			result = min(result, b.bit1[k])
		}
	}

	return result
}

func lsb(v int) int {
	return v & -v
}
