// Package partialsum solves exercise 3.10 3-28. It uses a Fenwick tree
// https://cp-algorithms.com/data_structures/fenwick.html
//
// Package partialsum can be extended to solve exercise 3.10 3-29. Its not super clear to me what
// constraints I have to adhere to with regards to the exercise stating "extends 3-28".
// https://algorist.com/algowiki/index.php/3.29 for example does not seem to extend 3-28 as there we
// are only allowed one array of O(N) while the solution uses a balanced BST. I am not sure a
// balanced BST can be implemented implicitly in an array while maintaining O(log N) on all the
// requested operations.
//
// I would extend the Fenwick tree solution in the following way
//
// Insert(k, v) is essentially Add(k, v) using `=` instead of `+=`.
// Delete(k) can be done by Add(k, -original[k]).
// The array could contain many unused elements over time. We could compact the array like in a
// dynamic array implementation and amortize the O(N) Fenwick tree construction.
//
// Providing the key k in all signatures could be done by translating keys k to Fenwick tree indices
// i via a BST. A balanced BST provides insert, update, delete in O(log N). One could then argue why
// not go for the balanced BST to provide all of the functionality. Good point :)

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
