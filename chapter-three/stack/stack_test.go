package stack

import (
	"testing"

	"github.com/teleivo/algorithm-design-manual/assert"
)

func TestStack(t *testing.T) {
	t.Run("DuplicateMins", func(t *testing.T) {
		st := &stack{}

		assert.True(t, "isEmpty", st.isEmpty())

		// PUSH
		st.push(2)

		got, err := st.findmin()
		assert.NoError(t, err)
		assert.Equals(t, "findmin", got, 2)

		st.push(4)

		got, err = st.findmin()
		assert.NoError(t, err)
		assert.Equals(t, "findmin", got, 2)

		st.push(2)
		st.push(1)

		got, err = st.findmin()
		assert.NoError(t, err)
		assert.Equals(t, "findmin", got, 1)

		st.push(3)
		st.push(1)

		got, err = st.findmin()
		assert.NoError(t, err)
		assert.Equals(t, "findmin", got, 1)

		// POP
		got, err = st.pop()
		assert.NoError(t, err)
		assert.Equals(t, "pop", got, 1)

		got, err = st.findmin()
		assert.NoError(t, err)
		assert.Equals(t, "findmin", got, 1)

		got, err = st.pop()
		assert.NoError(t, err)
		assert.Equals(t, "pop", got, 3)

		got, err = st.pop()
		assert.NoError(t, err)
		assert.Equals(t, "pop", got, 1)

		got, err = st.findmin()
		assert.NoError(t, err)
		assert.Equals(t, "findmin", got, 2)

		got, err = st.pop()
		assert.NoError(t, err)
		assert.Equals(t, "pop", got, 2)

		got, err = st.pop()
		assert.NoError(t, err)
		assert.Equals(t, "pop", got, 4)

		got, err = st.pop()
		assert.NoError(t, err)
		assert.Equals(t, "pop", got, 2)

		assert.True(t, "isEmpty", st.isEmpty())
	})

	t.Run("Basic", func(t *testing.T) {
		st := &stack{}

		assert.True(t, "isEmpty", st.isEmpty())

		st.push(1)
		assert.False(t, "isEmpty", st.isEmpty())

		got, err := st.findmin()
		assert.NoError(t, err)
		assert.Equals(t, "findmin", got, 1)

		got, err = st.pop()
		assert.NoError(t, err)
		assert.Equals(t, "pop", got, 1)

		assert.True(t, "isEmpty", st.isEmpty())

		st.push(1)
		st.push(4)

		got, err = st.findmin()
		assert.NoError(t, err)
		assert.Equals(t, "findmin", got, 1)

		st.push(3)
		st.push(5)

		got, err = st.findmin()
		assert.NoError(t, err)
		assert.Equals(t, "findmin", got, 1)

		got, err = st.pop()
		assert.NoError(t, err)
		assert.Equals(t, "pop", got, 5)

		got, err = st.findmin()
		assert.NoError(t, err)
		assert.Equals(t, "findmin", got, 1)

		got, err = st.pop()
		assert.NoError(t, err)
		assert.Equals(t, "pop", got, 3)

		got, err = st.pop()
		assert.NoError(t, err)
		assert.Equals(t, "pop", got, 4)

		got, err = st.pop()
		assert.NoError(t, err)
		assert.Equals(t, "pop", got, 1)
	})
}
