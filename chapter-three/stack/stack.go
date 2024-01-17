// Package implements a stack with push, pop and findmin operating in O(1).
package stack

import "errors"

// stack solves exercise 3.10 3-4.
type stack struct {
	vals []int
	mins []int
}

func (s *stack) push(x int) {
	s.vals = append(s.vals, x)

	// track min values in their own stack
	// duplicate a min for now as this makes the implementation easier at the cost of space. I could
	// save space by counting how often a particular min appears.
	if len(s.mins) == 0 || x <= s.mins[len(s.mins)-1] {
		s.mins = append(s.mins, x)
	}
}

func (s *stack) pop() (int, error) {
	if s.isEmpty() {
		return 0, errors.New("stack is empty")
	}

	last := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]

	// pop element from mins if it was a min
	if last == s.mins[len(s.mins)-1] {
		s.mins = s.mins[:len(s.mins)-1]
	}

	return last, nil
}

func (s *stack) findmin() (int, error) {
	if s.isEmpty() {
		return 0, errors.New("stack is empty")
	}

	return s.mins[len(s.mins)-1], nil
}

func (s *stack) isEmpty() bool {
	return len(s.vals) == 0
}
