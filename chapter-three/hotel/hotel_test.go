package hotel

import (
	"strconv"
	"testing"

	"github.com/teleivo/algorithm-design-manual/require"
)

func TestHotel(t *testing.T) {
	t.Run("CheckIntoAllRoomsInOrder", func(t *testing.T) {
		n := 17
		h := New(n)

		require.Equals(t, h.Count(1, n), n)

		availableRooms := n
		for i := 1; i <= n; i++ {
			t.Run(strconv.Itoa(i), func(t *testing.T) {
				require.Equals(t, h.Checkin(1, n), i)
				availableRooms--
				t.Logf("Count(1, %d) call #%d should equal %d", n, i, availableRooms)
				require.Equals(t, h.Count(1, n), availableRooms)
			})
		}
	})
}
