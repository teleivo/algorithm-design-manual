// Package set solves exercise 4.9 4-2.
package minmax

import (
	"math"
	"slices"
)

// maximizeUnsorted finds the pair of numbers most distant from each other.
// Time: O(N)
// Space: O(1)
func maximizeUnsorted(in []int) (int, int) {
	if len(in) < 2 {
		panic("input must have at least 2 elements")
	}

	minVal := math.MaxInt
	maxVal := math.MinInt
	for _, v := range in {
		minVal = min(minVal, v)
		maxVal = max(maxVal, v)
	}

	return minVal, maxVal
}

// maximizeSorted finds the pair of numbers most distant from each other.
// Time: O(1)
// Space: O(1)
func maximizeSorted(in []int) (int, int) {
	if len(in) < 2 {
		panic("input must have at least 2 elements")
	}

	return in[0], in[len(in)-1]
}

// minimizeUnsorted finds the pair of distinct numbers closest to each other.
// Time: O(N log N)
// Space: O(N) - could be done in place but I am using space to not mutate the input
func minimizeUnsorted(in []int) (int, int) {
	if len(in) < 2 {
		panic("input must have at least 2 elements")
	}

	sortedIn := make([]int, len(in))
	copy(sortedIn, in)
	slices.Sort(sortedIn)

	return minimizeSorted(sortedIn)
}

// minimizeSorted finds the pair of distinct numbers closest to each other.
// Time: O(N)
// Space: O(1)
func minimizeSorted(in []int) (int, int) {
	if len(in) < 2 {
		panic("input must have at least 2 elements")
	}

	firstIndex, secondIndex := 0, 1
	for i, j, diff := 1, 2, absDistance(in[firstIndex], in[secondIndex]); i < len(in)-1; i, j = i+1, j+1 {
		if absDistance(in[i], in[j]) < diff {
			firstIndex, secondIndex = i, j
			diff = absDistance(in[i], in[j])
		}
	}

	return in[firstIndex], in[secondIndex]
}

func absDistance(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
