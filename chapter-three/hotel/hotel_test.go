package hotel

import (
	"testing"

	"github.com/teleivo/algorithm-design-manual/require"
)

func TestHotel(t *testing.T) {
	n := 8
	availableRooms := n
	h := New(n)

	require.Equals(t, h.Count(1, n), n)

	require.Equals(t, h.Checkin(n), 1)
	availableRooms--
	require.Equals(t, h.Count(1, n), availableRooms)
}
