package diophantine_test

import (
	"testing"

	"github.com/lorenzotinfena/goji/math"
	"github.com/lorenzotinfena/goji/math/diophantine"
	"github.com/stretchr/testify/assert"
)

func TestExtendedEuclideanAlgorithm(t *testing.T) {
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
		gcd, x, y := diophantine.ExtendedEuclideanAlgorithm(ts.a, ts.b)
		assert.Equal(t, math.Abs(gcd), ts.gcd)
		assert.Equal(t, ts.a*x+ts.b*y, gcd)
	}
}

func TestComputeLinearDiophantine(t *testing.T) {
	testcases := []struct {
		a, b, c, gcd int
	}{
		// with solutions
		{0, 7, 49, 7},
		{7, 0, 14, 7},
		{11, 7, 101, 1},
		{126, 435, 0, 3},
		{6, 14, 2, 2},
		{12, 18, 12, 6},
		{41 * 129, 59 * 129, 129 * 5, 129},
		{-12, 18, 0, 6},
		{12, -18, -12, 6},
		{-12, -18, 12, 6},

		// no solutions
		{0, 7, 8, 7},
		{7, 0, 1, 7},
		{41 * 129, 59 * 129, 129*5 + 1, 129},
		{-12, 18, -1, 6},
		{12, -18, -11, 6},
		{-12, -18, 11, 6},
	}
	for _, ts := range testcases {
		exist, gcd, x, y := diophantine.ComputeLinearDiophantine(ts.a, ts.b, ts.c)
		assert.Equal(t, exist, ts.c%ts.gcd == 0, "ciao")
		if exist {
			assert.Equal(t, math.Abs(gcd), ts.gcd, "ciao")
			assert.Equal(t, ts.a*x+ts.b*y, ts.c, "ciao")
		}
	}
}
