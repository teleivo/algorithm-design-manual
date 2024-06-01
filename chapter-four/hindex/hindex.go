// Package set solves exercise 4.9 4-7.
package hindex

import (
	"slices"
)

// hindex computes the h-index of a researchers citations.
// Time: O(N log N)
// Space: O(1) - sorting a and b in place here as I don't assert on them, could obviously also copy
// them.
func hindex(citations []int) int {
	slices.Sort(citations)

	r := 1
	for i := len(citations) - 1; i >= 0; i, r = i-1, r+1 {
		if citations[i] < r {
			break
		}
	}

	return r - 1
}
