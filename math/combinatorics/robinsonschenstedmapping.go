package combinatorics

import (
	"github.com/lorenzotinfena/goji/sort"
	"github.com/lorenzotinfena/goji/utils"
)

// Using code from: codeforces.com/blog/entry/98167
func RobinsonSchenstedMapping(permutation []int) (p, q [][]int) {
	sqrt := 1
	for sqrt*sqrt < len(permutation) {
		sqrt++
	}
	partialRobinsonSchenstedMapping := func(permutation []int) (p [][]int) {
		youngTableaux := [][]int{}
		for i := 0; i < len(permutation); i++ {
			element := permutation[i]
			for j := 0; j < sqrt; j++ {
				if len(youngTableaux) == j {
					youngTableaux = append(youngTableaux, []int{})
				}
				ind := sort.UpperBound(youngTableaux[j], element, utils.Prioritize[int]())
				if ind == len(youngTableaux[j]) {
					youngTableaux[j] = append(youngTableaux[j], element)
					break
				}
				element, youngTableaux[j][ind] = youngTableaux[j][ind], element
			}
		}
		return youngTableaux
	}

	p = partialRobinsonSchenstedMapping(permutation)
	q = partialRobinsonSchenstedMapping(InversePermutation(permutation))
	for i, j := 0, len(permutation)-1; i < j; i, j = i+1, j-1 {
		permutation[i], permutation[j] = permutation[j], permutation[i]
	}
	pt := partialRobinsonSchenstedMapping(permutation)
	qt := partialRobinsonSchenstedMapping(InversePermutation(permutation))
	for i := len(p); i < len(pt[0]); i++ {
		p = append(p, []int{})
		q = append(q, []int{})
		for j := 0; j < len(pt) && i < len(pt[j]); j++ {
			p[i] = append(p[i], pt[j][i])
			q[i] = append(q[i], qt[j][i])
		}
	}
	return
}
