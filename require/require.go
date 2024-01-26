package require

import (
	"testing"

	"github.com/teleivo/algorithm-design-manual/test/report"
)

func NoError(t *testing.T, err error) {
	t.Helper()

	report.NoError(t.Fatalf, err)
}

func False(t *testing.T, got bool) {
	t.Helper()

	report.False(t.Fatalf, got)
}

func True(t *testing.T, got bool) {
	t.Helper()

	report.True(t.Fatalf, got)
}

func Nil(t *testing.T, got any) {
	t.Helper()

	report.Nil(t.Fatalf, got)
}

func Equals(t *testing.T, method string, got, want any) {
	t.Helper()

	report.Equals(t.Fatalf, got, want)
}

func EqualValues(t *testing.T, method string, in, got, want any) {
	t.Helper()

	report.EqualValues(t.Fatalf, got, want)
}
