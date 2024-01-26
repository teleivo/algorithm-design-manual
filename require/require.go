package require

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/teleivo/algorithm-design-manual/test/report"
)

func NoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("expected no error, instead got %v", err)
	}
}

func False(t *testing.T, got bool) {
	t.Helper()

	report.False(t.Fatalf, got)
}

func True(t *testing.T, got bool) {
	t.Helper()

	report.True(t.Fatalf, got)
}

func Nil(t *testing.T, method string, got any) {
	t.Helper()

	if got != nil {
		t.Fatalf("stack.%s() = %v want nil instead", method, got)
	}
}

func Equals(t *testing.T, method string, got, want any) {
	t.Helper()

	if got != want {
		t.Fatalf("stack.%s() = %v want %v instead", method, got, want)
	}
}

func EqualValues(t *testing.T, method string, in, got, want any) {
	t.Helper()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Fatalf("%s(%v) mismatch (-want +got):\n%s", method, in, diff)
	}
}
