package dp_test

import (
	"testing"

	. "github.com/lorenzotinfena/goji/collections"
	. "github.com/lorenzotinfena/goji/dp"
	. "github.com/lorenzotinfena/goji/math"
	"github.com/stretchr/testify/assert"
)

func TestDP(t *testing.T) {
	// knapsack 0-1
	values := []int{20, 5, 10, 40, 15, 25}
	weights := []int{1, 2, 3, 8, 7, 4}
	W := 10
	dp := NewDP(
		func(get func(Pair[int, int]) int,
			k Pair[int, int],
		) int {
			if k.First < 0 {
				return 0
			}
			max := get(MakePair(k.First-1, k.Second))
			if k.Second >= weights[k.First] {
				max = Max(max, values[k.First]+get(MakePair(k.First-1, k.Second-weights[k.First])))
			}
			return max
		},
	)
	assert.Equal(t, 60, dp.Get(MakePair(len(values)-1, W)))
}
