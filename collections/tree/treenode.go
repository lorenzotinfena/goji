package tree

import (
	"github.com/lorenzotinfena/goji/collections"
	"github.com/lorenzotinfena/goji/collections/graph"
)

type treeNode[V comparable] struct {
	Value    V
	Children []*treeNode[V]
}

func (t treeNode[V]) String() string {
	vertices := []collections.Pair[V, int]{}
	adjacents := [][]collections.Pair[V, int]{}
	i := 0
	var build func(t *treeNode[V])
	build = func(t *treeNode[V]) {
		vertices = append(vertices, collections.MakePair(t.Value, i))
		adjacents = append(adjacents, []collections.Pair[V, int]{})
		for j, c := range t.Children {
			adjacents[i] = append(adjacents[i], collections.MakePair(c.Value, i+j+1))
			build(c)
		}
		i++
	}
	return graph.UnitString(vertices, func(p collections.Pair[V, int]) []collections.Pair[V, int] { return adjacents[p.Second] }, nil)
}
