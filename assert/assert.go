package assert

import (
	"testing"

	"github.com/teleivo/algorithm-design-manual/test/report"
)

func NoError(t *testing.T, err error) {
	t.Helper()

	report.NoError(t, t.Errorf, err)
}

func False(t *testing.T, got bool) {
	t.Helper()

	report.False(t, t.Errorf, got)
}

func True(t *testing.T, got bool) {
	t.Helper()

	report.True(t, t.Errorf, got)
}

func Nil(t *testing.T, got any) {
	t.Helper()

	report.Nil(t, t.Errorf, got)
}

func Equals(t *testing.T, got, want any) {
	t.Helper()

	report.Equals(t, t.Errorf, got, want)
}

func EqualValues(t *testing.T, got, want any) {
	t.Helper()

	report.EqualValues(t, t.Errorf, got, want)
}
