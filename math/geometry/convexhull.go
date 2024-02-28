package geometry

import (
	"math/bits"
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
// Returns clockwise ordered convex hull
func GrahamScan[E constraints.Signed | constraints.Float](points [][2]E) [][2]E {
	if len(points) == 0 {
		return [][2]E{}
	}

	// Find starting point
	start := points[0]
	for i := 1; i < len(points); i++ {
		if points[i][1] < start[1] || (points[i][1] == start[1] && points[i][0] < start[0]) {
			start = points[i]
		}
	}

	// Remove duplicates of starting point (to avoid a bad sorting)
	v := [][2]E{start}
	for i := 0; i < len(points); i++ {
		if points[i] != start {
			v = append(v, points[i])
		}
	}

	// sort by orientation
	tmp := v[1:]
	sort.Slice(tmp, func(i, j int) bool {
		return orientation(start, tmp[i], tmp[j]) > E(0)
	})

	// Keep only farthest point in colinear points
	tmp = v[:1]
	for i := 1; i < len(v); {
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

	stack := v[:1]
	for i := 1; i < len(v); i++ {
		for len(stack) != 1 && orientation(stack[len(stack)-2], stack[len(stack)-1], v[i]) <= 0 {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, v[i])
	}
	return stack
}

// Returns clockwise ordered convex hull
func ChansAlgorithm[E constraints.Signed | constraints.Float](v [][2]E) [][2]E {
	maxT := (bits.LeadingZeros(0) - bits.LeadingZeros(uint(bits.LeadingZeros(0)-bits.LeadingZeros(uint(len(v)))-1)) - 1) + 1
	for t := 1; t <= maxT; t++ {
		var m int
		if t == maxT {
			m = len(v)
		} else {
			m = 1 << (1 << t)
		}
		convexhulls := make([][][2]E, (len(v)-1)/m+1)
		for i := 0; i < len(convexhulls); i++ {
			convexhulls[i] = GrahamScan(v[i*m : math.Min(len(v), (i+1)*m)])
		}
		kPointers := make([]int, len(convexhulls))
		for k := 0; k < len(convexhulls); k++ {
			c := convexhulls[k]
			ministart := 0
			for i := 1; i < len(c); i++ {
				if c[i][1] < c[ministart][1] || (c[i][1] == c[ministart][1] && c[i][0] < c[ministart][0]) {
					ministart = i
				}
			}
			kPointers[k] = ministart
		}
		vertexSoFar := convexhulls[0][kPointers[0]]
		for k := 1; k < len(kPointers); k++ {
			if convexhulls[k][kPointers[k]][1] < vertexSoFar[1] || (convexhulls[k][kPointers[k]][1] == vertexSoFar[1] && convexhulls[k][kPointers[k]][0] == vertexSoFar[0]) {
				vertexSoFar = v[kPointers[k]]
			}
		}
		result := [][2]E{vertexSoFar}
		for ; m >= 0; m-- {
			// update kPointers without binary search (ti.inf.ethz.ch/ew/lehre/CG09/materials/v2.pdf)
			for k := 0; k < len(kPointers); k++ {
				if len(convexhulls[k]) == 1 {
					continue
				}
				for {
					current := convexhulls[k][kPointers[k]]
					next := convexhulls[k][(kPointers[k]+1)%len(convexhulls[k])]
					tmp := orientation(vertexSoFar, current, next)
					if tmp > 0 && (tmp != 0 || (math.Abs(vertexSoFar[0]-next[0]) <= math.Abs(vertexSoFar[0]-current[0]) && math.Abs(vertexSoFar[1]-next[1]) <= math.Abs(vertexSoFar[1]-current[1]))) {
						break
					}

					kPointers[k] = (kPointers[k] + 1) % len(convexhulls[k])
				}
			}

			// Get next vertex of total convex hull (in case of colinear nexts, take the farthest)
			next := convexhulls[0][kPointers[0]]
			for k := 1; k < len(kPointers); k++ {
				candidate := convexhulls[k][kPointers[k]]
				tmp := orientation(vertexSoFar, candidate, next)
				if tmp > 0 || (tmp == 0 && (math.Abs(vertexSoFar[0]-candidate[0]) > math.Abs(vertexSoFar[0]-next[0]) || math.Abs(vertexSoFar[1]-candidate[1]) > math.Abs(vertexSoFar[1]-next[1]))) {
					next = candidate
				}
			}
			if next == result[0] {
				return result
			}
			result = append(result, next)
			vertexSoFar = next
		}
	}
	panic("error")
}
