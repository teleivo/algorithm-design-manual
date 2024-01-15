// Package parens checks if parenthesis in a given input are balanced.
package parens

// isBalanced solves exercise 3.10 3-1.
func isBalanced(in string) bool {
	openParens := make([]byte, 0, len(in))

	for i := 0; i < len(in); i++ {
		if in[i] == '(' {
			openParens = append(openParens, '(')
			continue
		}

		// closing paren without an opening one
		if len(openParens) == 0 {
			return false
		}

		// consume opening paren for closing one
		openParens = openParens[:len(openParens)-1]
	}

	// opening paren without an closing one
	return len(openParens) == 0
}
