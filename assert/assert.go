package assert

import "testing"

func NoError(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("expected no error, instead got %v", err)
	}
}

func False(t *testing.T, method string, got bool) {
	if got {
		t.Fatalf("stack.%s() = %t want %t instead", method, got, false)
	}
}
func True(t *testing.T, method string, got bool) {
	if !got {
		t.Fatalf("stack.%s() = %t want %t instead", method, got, true)
	}
}

func Equals(t *testing.T, method string, got, want int) {
	if got != want {
		t.Fatalf("stack.%s() = %d want %d instead", method, got, want)
	}
}
