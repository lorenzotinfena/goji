package math

import "github.com/lorenzotinfena/goji/utils/constraints"

// Pow to an integer using binary exponentiation
// Assumption:
// - Multiplication is associative
func PowGeneric[B any, P constraints.Integer](base B, power P, identity B, mul func(a, b B) B) B {
	res := identity
	for power > 0 {
		if power%2 == 1 {
			res = mul(res, base)
		}
		base = mul(base, base)
		power /= 2
	}
	return res
}

// Pow to an integer using binary exponentiation
// Assumption:
// - Multiplication is associative
func Pow[B constraints.Float | constraints.Integer, P constraints.Integer](base B, power P) B {
	res := B(1)
	for power > 0 {
		if power%2 == 1 {
			res *= base
		}
		base *= base
		power /= 2
	}
	return res
}
