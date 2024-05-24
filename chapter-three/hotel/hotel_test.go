package hotel

import (
	"math"
	"strconv"
	"testing"

	"github.com/teleivo/algorithm-design-manual/require"
)

func TestHotel(t *testing.T) {
	t.Run("CheckInAndOutOfAllRoomsInOrder", func(t *testing.T) {
		n := 17
		h := New(n)

		require.Equals(t, h.Count(1, n), n)

		availableRooms := n
		for i := 1; i <= n; i++ {
			t.Run("Checkin"+strconv.Itoa(i), func(t *testing.T) {
				require.Equals(t, h.Checkin(1, n), i)
				availableRooms--
				require.Equals(t, h.Count(1, n), availableRooms)
			})
		}
		for i := 1; i <= n; i++ {
			t.Run("Checkout"+strconv.Itoa(i), func(t *testing.T) {
				h.Checkout(i)
				availableRooms++
				require.Equals(t, h.Count(1, n), availableRooms)
			})
		}
	})

	t.Run("CheckInAndOutOfAllRoomsInReverseOrder", func(t *testing.T) {
		n := 17
		h := New(n)

		require.Equals(t, h.Count(1, n), n)

		availableRooms := n
		for i := n; i > 0; i-- {
			t.Run("Checkin"+strconv.Itoa(i), func(t *testing.T) {
				require.Equals(t, h.Checkin(i, i), i)
				availableRooms--
				require.Equals(t, h.Count(1, n), availableRooms)
			})
		}
		for i := n; i > 0; i-- {
			t.Run("Checkout"+strconv.Itoa(i), func(t *testing.T) {
				h.Checkout(i)
				availableRooms++
				require.Equals(t, h.Count(1, n), availableRooms)
			})
		}
	})

	t.Run("CheckInAndOutOfSameRoom", func(t *testing.T) {
		n := 5
		h := New(n)

		require.Equals(t, h.Count(1, n), n)

		availableRooms := n
		require.Equals(t, h.Checkin(2, 2), 2)
		availableRooms--

		require.Equals(t, h.Count(1, n), availableRooms)

		h.Checkout(2)
		availableRooms++

		require.Equals(t, h.Count(1, n), availableRooms)

		require.Equals(t, h.Checkin(2, 2), 2)
	})

	t.Run("CheckIntoOccupiedRoom", func(t *testing.T) {
		n := 3
		h := New(n)

		availableRooms := n
		require.Equals(t, h.Checkin(2, 2), 2)
		availableRooms--
		require.Equals(t, h.Count(1, n), availableRooms)

		require.Equals(t, h.Checkin(2, 2), math.MaxInt)
		require.Equals(t, h.Count(1, n), availableRooms)
	})

	t.Run("CheckInAndOutOfDifferentRanges", func(t *testing.T) {
		n := 9
		h := New(n)

		require.Equals(t, h.Count(1, n), n)

		availableRooms := n
		require.Equals(t, h.Checkin(3, 4), 3)
		availableRooms--

		require.Equals(t, h.Count(1, n), availableRooms)

		require.Equals(t, h.Checkin(1, 2), 1)
		availableRooms--

		require.Equals(t, h.Count(1, n), availableRooms)

		require.Equals(t, h.Checkin(1, 2), 2)
		availableRooms--

		require.Equals(t, h.Count(1, n), availableRooms)

		require.Equals(t, h.Checkin(1, 4), 4)
		availableRooms--

		require.Equals(t, h.Count(1, n), availableRooms)

		require.Equals(t, h.Checkin(3, 6), 5)
		availableRooms--

		require.Equals(t, h.Count(1, n), availableRooms)

		require.Equals(t, h.Checkin(7, 8), 7)
		availableRooms--

		require.Equals(t, h.Count(1, n), availableRooms)

		require.Equals(t, h.Checkin(3, 8), 6)
		availableRooms--

		require.Equals(t, h.Count(1, n), availableRooms)
	})

	t.Run("CheckInAndOutOfDifferentRanges2", func(t *testing.T) {
		n := 9
		h := New(n)

		require.Equals(t, h.Count(1, n), n)
		availableRooms := n

		require.Equals(t, h.Checkin(3, 4), 3)
		availableRooms--
		require.Equals(t, h.Count(1, n), availableRooms)

		require.Equals(t, h.Checkin(2, 4), 2)
		availableRooms--
		require.Equals(t, h.Count(1, n), availableRooms)

		require.Equals(t, h.Checkin(2, 4), 4)
		availableRooms--
		require.Equals(t, h.Count(1, n), availableRooms)

		require.Equals(t, h.Checkin(1, 4), 1)
		availableRooms--
		require.Equals(t, h.Count(1, n), availableRooms)

		require.Equals(t, h.Checkin(7, 8), 7)
		availableRooms--
		require.Equals(t, h.Count(1, n), availableRooms)

		require.Equals(t, h.Checkin(1, 8), 5)
		availableRooms--
		require.Equals(t, h.Count(1, n), availableRooms)

		h.Checkout(4)
		availableRooms++

		require.Equals(t, h.Checkin(1, 8), 4)
		availableRooms--
		require.Equals(t, h.Count(1, n), availableRooms)

		require.Equals(t, h.Checkin(7, 8), 8)
		availableRooms--
		require.Equals(t, h.Count(1, n), availableRooms)

		require.Equals(t, h.Checkin(6, 8), 6)
		availableRooms--
		require.Equals(t, h.Count(1, n), availableRooms)
	})
}
