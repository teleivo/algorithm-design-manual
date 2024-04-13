package hotel

import (
	"testing"

	"github.com/teleivo/algorithm-design-manual/require"
)

func TestHotel(t *testing.T) {
	t.Run("CheckInAllRoomsInOrder", func(t *testing.T) {
		n := 3
		availableRooms := n
		h := New(n)

		require.Equals(t, h.Count(1, n), n)

		for i := 1; i <= n; i++ {
			// t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Logf("checkin: %d", i)
			require.Equals(t, h.Checkin(1, n), i)
			availableRooms--
			t.Logf("count: %d", i)
			require.Equals(t, h.Count(1, n), availableRooms)
			// })
		}
	})
}
