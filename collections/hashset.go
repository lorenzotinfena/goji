package collections

import (
	"fmt"

	"github.com/lorenzotinfena/goji/math"
)

type hashSetNode[T comparable] struct {
	item T
	next *hashSetNode[T]
}
type HashSet[T comparable] struct {
	v      []*hashSetNode[T]
	length int
}

func NewHashSet[T comparable]() *HashSet[T] {
	return &HashSet[T]{v: []*hashSetNode[T]{nil}}
}

func (h *HashSet[T]) Add(item T) {
	i := math.Hash(item) % uint(len(h.v))
	for tmp := h.v[i]; tmp != nil; tmp = tmp.next {
		if tmp.item == item {
			return
		}
	}
	h.v[i] = &hashSetNode[T]{item, h.v[i]}
	h.length++
	if h.length*4 > len(h.v) {
		items := make([]T, h.length)
		j := 0
		for _, node := range h.v {
			for node != nil {
				items[j] = node.item
				j++
				node = node.next
			}
		}
		h.v = make([]*hashSetNode[T], len(h.v)*2)
		for _, item := range items {
			i := math.Hash(item) % uint(len(h.v))
			h.v[i] = &hashSetNode[T]{item, h.v[i]}
		}
	}
}

func (h *HashSet[T]) Remove(item T) {
	i := math.Hash(item) % uint(len(h.v))
	tmp := &h.v[i]
	for (*tmp).item != item {
		tmp = &(*tmp).next
	}
	*tmp = (*tmp).next
	h.length--
}

func (h HashSet[T]) Contains(item T) bool {
	i := math.Hash(item) % uint(len(h.v))
	for tmp := h.v[i]; tmp != nil; tmp = tmp.next {
		if tmp.item == item {
			return true
		}
	}
	return false
}

func (h HashSet[T]) Items() []T {
	items := make([]T, h.length)
	j := 0
	for _, node := range h.v {
		for node != nil {
			items[j] = node.item
			j++
			node = node.next
		}
	}
	return items
}

func (h HashSet[T]) String() string {
	return fmt.Sprint(h.Items())
}

func (h HashSet[T]) Len() int {
	return h.length
}
