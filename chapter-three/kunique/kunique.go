// Package kunique solves exercise 3.10 3-24.
package kunique

// IsKUnique returns true if there are no duplicates within k-positions from each other and false
// otherwise.
// IsKUnique soles exercise 3.10 3-24.
// Time: O(N)
// Space: O(N)
func IsKUnique(k int, a []int) bool {
	if len(a) == 0 {
		return true
	}

	seen := make(map[int]int)
	for i, v := range a {
		if j, ok := seen[v]; ok && i-j <= k {
			return false
		}
		seen[v] = i
	}

	return true
}
