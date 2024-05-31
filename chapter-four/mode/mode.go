// Package set solves exercise 4.9 4-5.
package mode

import (
	"math"
	"slices"
)

// Time: O(N)
func modeUsingMap(in []int) int {
	frequencies := make(map[int]int)
	mode := math.MinInt
	curMax := math.MinInt
	for _, num := range in {
		frequencies[num]++
		if frequencies[num] > curMax {
			mode = num
			curMax = frequencies[num]
		}
	}
	return mode
}

// Time: O(N log N)
// Space: O(N) - am making a copy for ease of testing, I could also do this in-place
func modeUsingSort(in []int) int {
	sortedIn := make([]int, len(in))
	copy(sortedIn, in)
	slices.Sort(sortedIn)

	mode := math.MinInt
	for i, curFrequency, maxFrequency := 1, 1, 1; i < len(sortedIn)-1; i++ {
		if sortedIn[i-1] == sortedIn[i] {
			curFrequency++
		} else {
			curFrequency = 1
		}

		if curFrequency > maxFrequency {
			maxFrequency = curFrequency
			mode = sortedIn[i]
		}
	}

	return mode
}
