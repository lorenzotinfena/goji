package math

import "github.com/lorenzotinfena/goji/utils/constraints"

// Euclidean algorithm
// It returns the positive solution
// Assumptions:
// - At least one != 0
func GCD[T constraints.Integer](a, b T) T {
	if b == 0 {
		return Abs(a)
	}
	return GCD(b, a%b)
}

// It returns the positive solution
// If one is 0, returns 0
// Assumptions:
// - Both != 0
func LCM[T constraints.Integer](a, b T) T {
	return Abs(a*b) / GCD(a, b)
}

// Using code from: cp-algorithms.com/algebra/phi-function.html
// Assumptions:
// - n >= 1
func EulerTotient(n int) int {
	result := n

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			for n%i == 0 {
				n /= i
			}
			result -= result / i
		}
	}
	if n > 1 {
		result -= result / n
	}
	return result
}
