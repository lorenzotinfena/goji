package geometry

import (
	"sort"

	"github.com/lorenzotinfena/goji/math"
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

	// Find starting point
	start := v[0]
	for i := 1; i < len(v); i++ {
		if v[i][1] < start[1] || (v[i][1] == start[1] && v[i][0] < start[0]) {
			start = v[i]
		}
	}

	// Remove duplicates of starting point (to avoid a bad sorting)
	tmp := v[:0]
	for i := 0; i < len(v); {
		if v[i] != start {
			tmp = append(tmp, v[i])
		}
	}
	v = tmp

	// sort by orientation
	sort.Slice(v, func(i, j int) bool {
		return orientation(start, v[i], v[j]) > E(0)
	})

	// Keep only farthest point in colinear points
	tmp = v[:0]
	for i := 0; i < len(v); {
		farthest := v[i]
		i++
		for i < len(v) && orientation(start, farthest, v[i]) == 0 {
			if math.Abs(v[i][0]-start[0]) > math.Abs(farthest[0]-start[0]) || math.Abs(v[i][1]-start[1]) > math.Abs(farthest[1]-start[1]) {
				farthest = v[i]
			}
			i++
		}
		tmp = append(tmp, farthest)
	}
	v = tmp

	stack := [][2]E{start}
	for i := 0; i < len(v); i++ {
		for len(stack) != 1 && orientation(stack[len(stack)-2], stack[len(stack)-1], v[i]) <= 0 {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, v[i])
	}
	return stack
}
