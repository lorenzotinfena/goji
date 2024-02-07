package heap

import (
	"github.com/lorenzotinfena/goji/collections"
)

// Assumptions:
// If a==b then not a<b and not a>b
type BinaryHeapWithRemove[T any] struct {
	v     []T
	m     collections.MultiMap[T, int]
	prior func(T, T) bool
}

// Note for function Prior: It's a strict order relation
func NewBinaryHeapWithRemoveFromSlice[T any](
	v []T,
	prior func(T, T) bool,
	m collections.MultiMap[T, int],
) *BinaryHeapWithRemove[T] {
	h := &BinaryHeapWithRemove[T]{
		v:     v,
		m:     m,
		prior: prior,
	}
	if h.Len() > 1 {
		for i := (h.Len() - 2) / 2; i >= 0; i-- {
			h.heapifyDown(i)
		}
	}
	return h
}

func NewBinaryHeapWithRemove[T any](prior func(T, T) bool, m collections.MultiMap[T, int]) *BinaryHeapWithRemove[T] {
	return &BinaryHeapWithRemove[T]{v: make([]T, 0), m: m, prior: prior}
}

func (h *BinaryHeapWithRemove[T]) Len() int {
	return len(h.v)
}

func (h *BinaryHeapWithRemove[T]) Push(value T) {
	h.m.Add(value, len(h.v))
	h.v = append(h.v, value)
	h.heapifyUp(h.Len() - 1)
}

func (h *BinaryHeapWithRemove[T]) Pop() (res T) {
	res = h.v[0]
	h.m.Remove(h.v[0], 0)
	h.v[0] = h.v[h.Len()-1]
	h.m.Remove(h.v[0], h.Len()-1)
	h.m.Add(h.v[0], 0)
	h.v = h.v[:h.Len()-1]
	h.heapifyDown(0)
	return
}

func (h *BinaryHeapWithRemove[T]) heapifyDown(index int) bool {
	origin := index
	for {
		j := index*2 + 2
		if j < h.Len() {
			if h.prior(h.v[j-1], h.v[j]) {
				j--
			}
		} else {
			j--
			if j >= h.Len() {
				break
			}
		}
		if h.prior(h.v[j], h.v[index]) {

			h.v[j], h.v[index] = h.v[index], h.v[j]
			h.m.Remove(h.v[j], index)
			h.m.Remove(h.v[index], j)
			h.m.Add(h.v[j], j)
			h.m.Add(h.v[index], index)
			index = j
		} else {
			break
		}
	}
	return origin != index
}

func (h *BinaryHeapWithRemove[T]) heapifyUp(index int) {
	for {
		if index == 0 {
			break
		}
		parent := (index - 1) / 2
		if h.prior(h.v[parent], h.v[index]) {
			break
		}
		h.v[index], h.v[parent] = h.v[parent], h.v[index]
		h.m.Remove(h.v[index], parent)
		h.m.Remove(h.v[parent], index)
		h.m.Add(h.v[index], index)
		h.m.Add(h.v[parent], parent)
		index = parent
	}
}

func (h *BinaryHeapWithRemove[T]) Preview() T {
	return h.v[0]
}

func (h BinaryHeapWithRemove[T]) String() string {
	return "" // #TODO
}

func (h *BinaryHeapWithRemove[T]) Remove(element T) {
	i := h.m.GetOne(element)
	h.m.Remove(element, i)
	if i == h.Len()-1 {
		h.v = h.v[:h.Len()-1]
	} else {
		h.m.Remove(h.v[h.Len()-1], h.Len()-1)
		h.m.Add(h.v[h.Len()-1], i)
		h.v[i] = h.v[h.Len()-1]
		h.v = h.v[:h.Len()-1]
		if !h.heapifyDown(i) {
			h.heapifyUp(i)
		}
	}
}
