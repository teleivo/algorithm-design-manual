// Package rangeminimum solves exercise 3.10 3-30. The solution is using a Fenwick tree for the
// prefix sum of available rooms. It is also using Fenwick trees as discussed in the paper
// "Efficient Range Minimum Queries using Binary Indexed Trees" Mircea DIMA, Rodica CETERCHI
// https://ioinformatics.org/journal/v9_2015_39_44.pdf
package hotel

import (
	"fmt"
	"math"
)

type rooms struct {
	occupied       []int // indicates if the room at given index is occupied via MaxInt
	availableCount []int // Fenwick tree storing prefix sums to provide Count(l, h) of available rooms
	checkin1       []int // Fenwick tree storing min range to provide Checkin(l, h)
	checkin2       []int // Fenwick tree storing min range to provide Checkin(l, h) mirror of checkin2
}

// Time: O(N)
// Space: O(N)
func New(n int) *rooms {
	n++ // using 1-indexed Fenwick trees

	occupied := make([]int, n)
	for i := range occupied {
		occupied[i] = i
	}

	availableCount := make([]int, n)
	for i := 1; i < n; i++ {
		availableCount[i]++
		idx := i + lsb(i)
		if idx < n {
			availableCount[idx] += availableCount[i]
		}
	}

	checkin1 := make([]int, n)
	for i := range checkin1 {
		checkin1[i] = math.MaxInt
	}
	// bottom-up initialization of the bit
	// move bit value to immediate parent leads to O(N)
	for i := 1; i < n; i++ {
		checkin1[i] = min(checkin1[i], i)
		if j := i + lsb(i); j < n {
			checkin1[j] = min(checkin1[j], checkin1[i])
		}
	}

	checkin2 := make([]int, n)
	for i := range checkin2 {
		checkin2[i] = math.MaxInt
	}
	// bottom-up initialization of the bit
	// move bit value to immediate parent leads to O(N)
	for i := n - 1; i > 0; i-- {
		checkin2[i] = min(checkin2[i], i)
		if j := i - lsb(i); j > 0 {
			checkin2[j] = min(checkin2[j], checkin2[i])
		}
	}

	return &rooms{
		occupied:       occupied,
		availableCount: availableCount,
		checkin1:       checkin1,
		checkin2:       checkin2,
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
		result += r.availableCount[idx]
	}

	return result
}

// Checkin returns the first available room in inclusive range l to h and marks the room as
// occupied. Returns -1 if all rooms in given range are occupied.
// Time: O(log^2 N)
func (r *rooms) Checkin(l, h int) int {
	fmt.Printf("Checkin(%d, %d) rooms before %#v\n", l, h, r)
	room := r.minRange(l, h)
	if room == math.MaxInt {
		return room
	}

	r.occupy(room)
	fmt.Printf("Checkin(%d, %d) rooms after %#v\n", l, h, r)
	return room
}

func (r *rooms) minRange(l, h int) int {
	room := math.MaxInt

	for idx := l; idx <= h; idx = idx + lsb(idx) {
		lowerBoundCheckin1 := idx - lsb(idx) + 1
		upperBoundCheckin2 := idx + lsb(idx) - 1
		// compare to value that is in range i-j which is either bit1, bit2 or original
		if lowerBoundCheckin1 < l {
			if upperBoundCheckin2 > h {
				room = min(room, r.occupied[idx])
			} else {
				room = min(room, r.checkin2[idx])
			}
		} else {
			room = min(room, r.checkin1[idx])
		}
	}

	for idx := h; idx >= l; idx = idx - lsb(idx) {
		lowerBoundCheckin1 := idx - lsb(idx) + 1
		upperBoundCheckin2 := idx + lsb(idx) - 1
		// compare to value that is in range i-j which is either bit1, bit2 or original
		if lowerBoundCheckin1 < l {
			if upperBoundCheckin2 > h {
				room = min(room, r.occupied[idx])
			} else {
				room = min(room, r.checkin2[idx])
			}
		} else {
			room = min(room, r.checkin1[idx])
		}
	}

	fmt.Printf("minRange(%d, %d) = %d\n", l, h, room)
	return room
}

func (r *rooms) occupy(l int) {
	r.update(l, math.MaxInt)
}

// Time: O(log N)
func (r *rooms) update(l, value int) {
	n := len(r.occupied)
	for idx := l; idx < n; idx = idx + lsb(idx) {
		if value == math.MaxInt {
			r.availableCount[idx]--
		} else {
			r.availableCount[idx]++
		}
	}

	r.occupied[l] = value

	// checkin1/bit1 update
	currentMin := math.MaxInt
	for idx := l; idx < n; idx = idx + lsb(idx) { // climb bit1
		// get min in range [prevIdx,idx-1] if that range is in the range of idx [idx-lsb(idx)+1, idx]
		for idxLower := idx - 1; idxLower >= idx-lsb(idx)+1 && idxLower > 0; idxLower = idxLower - lsb(idxLower) { // climb bit2
			currentMin = min(currentMin, r.checkin1[idxLower])
		}

		currentMin = min(currentMin, r.occupied[idx]) // include actual value into current range of min
		r.checkin1[idx] = currentMin
	}

	// checkin2/bit2 update (inverse of bit1 update)
	currentMin = math.MaxInt
	for idx := l; idx > 0; idx = idx - lsb(idx) {
		for idxUpper := idx + 1; idxUpper <= idx+lsb(idx)-1 && idxUpper < n; idxUpper = idxUpper + lsb(idxUpper) { // climb bit1
			currentMin = min(currentMin, r.checkin2[idxUpper])
		}

		currentMin = min(currentMin, r.occupied[idx]) // include actual value into current range of min
		r.checkin2[idx] = currentMin
	}
}

func (r *rooms) Checkout(l int) {
	r.update(l, l)
}

// lsb returns the least significant bit.
func lsb(i int) int {
	return i & -i
}
