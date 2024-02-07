package tree_test

import (
	"testing"

	"github.com/lorenzotinfena/goji/collections/tree"
	"github.com/lorenzotinfena/goji/math"
	"github.com/stretchr/testify/assert"
)

func TestLazySegmentTree(t *testing.T) {
	s := tree.NewLazySegmentTree[int, int](
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
		func(element int) int { return element },
		func(q1, q2 int) int { return math.Max(q1, q2) },
		func(q1, q2 int) int { return math.Max(q1, q2) },
		func(f1, f2 func(int) int) func(int) int {
			tmp := f1(0) + f2(0)
			return func(i int) int { return i + tmp }
		},
	)
	s.UpdateRange(0, 4, func(q int) int { return q + 1 }, func(l, r, old int) int { return math.Max(l, r) })
	s.UpdateRange(1, 2, func(q int) int { return q + 1 }, func(l, r, old int) int { return math.Max(l, r) })
	s.UpdateRange(5, 5, func(q int) int { return q + 1 }, func(l, r, old int) int { return math.Max(l, r) })
	s.UpdateRange(3, 7, func(q int) int { return q + 1 }, func(l, r, old int) int { return math.Max(l, r) })
	s.UpdateRange(6, 8, func(q int) int { return q + 1 }, func(l, r, old int) int { return math.Max(l, r) })
	elements := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	for i := 0; i <= 4; i++ {
		elements[i]++
	}
	for i := 1; i <= 2; i++ {
		elements[i]++
	}
	for i := 5; i <= 5; i++ {
		elements[i]++
	}
	for i := 3; i <= 7; i++ {
		elements[i]++
	}
	for i := 6; i <= 8; i++ {
		elements[i]++
	}
	// after updates: 2 4 5 6 6 8 9 10 9 10 11
	assert.Equal(t, s.Query(1, 4), math.Max(elements[1:5]...))
	assert.Equal(t, s.Query(3, 7), math.Max(elements[3:8]...))
	assert.Equal(t, s.Query(3, 9), math.Max(elements[3:10]...))
	assert.Equal(t, s.Query(1, 4), math.Max(elements[1:5]...))
	assert.Equal(t, s.Query(3, 7), math.Max(elements[3:8]...))
	assert.Equal(t, s.Query(3, 9), math.Max(elements[3:10]...))
	assert.Equal(t, s.Query(0, 2), math.Max(elements[0:3]...))
}
