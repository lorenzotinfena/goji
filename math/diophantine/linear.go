package diophantine

import (
	"github.com/lorenzotinfena/goji/utils/constraints"
)

// gcd can be negative
// source: https://cp-algorithms.com/algebra/extended-euclid-algorithm.html#algorithm
// Assumptions:
// - At least one != 0
func ExtendedEuclideanAlgorithm[T constraints.Integer](a, b T) (gcd T, x, y int) {
	if b == 0 {
		return a, 1, 0
	}
	gcd, xTmp, yTmp := ExtendedEuclideanAlgorithm(b, a%b)
	y = xTmp - yTmp*int(a/b)
	x = yTmp
	return gcd, x, y
}

// See https://cp-algorithms.com/algebra/linear-diophantine-equation.html#algorithmic-solution
// Given ax+by=c returns:
// - if it has solutions (zero or infinite)
// - gcd(a,b)
// - one solution
func ComputeLinearDiophantine[T constraints.Integer](a, b, c T) (hasSolutions bool, gcd T, x, y int) {
	g, x, y := ExtendedEuclideanAlgorithm(a, b)
	if c%g != 0 {
		return false, 0, 0, 0
	}

	factor := int(c / g)
	x *= factor
	y *= factor
	return true, g, x, y
}
