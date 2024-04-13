package hotel

import (
	"math"
)

type rooms struct {
	occupied       []bool // indicates if the room at given index is available
	availableCount []int  // Fenwick tree storing prefix sums to provide Count(l, h) of available rooms
	checkin1       []int  // Fenwick tree storing min range to provide Checkin(l, h)
	checkin2       []int  // Fenwick tree storing min range to provide Checkin(l, h) mirror of checkin2
}

// Time: O(N)
// Space: O(N)
func New(n int) *rooms {
	n++ // using 1-indexed Fenwick trees

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
		idx := i + lsb(i)
		if idx < n {
			checkin1[idx] = min(checkin1[idx], checkin1[i])
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
		idx := i - lsb(i)
		if idx > 0 {
			checkin2[idx] = min(checkin2[idx], checkin1[i])
		}
	}

	return &rooms{
		occupied:       make([]bool, n),
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
// Time: O(log N)
func (r *rooms) Checkin(l, h int) int {
	room := r.minRange(l, h)
	if room == math.MaxInt {
		return room
	}
	r.occupy(room)
	return room
}

func (r *rooms) minRange(l, h int) int {
	room := math.MaxInt

	for idx := l; idx <= h; idx = idx + lsb(idx) {
		lowerBoundBit1 := idx - lsb(idx) + 1
		upperBoundBit2 := idx + lsb(idx) - 1
		// compare to value that is in range i-j which is either bit1, bit2 or original
		if lowerBoundBit1 < l {
			if upperBoundBit2 > h {
				room = min(room, idx)
			} else {
				room = min(room, r.checkin2[idx])
			}
		} else {
			room = min(room, r.checkin1[idx])
		}
	}

	for idx := h; idx >= l; idx = idx - lsb(idx) {
		lowerBoundBit1 := idx - lsb(idx) + 1
		upperBoundBit2 := idx + lsb(idx) - 1
		// compare to value that is in range i-j which is either bit1, bit2 or original
		if lowerBoundBit1 < l {
			if upperBoundBit2 > h {
				room = min(room, idx)
			} else {
				room = min(room, r.checkin2[idx])
			}
		} else {
			room = min(room, r.checkin1[idx])
		}
	}

	return room
}

func (r *rooms) occupy(l int) {
	r.occupied[l] = true

	for idx := l; idx < len(r.availableCount); idx = idx + lsb(idx) {
		r.availableCount[idx]--
	}

	for idx := l; idx < len(r.checkin1); idx = idx + lsb(idx) {
		if idx == l {
			if r.checkin1[idx] == l {
				// make the room itself unavailable if the ranges' min is set to the room itself
				// otherwise it has to point to a lower room number and thus keep it as is
				r.checkin1[idx] = math.MaxInt
			} else {
				break
			}
		} else {
			if r.occupied[idx] {
				r.checkin1[idx] = math.MaxInt
			} else {
				r.checkin1[idx] = min(idx, r.checkin1[idx-lsb(idx)+1])
			}
		}
	}

	for idx := l; idx > 0; idx = idx - lsb(idx) {
		if idx == l {
			if r.checkin2[idx] == l {
				// make the room itself unavailable if the ranges' min is set to the room itself
				// otherwise it has to point to a lower room number and thus keep it as is
				r.checkin2[idx] = math.MaxInt
			} else {
				break
			}
		} else {
			if r.occupied[idx] {
				r.checkin2[idx] = math.MaxInt
			} else {
				r.checkin2[idx] = min(idx, r.checkin2[idx+lsb(idx)-1])
			}
		}
	}
}

// TODO implement
// Time: O(log N)
func (r *rooms) Checkout(l int) int {
	return 0
}

// lsb returns the least significant bit.
func lsb(i int) int {
	return i & -i
}
