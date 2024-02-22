package math

import "github.com/lorenzotinfena/goji/utils/constraints"

// Assumptions:
// len(elements) > 0
func Max[T constraints.Ordered](elements ...T) T {
	m := elements[0]
	for i := 1; i < len(elements); i++ {
		if elements[i] > m {
			m = elements[i]
		}
	}
	return m
}

// Assumptions:
// len(elements) > 0
func Min[T constraints.Ordered](elements ...T) T {
	m := elements[0]
	for i := 1; i < len(elements); i++ {
		if elements[i] < m {
			m = elements[i]
		}
	}
	return m
}

func Abs[T constraints.Integer | constraints.Float](a T) T {
	if a < 0 {
		return -a
	}
	return a
}

func Diff[T constraints.Integer | constraints.Float](a, b T) T {
	if a >= b {
		return a - b
	} else {
		return b - a
	}
}

// Assumption:
// - Multiplication is associative
func Pow2[B constraints.Float | constraints.Integer, P constraints.Integer](base B) B {
	return base * base
}
