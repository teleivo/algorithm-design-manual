// Package set solves exercise 4.9 4-6.
package twosum

import "slices"

// existsUsingMap returns true if there is a pair with one element in each of a and b adding up to
// x and false otherwise.
// Time: O(N)
// Space: O(N)
func existsUsingMap(a, b []int, x int) bool {
	sum := make(map[int]struct{}, len(a))

	for _, num := range b {
		sum[num] = struct{}{}
	}

	for _, num := range a {
		if _, ok := sum[x-num]; ok {
			return true
		}
	}
	return false
}

// existsUsingSort returns true if there is a pair with one element in each of a and b adding up to
// x and false otherwise.
// Time: O(N log N) - sorting both slices is O(N log N) each + worst-case doing a binary search N
// times for the other operand
// Space: O(1) - sorting a and b in place here as I don't assert on them, could obviously also copy
// them.
func existsUsingSort(a, b []int, x int) bool {
	slices.Sort(a)
	slices.Sort(b)

	for _, num := range a {
		_, ok := slices.BinarySearch(b, x-num)
		if ok {
			return true
		}
	}
	return false
}
