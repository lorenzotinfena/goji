package combinatorics

import "github.com/lorenzotinfena/goji/math"

func HookLength(partition []int, primeMod int) int {
	partitionTrasnposed := make([]int, partition[0])
	for i := 0; i < len(partition); i++ {
		partitionTrasnposed[partition[i]-1] = i + 1
	}

	productOfHookLengths := 1
	n := 0
	for i := 0; i < len(partition); i++ {
		n += partition[i]
		for j := 0; j < partition[i]; j++ {
			productOfHookLengths *= partitionTrasnposed[j] - i + partition[i] - j - 1
			productOfHookLengths %= primeMod
		}
	}

	hookLength := math.Factorial(n, primeMod) * math.ModularInverse(productOfHookLengths, primeMod)
	hookLength %= primeMod
	return (hookLength * hookLength) % primeMod
}
