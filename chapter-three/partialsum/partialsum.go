// Package partialsum solves exercise 3.10 3-28. It uses a Fenwick tree
// https://cp-algorithms.com/data_structures/fenwick.html
package partialsum

type Partial struct {
	t []int
}

// New creates a new Partial. Uses the naive way to initialize the Fenwick tree.
// Time: O(N log N) can be reduced to O(log N)
// Space: O(N)
func New(vals ...int) *Partial {
	p := &Partial{t: make([]int, len(vals)+1)}
	for i, v := range vals {
		p.Add(i, v)
	}

	return p
}

// Sum cumulatively adds all values up to and including index of i.
func (p *Partial) Sum(i int) int {
	var result int

	for idx := i + 1; idx > 0; idx -= idx & -idx {
		result += p.t[idx]
	}

	return result
}

// Add adds given value to the value at index of i.
func (p *Partial) Add(i, val int) {
	for idx := i + 1; idx < len(p.t); idx += idx & -idx {
		p.t[idx] += val
	}
}
