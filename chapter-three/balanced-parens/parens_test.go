package parens

import "testing"

func TestBalancedParens(t *testing.T) {

	tests := []struct {
		in   string
		want bool
	}{
		{
			in:   "()",
			want: true,
		},
		{
			in:   "((())())()",
			want: true,
		},
		{
			in:   ")",
			want: false,
		},
		{
			in:   "(",
			want: false,
		},
		{
			in:   "())",
			want: false,
		},
		{
			in:   ")()(",
			want: false,
		},
	}

	for _, tc := range tests {
		got := isBalanced(tc.in)

		if got != tc.want {
			t.Errorf("isBalanced(%q) = %t want %t", tc.in, got, tc.want)
		}
	}
}
