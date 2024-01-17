package dynamic

import (
	"testing"

	"github.com/teleivo/algorithm-design-manual/require"
)

func TestArray(t *testing.T) {
	t.Run("GrowAndShrink", func(t *testing.T) {
		a := &array{}

		{
			t.Log("Grow")
			for i := 0; i < 16; i++ {
				t.Logf("len %d, cap %d\n", a.len(), a.cap())
				t.Logf("insert %d\n", i)

				curLen := a.len()
				curCap := a.cap()

				a.insert(i, i)

				require.Equals(t, "len", a.len(), i+1)
				// double if full
				if curLen == curCap {
					wantCap := max(curCap*2, 1)
					require.Equals(t, "cap", a.cap(), wantCap)
				} else {
					require.Equals(t, "cap", a.cap(), curCap)
				}
			}
		}

		{
			t.Log("InsertBelowLen")

			require.Equals(t, "len", a.len(), 16)
			require.Equals(t, "cap", a.cap(), 16)

			a.insert(15, 90)

			require.Equals(t, "len", a.len(), 16)
			require.Equals(t, "cap", a.cap(), 16)
		}

		{
			t.Log("Shrink")
			for i := a.len(); i >= 1; i-- {
				t.Logf("len %d, cap %d\n", a.len(), a.cap())
				t.Logf("delete %d\n", i)

				curLen := a.len()
				curCap := a.cap()

				a.delete()

				require.Equals(t, "len", a.len(), i-1)
				// shrink by half if quarter full
				if curLen == curCap/4 {
					wantCap := curCap / 2
					require.Equals(t, "cap", a.cap(), wantCap)
				} else {
					require.Equals(t, "cap", a.cap(), curCap)
				}
			}
		}
	})

	t.Run("InsertMoreThanOneIndexPastCapacity", func(t *testing.T) {
		a := &array{}

		a.insert(110, 32)

		require.Equals(t, "get", a.get(110), 32)
		require.Equals(t, "len", a.len(), 111)
		require.Equals(t, "cap", a.cap(), 111)
	})
}
