package combinatorics

import "github.com/lorenzotinfena/goji/utils/constraints"

func InversePermutation[T constraints.Integer](permutation []T) []T {
	inverse := make([]T, len(permutation))
	for i := 0; i < len(permutation); i++ {
		inverse[permutation[i]-1] = T(i + 1)
	}
	return inverse
}
