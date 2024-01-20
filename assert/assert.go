package assert

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func NoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("expected no error, instead got %v", err)
	}
}

func False(t *testing.T, method string, got bool) {
	t.Helper()

	if got {
		t.Errorf("stack.%s() = %t want %t instead", method, got, false)
	}
}

func True(t *testing.T, method string, got bool) {
	t.Helper()

	if !got {
		t.Errorf("stack.%s() = %t want %t instead", method, got, true)
	}
}

func Equals(t *testing.T, method string, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("stack.%s() = %d want %d instead", method, got, want)
	}
}

func EqualValues(t *testing.T, method string, in, got, want any) {
	t.Helper()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("%s(%v) mismatch (-want +got):\n%s", method, in, diff)
	}
}
