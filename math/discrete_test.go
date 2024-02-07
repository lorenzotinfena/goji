package math_test

import (
	"testing"

	"github.com/lorenzotinfena/goji/math"
	"github.com/stretchr/testify/assert"
)

func TestGCD(t *testing.T) {
	testcases := []struct {
		a, b, gcd int
	}{
		{0, 7, 7},
		{7, 0, 7},
		{11, 7, 1},
		{126, 435, 3},
		{6, 14, 2},
		{12, 18, 6},
		{41 * 129, 59 * 129, 129},
		{-12, 18, 6},
		{12, -18, 6},
		{-12, -18, 6},
	}
	for _, ts := range testcases {
		gcd := math.GCD(ts.a, ts.b)
		assert.Equal(t, math.Abs(gcd), ts.gcd)
	}
}
