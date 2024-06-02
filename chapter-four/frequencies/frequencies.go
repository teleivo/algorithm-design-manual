// Package set solves exercise 4.9 4-11.
package frequencies

// majorityElementUsinMap returns the element that appears more than n/2 times.
// Time: O(N)
// Space: O(N)
func majorityElementUsinMap(in []int) (int, bool) {
	frequencies := make(map[int]int)
	for _, num := range in {
		frequencies[num]++
	}

	for num, frequency := range frequencies {
		if frequency > len(in)/2 {
			return num, true
		}
	}

	return 0, false
}
