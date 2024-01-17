package parens

import "testing"

func TestBalancedParens(t *testing.T) {
	tests := []struct {
		in      string
		wantPos int
		wantOk  bool
	}{
		{
			in:      "()",
			wantPos: 0,
			wantOk:  true,
		},
		{
			in:      "((())())()",
			wantPos: 0,
			wantOk:  true,
		},
		{
			in:      ")",
			wantPos: 0,
			wantOk:  false,
		},
		{
			in:      "(",
			wantPos: 0,
			wantOk:  false,
		},
		{
			in:      "())",
			wantPos: 2,
			wantOk:  false,
		},
		{
			in:      ")()(",
			wantPos: 0,
			wantOk:  false,
		},
	}

	for _, tc := range tests {
		gotPos, gotOk := isBalanced(tc.in)

		if gotPos != tc.wantPos || gotOk != tc.wantOk {
			t.Errorf("isBalanced(%q) = (%d, %t) want (%d, %t)", tc.in, gotPos, gotOk, tc.wantPos, tc.wantOk)
		}
	}
}

func TestCountLongestBalanced(t *testing.T) {
	tests := []struct {
		in     string
		want   int
		wantOk bool
	}{
		{
			in:   "",
			want: 0,
		},
		{
			in:   "()",
			want: 2,
		},
		{
			in:   "((())())()",
			want: 10,
		},
		{
			in:   ")",
			want: 0,
		},
		{
			in:   "(",
			want: 0,
		},
		{
			in:   "())",
			want: 2,
		},
		{
			in:   ")()(",
			want: 2,
		},
		{
			in:   ")()(())()()))())))(",
			want: 12,
		},
	}

	for _, tc := range tests {
		got := countLongestBalanced(tc.in)

		if got != tc.want {
			t.Errorf("countLongestBalanced(%q) = %d want %d", tc.in, got, tc.want)
		}
	}
}
