package anagram

import (
	"testing"

	"github.com/teleivo/assertive/assert"
)

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		a    string
		b    string
		want bool
	}{
		{
			a:    "silent",
			b:    "silent",
			want: true,
		},
		{
			a:    "silent",
			b:    "listen",
			want: true,
		},
		{
			a:    "listen",
			b:    "silent",
			want: true,
		},
		{
			a:    "silend",
			b:    "listen",
			want: false,
		},
		{
			a:    "silence",
			b:    "listening",
			want: false,
		},
		{
			a:    "aaabbc",
			b:    "aabccc",
			want: false,
		},
	}

	for _, tc := range tests {
		got := isAnagram(tc.a, tc.b)

		t.Logf("%q %q", tc.a, tc.b)
		assert.Equals(t, got, tc.want)
	}
}
