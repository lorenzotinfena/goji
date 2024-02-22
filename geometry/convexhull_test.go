package geometry_test

import (
	"sort"
	"testing"

	"github.com/lorenzotinfena/goji/geometry"
	"github.com/stretchr/testify/assert"
)

func TestGrahamScan(t *testing.T) {
	points := [][2]int{
		{-1, -3},
		{-1, -3},
		{-3, -1},
		{-2, -1},
		{-1, -2},
		{-2, 1},
		{-1, 3},
		{-1, 1},
		{-1, 1},
		{0, 0},
		{0, 0},
		{1, 1},
		{1, 1},
		{2, 3},
		{1, -1},
		{0, -2},
		{0, -3},
	}
	convexhull := geometry.GrahamScan(points)
	sort.Slice(convexhull, func(i, j int) bool {
		if convexhull[i][0] == convexhull[j][0] {
			return convexhull[i][1] < convexhull[j][1]
		} else {
			return convexhull[i][0] < convexhull[j][0]
		}
	})
	expectedConvexHull := [][2]int{
		{-3, -1},
		{-1, 3},
		{2, 3},
		{1, -1},
		{0, -3},
		{-1, -3},
	}
	sort.Slice(expectedConvexHull, func(i, j int) bool {
		if expectedConvexHull[i][0] == expectedConvexHull[j][0] {
			return expectedConvexHull[i][1] < expectedConvexHull[j][1]
		} else {
			return expectedConvexHull[i][0] < expectedConvexHull[j][0]
		}
	})
	assert.Equal(t, convexhull, expectedConvexHull)
}
