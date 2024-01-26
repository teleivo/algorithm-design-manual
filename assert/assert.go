package assert

import (
	"testing"

	"github.com/teleivo/algorithm-design-manual/test/report"
)

func NoError(t *testing.T, err error) {
	t.Helper()

	report.NoError(t.Errorf, err)
}

func False(t *testing.T, method string, got bool) {
	t.Helper()

	report.False(t.Errorf, got)
}

func True(t *testing.T, method string, got bool) {
	t.Helper()

	report.True(t.Errorf, got)
}

func Nil(t *testing.T, got any) {
	t.Helper()

	report.Nil(t.Errorf, got)
}

func Equals(t *testing.T, method string, got, want int) {
	t.Helper()

	report.Equals(t.Errorf, got, want)
}

func EqualValues(t *testing.T, method string, in, got, want any) {
	t.Helper()

	report.EqualValues(t.Errorf, got, want)
}
