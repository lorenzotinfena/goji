package combinatorics_test

import (
	"testing"

	"github.com/lorenzotinfena/goji/math/combinatorics"
	"github.com/stretchr/testify/assert"
)

func TestRobinsonSchenstedMapping(t *testing.T) {
	p, q := combinatorics.RobinsonSchenstedMapping([]int{3, 2, 5, 4, 6, 1})
	assert.Equal(t, p, [][]int{{1, 4, 6}, {2, 5}, {3}})
	assert.Equal(t, q, [][]int{{1, 3, 5}, {2, 4}, {6}})

	p, q = combinatorics.RobinsonSchenstedMapping([]int{3, 2, 1})
	assert.Equal(t, p, [][]int{{1}, {2}, {3}})
	assert.Equal(t, q, [][]int{{1}, {2}, {3}})
}
