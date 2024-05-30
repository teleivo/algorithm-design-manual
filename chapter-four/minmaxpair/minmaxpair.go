// Package set solves exercise 4.9 4-3.
package minmaxpair

import (
	"slices"
)

// partition returns pairs by minimizing the maximum sum of a pair.
// Time: O(N log N)
// Space: O(N) - could be done in place but I am using space to not mutate the input
func partition(in []int) [][2]int {
	if len(in) < 4 {
		panic("input must have at least 4 elements")
	}
	if len(in)%2 != 0 {
		panic("input must have an even number of elements")
	}

	sortedIn := make([]int, len(in))
	copy(sortedIn, in)
	slices.Sort(sortedIn)

	pairs := make([][2]int, len(sortedIn)/2)

	for i := 0; i < len(sortedIn)/2; i++ {
		pairs[i] = [2]int{sortedIn[i], sortedIn[len(sortedIn)-i-1]}
	}

	return pairs
}
