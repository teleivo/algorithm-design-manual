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
// Time: O(log N)
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
	// fmt.Printf("minRange(%d, %d) rooms: %#v\n", l, h, r)

	for idx := l; idx <= h; idx = idx + lsb(idx) {
		lowerBoundCheckin1 := idx - lsb(idx) + 1
		upperBoundCheckin2 := idx + lsb(idx) - 1
		// compare to value that is in range i-j which is either bit1, bit2 or original
		if lowerBoundCheckin1 < l {
			if upperBoundCheckin2 > h {
				room = min(room, idx)
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
				room = min(room, idx)
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

// TODO implement
func (r *rooms) update(l, value int) {
	n := len(r.occupied)
	for idx := l; idx < n; idx = idx + lsb(idx) {
		if value == math.MaxInt {
			r.availableCount[idx]--
		} else {
			r.availableCount[idx]++
		}
	}

	oldValue := r.occupied[l]
	r.occupied[l] = value

	for idx := l; idx < n; idx = idx + lsb(idx) {
		if r.checkin1[idx] != oldValue {
			r.checkin1[idx] = min(r.checkin1[idx], value)
		} else {
			x, y := idx-lsb(idx)+1, idx
			fmt.Printf("update(%d) tree1\n", idx)
			newMin := value
			if x <= l-1 {
				newMin = min(newMin, r.minRange(x, l-1))
			}
			// TODO Add && l+1 > 0?
			if l+1 <= y {
				newMin = min(newMin, r.minRange(l+1, y))
			}
			r.checkin1[idx] = newMin
		}
	}

	for idx := l; idx > 0; idx = idx - lsb(idx) {
		if r.checkin2[idx] != oldValue {
			r.checkin2[idx] = min(r.checkin2[idx], value)
		} else {
			x, y := idx, idx+lsb(idx)-1
			fmt.Printf("update(%d) tree2\n", idx)
			newMin := value
			if x <= idx-1 {
				newMin = min(newMin, r.minRange(x, idx-1))
			}
			if idx+1 <= y && y < n {
				newMin = min(newMin, r.minRange(idx+1, y))
			}
			r.checkin2[idx] = newMin
		}
	}
}

func (r *rooms) Checkout(l int) {
	r.update(l, l)
}

// lsb returns the least significant bit.
func lsb(i int) int {
	return i & -i
}
