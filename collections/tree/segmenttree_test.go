package tree_test

import (
	"testing"

	"github.com/lorenzotinfena/goji/collections/tree"
	"github.com/stretchr/testify/assert"
)

func TestSegmentTree(t *testing.T) {
	s := tree.NewSegmentTree[int, int](
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
		func(element int) int { return element },
		func(q1, q2 int) int { return q1 + q2 },
		func(q1, q2 int) int { return q1 + q2 },
	)

	assert.Equal(t, s.Query(1, 4), 2+3+4+5)
	assert.Equal(t, s.Query(3, 7), 4+5+6+7+8)
	assert.Equal(t, s.Query(3, 9), 4+5+6+7+8+9+10)
	s.Update(4, func(oldQ int) (newQ int) { return oldQ - 5 + 0 })
	assert.Equal(t, s.Query(1, 4), 2+3+4+5-5)
	assert.Equal(t, s.Query(3, 7), 4+5+6+7+8-5)
	assert.Equal(t, s.Query(3, 9), 4+5+6+7+8+9+10-5)

	s = tree.NewSegmentTree[int, int](
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
		func(element int) int { return element },
		func(q1, q2 int) int { return q1 + q2 },
		func(q1, q2 int) int { return q1 + q2 },
	)

	assert.Equal(t, s.RightMost(3, func(q1, q2 int) int { return q1 + q2 }, func(q int) bool { return q <= 4+5+6+7+8+9+10 }), 10)
	assert.Equal(t, s.RightMost(3, func(q1, q2 int) int { return q1 + q2 }, func(q int) bool { return q <= 4+5+6+7 }), 7)
	assert.Equal(t, s.RightMost(4, func(q1, q2 int) int { return q1 + q2 }, func(q int) bool { return q <= 5 }), 5)
	assert.Equal(t, s.RightMost(4, func(q1, q2 int) int { return q1 + q2 }, func(q int) bool { return q <= 4 }), 4)

	s = tree.NewSegmentTree[int, int](
		[]int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		func(element int) int { return element },
		func(q1, q2 int) int { return q1 + q2 },
		func(q1, q2 int) int { return q1 + q2 },
	)

	assert.Equal(t, s.LeftMost(7, func(q1, q2 int) int { return q1 + q2 }, func(q int) bool { return q <= 4+5+6+7+8+9+10 }), 0)
	assert.Equal(t, s.LeftMost(7, func(q1, q2 int) int { return q1 + q2 }, func(q int) bool { return q <= 4+5+6+7 }), 3)
	assert.Equal(t, s.LeftMost(6, func(q1, q2 int) int { return q1 + q2 }, func(q int) bool { return q <= 5 }), 5)
	assert.Equal(t, s.LeftMost(6, func(q1, q2 int) int { return q1 + q2 }, func(q int) bool { return q <= 4 }), 6)
}
