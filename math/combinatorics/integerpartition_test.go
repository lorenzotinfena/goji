package combinatorics_test

import (
	"testing"

	"github.com/lorenzotinfena/goji/math/combinatorics"
	"github.com/stretchr/testify/assert"
)

func TestIntegerPartition(t *testing.T) {
	assert.Equal(t, combinatorics.IntegerPartitions(3), [][]int{{1, 1, 1}, {2, 1}, {3}})
}
