package heap_test

import (
	"testing"

	"github.com/lorenzotinfena/goji/collections"
	"github.com/lorenzotinfena/goji/collections/heap"
	"github.com/lorenzotinfena/goji/utils"
	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	b := heap.NewBinaryHeapWithRemove[int](utils.Prioritize[int](), collections.NewMultiHashMap[int, int]())
	m := collections.NewHashMap[int, []*heap.FibonacciHeapNode[int]]()
	f := heap.NewFibonacciHeap[int](
		utils.Prioritize[int](),
		m.Set, func(i int) []*heap.FibonacciHeapNode[int] { return m.Get(i) }, func(i int) bool { return m.Contains(i) }, m.Remove)

	f.Push(0)
	f.Push(1)
	f.Remove(0)
	f.Pop()

	b.Push(5)
	b.Push(6)
	b.Push(7)
	b.Push(4)
	b.Push(6)
	b.Remove(5)

	f.Push(5)
	f.Push(6)
	f.Push(7)
	f.Push(4)
	f.Push(6)
	f.Remove(5)

	b.Push(51)
	b.Push(61)
	b.Push(71)
	b.Push(41)
	b.Push(61)

	f.Push(51)
	f.Push(61)
	f.Push(71)
	f.Push(41)
	f.Push(61)

	assert.Equal(t, b.Preview(), f.Preview())
	assert.Equal(t, b.Len(), f.Len())
	assert.Equal(t, b.Pop(), f.Pop())
	assert.Equal(t, b.Len(), f.Len())
	f.Remove(61)
	b.Remove(61)
	assert.Equal(t, b.Pop(), f.Pop())
	assert.Equal(t, b.Len(), f.Len())
	b.Push(5)
	f.Push(5)
	assert.Equal(t, b.Pop(), f.Pop())
	assert.Equal(t, b.Pop(), f.Pop())
	b.Push(2)
	b.Push(2)
	b.Push(1)
	b.Push(3)
	f.Push(2)
	f.Push(2)
	f.Push(1)
	f.Push(3)
	assert.Equal(t, b.Pop(), f.Pop())
	assert.Equal(t, b.Len(), f.Len())
}

func TestFibonacci2(t *testing.T) {
	m := collections.NewHashMap[int, []*heap.FibonacciHeapNode[int]]()
	f := heap.NewFibonacciHeap[int](utils.Prioritize[int](),
		m.Set, func(i int) []*heap.FibonacciHeapNode[int] { return m.Get(i) }, func(i int) bool { return m.Contains(i) }, m.Remove)

	f.Push(5)
	f.Push(6)
	f.Push(7)
	f.Push(4)
	f.Push(6)
	f.Remove(5)
	f.Push(51)
	f.Push(61)
	f.Push(71)
	f.Push(41)
	f.Push(61)
	assert.Equal(t, 4, f.Pop())
	f.DecreaseKey(61, 1)
	assert.Equal(t, 1, f.Pop())
	f.DecreaseKey(41, 7)
	assert.Equal(t, 6, f.Pop())
	f.Remove(61)
	f.Remove(7)
	f.Remove(51)
	assert.Equal(t, 6, f.Pop())
	assert.Equal(t, 7, f.Pop())
	assert.Equal(t, 71, f.Pop())
}
