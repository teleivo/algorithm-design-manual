package assert

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/teleivo/algorithm-design-manual/test/report"
)

func NoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("expected no error, instead got %v", err)
	}
}

func False(t *testing.T, method string, got bool) {
	t.Helper()

	report.False(t.Errorf, got)
}

func True(t *testing.T, method string, got bool) {
	t.Helper()

	report.True(t.Errorf, got)
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
