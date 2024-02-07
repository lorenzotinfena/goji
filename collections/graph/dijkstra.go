// dir graph, undir graph, weighted, unweighted
package graph

import (
	cl "github.com/lorenzotinfena/goji/collections"
	"github.com/lorenzotinfena/goji/utils"
	constr "github.com/lorenzotinfena/goji/utils/constraints"
)

type ShortedPathVertex[V comparable, W constr.Integer | constr.Float] struct {
	Vertex   V
	Previous *ShortedPathVertex[V, W]
	Cost     W
}

type unitDijkstraIterator[V comparable] struct {
	adjacents func(V) []V
	toAnalyze cl.Queue[ShortedPathVertex[V, int]]
	visited   cl.HashSet[V]
}

func (it *unitDijkstraIterator[V]) HasNext() bool {
	return it.toAnalyze.Len() != 0
}

func (it *unitDijkstraIterator[V]) Next() ShortedPathVertex[V, int] {
	cur := it.toAnalyze.Dequeue()
	for _, v := range it.adjacents(cur.Vertex) {
		if !it.visited.Contains(v) {
			it.toAnalyze.Enqueue(ShortedPathVertex[V, int]{Vertex: v, Previous: &cur, Cost: cur.Cost + 1})
			it.visited.Add(v)
		}
	}
	return cur
}

func UnitDijkstra[V comparable](start V, adjacents func(V) []V, addV func() V, containsV func(V)) utils.Iterator[ShortedPathVertex[V, int]] {
	toAnalyze := *cl.NewQueue[ShortedPathVertex[V, int]]()
	toAnalyze.Enqueue(ShortedPathVertex[V, int]{Vertex: start, Previous: nil, Cost: 0})
	visited := *cl.NewHashSet[V]()
	visited.Add(start)

	return &unitDijkstraIterator[V]{
		adjacents: adjacents,
		toAnalyze: toAnalyze,
		visited:   visited,
	}
}

/* //TODO
type weightedGraphDijkstraIterator[V comparable] struct {
	bfsIterator utils.Iterator[V]
	last        *ShortedPathVertex[V, int]
}

func (it *unitGraphDijkstraIterator[V]) HasNext() bool {
	return it.bfsIterator.HasNext()
}

func (it *unitGraphDijkstraIterator[V]) Next() ShortedPathVertex[V, int] {
	current := it.bfsIterator.Next()
	if it.last == nil {
		return ShortedPathVertex[V, int]{
			Vertex:   current,
			Previous: it.last,
			Cost:     0,
		}
	} else {
		return ShortedPathVertex[V, int]{
			Vertex:   current,
			Previous: it.last,
			Cost:     it.last.Cost + 1,
		}
	}
}

func (g UnitGraph[V]) Dijkstra(from V) utils.Iterator[ShortedPathVertex[V, int]] {
	return &unitGraphDijkstraIterator[V]{
		bfsIterator: g.GetIteratorBFS(from),
		last:        nil,
	}
}
*/
