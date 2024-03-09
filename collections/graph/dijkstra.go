// dir graph, undir graph, weighted, unweighted
package graph

import (
	cl "github.com/lorenzotinfena/goji/collections"
	"github.com/lorenzotinfena/goji/collections/heap"
	"github.com/lorenzotinfena/goji/utils"
	constr "github.com/lorenzotinfena/goji/utils/constraints"
)

type DijkstraNode[V comparable, W constr.Integer | constr.Float] struct {
	Vertex   V
	Previous *DijkstraNode[V, W]
	Cost     W
}

type unitDijkstraIterator[V comparable] struct {
	toAnalyze cl.Queue[DijkstraNode[V, int]]
	adjacents func(V) []V
	addV      func(V)
	containsV func(V) bool
}

func (it *unitDijkstraIterator[V]) HasNext() bool {
	return it.toAnalyze.Len() != 0
}

func (it *unitDijkstraIterator[V]) Next() DijkstraNode[V, int] {
	cur := it.toAnalyze.Dequeue()
	for _, v := range it.adjacents(cur.Vertex) {
		if !it.containsV(v) {
			it.toAnalyze.Enqueue(DijkstraNode[V, int]{Vertex: v, Previous: &cur, Cost: cur.Cost + 1})
			it.addV(v)
		}
	}
	return cur
}

func UnitDijkstra[V comparable](start V, adjacents func(V) []V, addV func(V), containsV func(V) bool) utils.Iterator[DijkstraNode[V, int]] {
	toAnalyze := *cl.NewQueue[DijkstraNode[V, int]]()
	toAnalyze.Enqueue(DijkstraNode[V, int]{Vertex: start, Previous: nil, Cost: 0})
	addV(start)

	return &unitDijkstraIterator[V]{
		adjacents: adjacents,
		toAnalyze: toAnalyze,
		addV:      addV,
		containsV: containsV,
	}
}

type weightedDijkstraIterator[V comparable, W constr.Integer | constr.Float] struct {
	toAnalyze        *heap.FibonacciHeap[V]
	adjacents        func(V) []cl.Pair[V, W]
	dijkstraSet      func(V, *DijkstraNode[V, W])
	dijkstraGet      func(V) *DijkstraNode[V, W]
	dijkstraContains func(V) bool
}

func (it *weightedDijkstraIterator[V, W]) HasNext() bool {
	return it.toAnalyze.Len() != 0
}

func (it *weightedDijkstraIterator[V, W]) Next() DijkstraNode[V, W] {
	curr := it.toAnalyze.Pop()
	for _, next := range it.adjacents(curr) {
		costToNext := it.dijkstraGet(curr).Cost + next.Second
		if !it.dijkstraContains(next.First) {
			it.dijkstraSet(next.First, &DijkstraNode[V, W]{next.First, it.dijkstraGet(curr), it.dijkstraGet(curr).Cost + next.Second})
			it.toAnalyze.Push(next.First)
		} else if costToNext < it.dijkstraGet(next.First).Cost {
			it.dijkstraGet(next.First).Cost = costToNext
			it.dijkstraGet(next.First).Previous = it.dijkstraGet(curr)
			it.toAnalyze.DecreaseKey(next.First, next.First)
		}
	}
	return *it.dijkstraGet(curr)
}

func WeightedDijkstra[V comparable, W constr.Integer | constr.Float](
	start V,
	adjacents func(V) []cl.Pair[V, W],

	dijkstraSet func(V, *DijkstraNode[V, W]),
	dijkstraGet func(V) *DijkstraNode[V, W],
	dijkstraContains func(V) bool,
	fibonacciSet func(V, []*heap.FibonacciHeapNode[V]),
	fibonacciGet func(V) []*heap.FibonacciHeapNode[V],
	fibonacciContains func(V) bool,
	fibonacciRemove func(V),
) utils.Iterator[DijkstraNode[V, W]] {
	toAnalyze := heap.NewFibonacciHeap[V](
		func(v1, v2 V) bool { return dijkstraGet(v1).Cost < dijkstraGet(v2).Cost },
		fibonacciSet, fibonacciGet, dijkstraContains, fibonacciRemove)
	dijkstraSet(start, &DijkstraNode[V, W]{start, nil, W(0)})
	toAnalyze.Push(start)

	return &weightedDijkstraIterator[V, W]{
		toAnalyze,
		adjacents,
		dijkstraSet,
		dijkstraGet,
		dijkstraContains,
	}
}
