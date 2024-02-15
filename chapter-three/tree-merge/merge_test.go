package treemerge

import (
	"testing"

	"github.com/teleivo/algorithm-design-manual/assert"
)

func TestMergeImplementations(t *testing.T) {
	implementations := map[string]func(*node, *node, *list){
		"mergeIter": mergeIter,
	}

	for name, f := range implementations {
		t.Run(name+"/OnlyLeftTree", func(t *testing.T) {
			n1 := New(2, 1, 3)
			got := &list{}

			f(n1, nil, got)

			want := []int{
				1, 2, 3,
			}
			assertList(t, got, want)
		})
		t.Run(name+"/LeftTreeSmallerThanRight", func(t *testing.T) {
			n1 := New(2, 1, 3)
			n2 := New(5, 4, 6)
			got := &list{}

			f(n1, n2, got)

			want := []int{
				1, 2, 3, 4, 5, 6,
			}
			assertList(t, got, want)
		})
		t.Run(name+"/RightTreeSmallerThanLeft", func(t *testing.T) {
			n1 := New(5, 4, 6)
			n2 := New(2, 1, 3)
			got := &list{}

			f(n1, n2, got)

			want := []int{
				1, 2, 3, 4, 5, 6,
			}
			assertList(t, got, want)
		})
		t.Run(name+"/EqualValues", func(t *testing.T) {
			n1 := New(2, 1, 4)
			n2 := New(2, 4, 6)
			got := &list{}

			f(n1, n2, got)

			want := []int{
				1, 2, 4, 6,
			}
			assertList(t, got, want)
		})
		t.Run(name+"/Mixed", func(t *testing.T) {
			n1 := New(11, 7, 4, 8, 24, 1, 6)
			n2 := New(3, 22, 2, 9, 5, 36)
			got := &list{}

			f(n1, n2, got)

			want := []int{
				1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 22, 24, 36,
			}
			assertList(t, got, want)
		})
	}
}
func assertList(t *testing.T, got *list, want []int) {
	var result []int
	for cur := got.head; cur != nil; cur = cur.next {
		result = append(result, cur.Value)
	}

	assert.EqualValues(t, result, want)
}
