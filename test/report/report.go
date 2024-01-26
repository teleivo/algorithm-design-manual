package report

import (
	"fmt"
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
