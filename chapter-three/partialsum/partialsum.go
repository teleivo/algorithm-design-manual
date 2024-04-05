// Package partialsum solves exercise 3.10 3-28. It uses a Fenwick tree
// https://cp-algorithms.com/data_structures/fenwick.html
package partialsum

type Partial struct {
	bit []int
}

// New creates a new Partial. Uses the naive way to initialize the Fenwick tree.
// Time: O(log N)
// Space: O(N)
func New(vals ...int) *Partial {
	bit := make([]int, len(vals)+1) // using 1-indexed fenwick tree
	for i, v := range vals {
		idx := i + 1
		bit[idx] += v
		if r := idx + lowestOneBit(idx); r < len(vals)+1 {
			bit[r] += bit[idx]
		}
	}

	return &Partial{bit: bit}
}

// Sum cumulatively adds all values up to and including index of i.
func (p *Partial) Sum(i int) int {
	var result int

	for idx := i + 1; idx > 0; idx -= lowestOneBit(idx) {
		result += p.bit[idx]
	}

	return result
}

// Add adds given value to the value at index of i.
func (p *Partial) Add(i, val int) {
	for idx := i + 1; idx < len(p.bit); idx += lowestOneBit(idx) {
		p.bit[idx] += val
	}
}

// lowestOneBit returns the lowest bit that is set to one.
func lowestOneBit(v int) int {
	return v & -v
}
