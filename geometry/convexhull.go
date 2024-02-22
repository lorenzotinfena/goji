package geometry

import (
	"sort"

	"github.com/lorenzotinfena/goji/utils/constraints"
)

// Return -1 if p1p2p3 makes a left turn, 1 if right turn, and 0 if they are colinear
func orientation[E constraints.Signed | constraints.Float](p1, p2, p3 [2]E) E {
	crossProductNorm := (p3[0]-p1[0])*(p2[1]-p1[1]) - (p2[0]-p1[0])*(p3[1]-p1[1])
	if crossProductNorm > 0 {
		return 1
	} else if crossProductNorm < 0 {
		return -1
	} else {
		return 0
	}
}

// Graham scan algorithm on 2-dim euclidean space E
func GrahamScan[E constraints.Signed | constraints.Float](v [][2]E) [][2]E {
	if len(v) == 0 {
		return [][2]E{}
	}
	for i := 1; i < len(v); i++ {
		if v[i][1] < v[0][1] || (v[i][1] == v[0][1] && v[i][0] < v[0][0]) {
			v[0], v[i] = v[i], v[0]
		}
	}
	start := v[0]
	v = v[1:]
	sort.Slice(v, func(i, j int) bool {
		return orientation(start, v[i], v[j]) > E(0)
	})

	stack := [][2]E{start}
	for i := 0; i < len(v); i++ {
		if v[i] == start {
			continue
		}
		for len(stack) != 1 && orientation(stack[len(stack)-2], stack[len(stack)-1], v[i]) <= 0 {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, v[i])
	}
	return stack
}
