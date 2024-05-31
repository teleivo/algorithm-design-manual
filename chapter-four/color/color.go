// Package set solves exercise 4.9 4-4.
package color

// sort sorts given sequence of number pairs by the pairs second number while keeping their relative
// order stable. Note that I replaced the colors red, blue, yellow with numbers 1, 2, 3 for
// simplicity.
// Time: O(N)
// Space: O(N) - I think an inplace version cannot be done in O(N) due to the stable constraint. I
// would need to shift elements which would make it O(N^2). If the input where a linked list
// in-place could be done in O(N) time.
func sort(in [][2]int) [][2]int {
	result := make([][2]int, 0, len(in))

	for i := 1; i < 4; i++ { // collect elements in order of second pairs key i.e. color
		for _, v := range in {
			if v[1] == i {
				result = append(result, v)
			}
		}
	}

	return result
}
