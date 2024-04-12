package hotel

import (
	"fmt"
	"math"
)

type rooms struct {
	available []int // Fenwick tree storing prefix sums to provide Count(l, h) of available rooms
	checkin1  []int // Fenwick tree storing min range to provide Checkin(l, h)
	checkin2  []int // Fenwick tree storing min range to provide Checkin(l, h) mirror of checkin2
}

func New(n int) *rooms {
	n++

	available := make([]int, n)
	for i := 1; i < n; i++ {
		available[i]++
		idx := i + lsb(i)
		if idx < n {
			available[idx] += available[i]
		}
	}

	checkin1 := make([]int, n)
	for i := range checkin1 {
		checkin1[i] = math.MaxInt
	}
	for i := 1; i < n; i++ {
		checkin1[i] = min(checkin1[i], i)
		idx := i + lsb(i)
		if idx < n {
			checkin1[idx] = min(checkin1[idx], available[i])
		}
	}

	checkin2 := make([]int, n)
	for i := range checkin2 {
		checkin2[i] = math.MaxInt
	}
	for i := n - 1; i > 0; i-- {
		checkin2[i] = min(checkin2[i], i)
		idx := i - lsb(i)
		if idx > 0 {
			checkin2[idx] = min(checkin2[idx], available[i])
		}
	}

	return &rooms{
		available: available,
		checkin1:  checkin1,
		checkin2:  checkin2,
	}
}

// Count returns the number of available rooms in inclusive rang l to h.
// Time: O(log N)
func (r *rooms) Count(l, h int) int {
	return r.sum(h) - r.sum(l-1)
}

// sum computes the prefix sum of available rooms up to and including l.
// Time: O(log N)
func (r *rooms) sum(l int) int {
	var result int
	for idx := l; idx > 0; idx = idx - lsb(idx) {
		result += r.available[idx]
	}
	fmt.Println(result)
	return result
}

// Checkin returns the first available room in inclusive range l to h and marks the room as
// occupied. Returns -1 if all rooms in given range are occupied.
// Time: O(log N)
func (r *rooms) Checkin(l, h int) int {
	var room int

	// TODO implement range min

	// TODO implement Add(i, i) of the selected room?

	return room
}

// lsb returns the least significant bit.
func lsb(i int) int {
	return i & -i
}
