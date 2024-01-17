// Package dynamic implements a dynamic array growing and shrinking depending on the length and
// capacity.
package dynamic

// array solves exercise 3.10 3-5.
// a) I could end up shrinking and doubling the dynamic array repeatedly if I were to double the
// array at full capacity and shrink when the array is just below 50%. For example if I repeatedly
// insert and delete at full capacity.
// b) By decoupling the doubling amount and the shrinking point I can prevent the above. For example
// grow by 50% at full capacity and shrink by 50% at 25% capacity. An array with for example
// capacity of 32 would then shrink to 16 only when at length 8. This means that after doubling it
// from 16 it can grow by 15 before doubling again and shrink by 9 before halving again.
// I could also add a threshold from which halving occurs as the amount of work might not be worth
// the saving of space at low capacities.
type array struct {
	vals []int
}

func (a *array) insert(i, val int) {
	if i < len(a.vals) {
		a.vals[i] = val
		return
	}

	n := i + 1
	if n > cap(a.vals) {
		newCap := max(cap(a.vals)*2, n)
		newVals := make([]int, newCap)
		copy(newVals, a.vals)
		a.vals = newVals
	}
	a.vals = a.vals[:n]
	a.vals[i] = val
}

// get element at index i.
func (a *array) get(i int) int {
	return a.vals[i]
}

// delete last element.
func (a *array) delete() int {
	n := len(a.vals) - 1
	if n < 0 {
		panic("empty array: no element left to delete")
	}
	val := a.vals[n]

	if a.len() == a.cap()/4 {
		newVals := make([]int, a.cap()/2)
		copy(newVals, a.vals)
		a.vals = newVals
	}
	a.vals = a.vals[:n]

	return val
}

// len returns the current length of the array.
func (a *array) len() int {
	return len(a.vals)
}

// len returns the current capacity of the array.
func (a *array) cap() int {
	return cap(a.vals)
}
