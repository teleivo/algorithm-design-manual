// Package set solves exercise 4.9 4-9.
package nsum

import "slices"

// exists returns true if there are k elements in a adding up to sum and false otherwise.
// Time: O(N log N) - sorting the slice is O(N log N) + worst-case doing a binary search N^k-1 times
// for the other operands
// Space: O(1) - sorting a in place here as I don't assert on them, could obviously also copy them.
func exists(k int, a []int, sum int) bool {
	slices.Sort(a)
	return find(k, a, sum)
}

func find(k int, a []int, sum int) bool {
	if k == 1 {
		_, ok := slices.BinarySearch(a, sum)
		return ok
	}
	k--
	for i := 0; i < len(a)-1; i++ {
		if find(k, a[i+1:], sum-a[i]) {
			return true
		}
	}
	return false
}
