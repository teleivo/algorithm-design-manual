// Package parens checks if parenthesis in a given input are balanced.
package parens

// isBalanced solves exercise 3.10 3-1.
func isBalanced(in string) (int, bool) {
	var openParens int

	for i := 0; i < len(in); i++ {
		if in[i] == '(' {
			openParens++
			continue
		}

		// closing paren without an opening one
		if openParens == 0 {
			return i, false
		}

		// consume opening paren for closing one
		openParens--
	}

	// opening paren without an closing one
	return 0, openParens == 0
}

// countLongestBalanced solves exercise 3.10 3-2.
// Note: I am confused by the hint which to me seems to contradict the problem statement. Finding
// the longest balanced parenthesis while it is not necessarily a contiguous run of parenthesis does
// not make sense to me.
func countLongestBalanced(in string) int {
	var count int
	var openParens int

	for i := 0; i < len(in); i++ {
		if in[i] == '(' {
			openParens++
			continue
		}

		// closing paren without an opening one
		if openParens == 0 {
			continue
		}

		// consume opening paren for closing one
		openParens--
		// count both parens
		count += 2
	}

	return count
}
