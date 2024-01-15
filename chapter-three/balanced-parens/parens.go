// Package parens checks if parenthesis in a given input are balanced.
package parens

// isBalanced solves exercise 3.10 3-1.
func isBalanced(in string) bool {
	var openParens int

	for i := 0; i < len(in); i++ {
		if in[i] == '(' {
			openParens++
			continue
		}

		// closing paren without an opening one
		if openParens == 0 {
			return false
		}

		// consume opening paren for closing one
		openParens--
	}

	// opening paren without an closing one
	return openParens == 0
}
