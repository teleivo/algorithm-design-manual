// Package set solves exercise 4.9 4-13.
// As described in the exercise all entry and exit times are distinct i.e. no interval has equal
// start and end values and no 2 intervals have the same start or end value.
package party

type interval [2]int

func (a interval) Intersect(b interval) bool {
	return (a[0] < b[0] && b[0] < a[1]) || (a[0] < b[1] && b[1] < a[1])
}
