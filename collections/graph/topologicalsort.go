package graph

import "github.com/lorenzotinfena/goji/utils/slices"

// Assumptions:
// - the graph is acyclic
func TopologicalSort[V any](vertices []V, adjacents func(V) []V, setAdd func(V), setContains func(V) bool) []V {
	res := []V{}

	var dfs func(v V)
	dfs = func(v V) {
		setAdd(v)
		for _, v1 := range adjacents(v) {
			if !setContains(v1) {
				dfs(v1)
			}
		}
		res = append(res, v)
	}
	for _, v := range vertices {
		if !setContains(v) {
			dfs(v)
		}
	}
	slices.Reverse(res)
	return res
}
