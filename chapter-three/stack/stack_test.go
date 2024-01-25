package stack

import (
	"testing"

	"github.com/teleivo/algorithm-design-manual/require"
)

func TestStack(t *testing.T) {
	t.Run("DuplicateMins", func(t *testing.T) {
		st := &stack{}

		require.True(t, st.isEmpty())

		// PUSH
		st.push(2)

		got, err := st.findmin()
		require.NoError(t, err)
		require.Equals(t, "findmin", got, 2)

		st.push(4)

		got, err = st.findmin()
		require.NoError(t, err)
		require.Equals(t, "findmin", got, 2)

		st.push(2)
		st.push(1)

		got, err = st.findmin()
		require.NoError(t, err)
		require.Equals(t, "findmin", got, 1)

		st.push(3)
		st.push(1)

		got, err = st.findmin()
		require.NoError(t, err)
		require.Equals(t, "findmin", got, 1)

		// POP
		got, err = st.pop()
		require.NoError(t, err)
		require.Equals(t, "pop", got, 1)

		got, err = st.findmin()
		require.NoError(t, err)
		require.Equals(t, "findmin", got, 1)

		got, err = st.pop()
		require.NoError(t, err)
		require.Equals(t, "pop", got, 3)

		got, err = st.pop()
		require.NoError(t, err)
		require.Equals(t, "pop", got, 1)

		got, err = st.findmin()
		require.NoError(t, err)
		require.Equals(t, "findmin", got, 2)

		got, err = st.pop()
		require.NoError(t, err)
		require.Equals(t, "pop", got, 2)

		got, err = st.pop()
		require.NoError(t, err)
		require.Equals(t, "pop", got, 4)

		got, err = st.pop()
		require.NoError(t, err)
		require.Equals(t, "pop", got, 2)

		require.True(t, st.isEmpty())
	})

	t.Run("Basic", func(t *testing.T) {
		st := &stack{}

		require.True(t, st.isEmpty())

		st.push(1)
		require.False(t, st.isEmpty())

		got, err := st.findmin()
		require.NoError(t, err)
		require.Equals(t, "findmin", got, 1)

		got, err = st.pop()
		require.NoError(t, err)
		require.Equals(t, "pop", got, 1)

		require.True(t, st.isEmpty())

		st.push(1)
		st.push(4)

		got, err = st.findmin()
		require.NoError(t, err)
		require.Equals(t, "findmin", got, 1)

		st.push(3)
		st.push(5)

		got, err = st.findmin()
		require.NoError(t, err)
		require.Equals(t, "findmin", got, 1)

		got, err = st.pop()
		require.NoError(t, err)
		require.Equals(t, "pop", got, 5)

		got, err = st.findmin()
		require.NoError(t, err)
		require.Equals(t, "findmin", got, 1)

		got, err = st.pop()
		require.NoError(t, err)
		require.Equals(t, "pop", got, 3)

		got, err = st.pop()
		require.NoError(t, err)
		require.Equals(t, "pop", got, 4)

		got, err = st.pop()
		require.NoError(t, err)
		require.Equals(t, "pop", got, 1)
	})
}
