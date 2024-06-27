// Package set solves exercise 4.9 4-12.
package union

import (
	"slices"
)

// unionUnsorted returns the union of two sets.
// Time: O(N log N) - sorting the slices is O(N log N)
// Space: O(1) - sorting a and b in place here as I don't assert on them, could obviously also copy
// them.
func unionUnsorted(a, b []int) []int {
	slices.Sort(a)
	slices.Sort(b)
	return unionSorted(a, b)
}

// unionSorted returns the union of two sets.
// Time: O(N)
// Space: O(1)
func unionSorted(a, b []int) []int {
	var out []int
	for i, j := 0, 0; i < len(a) || j < len(b); {
		if i == len(a) {
			out = append(out, b[j:]...)
			break
		} else if j == len(b) {
			out = append(out, a[i:]...)
			break
		}

		if a[i] == b[j] {
			out = append(out, a[i])
			i++
			j++
		} else if a[i] < b[j] {
			out = append(out, a[i])
			i++
		} else {
			out = append(out, b[j])
			j++
		}
	}
	return out
}
