package sort_test

import (
	"testing"

	"github.com/lorenzotinfena/goji/sort"
	"github.com/lorenzotinfena/goji/utils"
	"github.com/stretchr/testify/assert"
)

func TestLowerUpperBound(t *testing.T) {
	testcases := []struct {
		v                  []int
		element            int
		expectedLowerBound int
		expectedUpperBound int
	}{
		{[]int{1, 2, 2, 3}, 2, 1, 3},
		{[]int{1, 2, 2, 2, 3}, 2, 1, 4},
		{[]int{1, 2, 4, 5}, 3, 2, 2},
	}
	for _, ts := range testcases {
		assert.Equal(t, ts.expectedLowerBound, sort.LowerBound(ts.v, ts.element, utils.Prioritize[int]()))
		assert.Equal(t, ts.expectedUpperBound, sort.UpperBound(ts.v, ts.element, utils.Prioritize[int]()))
	}
}
