// Package set solves exercise 4.9 4-10.
package twosum

import "slices"

// existsUsingMap  returns true if there is a pair adding up to x and false otherwise.
// x and false otherwise.
// Package set solves exercise 4.9 4-10 b.
// Time: O(N)
// Space: O(N)
func existsUsingMap(a []int, x int) bool {
	sum := make(map[int]struct{}, len(a))

	for _, num := range a {
		if _, ok := sum[x-num]; ok {
			return true
		}
		sum[num] = struct{}{}
	}
	return false
}

// existsUsingSort returns true if there is a pair adding up to x and false otherwise.
// Package set solves exercise 4.9 4-10 a.
// Time: O(N log N) - sorting is O(N log N) + worst-case doing a binary search N
// times for the other operand
// Space: O(1) - sorting a in place here as I don't assert on them, could obviously also copy them.
func existsUsingSort(a []int, x int) bool {
	slices.Sort(a)

	for i, num := range a {
		_, ok := slices.BinarySearch(a[i+1:], x-num)
		if ok {
			return true
		}
	}
	return false
}
