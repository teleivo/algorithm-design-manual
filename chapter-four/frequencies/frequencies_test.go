package frequencies

import (
	"testing"

	"github.com/teleivo/assertive/assert"
)

func TestMajority(t *testing.T) {
	tests := []struct {
		in      []int
		wantNum int
		wantOk  bool
	}{
		{
			in:      []int{1, 10, 3, 1, 1},
			wantNum: 1,
			wantOk:  true,
		},
		{
			in:      []int{1, 10, 1, 1, 3, 1},
			wantNum: 1,
			wantOk:  true,
		},
		{
			in:      []int{1, 10, 1, 1, 3, 7},
			wantNum: 0,
			wantOk:  false,
		},
		{
			in:      []int{1, 10, 3, 1, 2},
			wantNum: 0,
			wantOk:  false,
		},
		{
			in:      []int{1, 10, 3, 1, 2, 4},
			wantNum: 0,
			wantOk:  false,
		},
		{
			in:      []int{1, 10, 3, 7, 2, 4},
			wantNum: 0,
			wantOk:  false,
		},
	}

	t.Run("UsingMap", func(t *testing.T) {
		for _, tc := range tests {
			gotNum, gotOk := majorityElementUsingMap(tc.in)

			assert.Equals(t, gotNum, tc.wantNum)
			assert.Equals(t, gotOk, tc.wantOk)
		}
	})

	t.Run("UsingStack", func(t *testing.T) {
		for _, tc := range tests {
			gotNum, gotOk := majorityElementUsingStack(tc.in)

			assert.Equals(t, gotNum, tc.wantNum)
			assert.Equals(t, gotOk, tc.wantOk)
		}
	})
}
