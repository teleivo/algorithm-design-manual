package report

import (
	"fmt"
)

// report notifies a user of a failed assertion. Functions like t.Errorf, t.Fatalf.
type report func(string, ...any)

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
