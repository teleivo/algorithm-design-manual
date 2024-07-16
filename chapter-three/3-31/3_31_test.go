package problem_331

import (
	"testing"

	"github.com/teleivo/assertive/assert"
)

func TestNumbers(t *testing.T) {
	n := New(6, 3)

	assert.False(t, n.Search(4))

	n.Insert(3)
	assert.True(t, n.Search(3))
	assert.False(t, n.Search(4))

	n.Insert(2)
	n.Insert(2)
	assert.True(t, n.Search(2))
	assert.True(t, n.Search(3))
	assert.False(t, n.Search(4))

	n.Delete(3)
	assert.True(t, n.Search(2))
	assert.False(t, n.Search(3))
	assert.False(t, n.Search(4))

	n.Delete(4)
	assert.True(t, n.Search(2))
	assert.False(t, n.Search(3))
	assert.False(t, n.Search(4))
}

func TestNumbersArrayContainingGarbage(t *testing.T) {
	n := &numbers{
		a: make([]int, 4),
		b: make([]int, 4),
	}

	n.a[1] = 3
	n.b[3] = 1
	assert.False(t, n.Search(1))

	n.Insert(2)
	assert.True(t, n.Search(2))
	assert.False(t, n.Search(1))

	n.Insert(1)
	assert.True(t, n.Search(2))
	assert.True(t, n.Search(1))
}
