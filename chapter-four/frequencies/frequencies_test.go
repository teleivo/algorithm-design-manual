package frequencies

import (
	"testing"

	"github.com/teleivo/algorithm-design-manual/assert"
)

func TestCount(t *testing.T) {
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
	}

	for _, tc := range tests {
		gotNum, gotOk := majorityElementUsinMap(tc.in)

		assert.Equals(t, gotNum, tc.wantNum)
		assert.Equals(t, gotOk, tc.wantOk)
	}
}
