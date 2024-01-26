package report

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

// report notifies a user of a failed assertion. Functions like t.Errorf, t.Fatalf.
type report func(string, ...any)

func NoError(fn report, err error) {
	if err != nil {
		fn(fmt.Sprintf("expected no error, instead got %v", err))
	}
}

func False(fn report, got bool) {
	if got {
		fn(fmt.Sprintf("got %t want %t instead", got, false))
	}
}

func True(fn report, got bool) {
	if !got {
		fn(fmt.Sprintf("got %t want %t instead", got, true))
	}
}

func Nil(fn report, got any) {
	if got != nil {
		fn(fmt.Sprintf("got %v want nil instead", got))
	}
}

func Equals(fn report, got, want any) {
	if got != want {
		fn(fmt.Sprintf("got %v want %v instead", got, want))
	}
}

func EqualValues(fn report, got, want any) {
	if diff := cmp.Diff(want, got); diff != "" {
		fn(fmt.Sprintf("mismatch (-want +got):\n%s", diff))
	}
}
