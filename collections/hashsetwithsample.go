package collections

import (
	"fmt"

	"github.com/lorenzotinfena/goji/math"
)

type hashSetWithSampleNode[T comparable] struct {
	item  T
	index int
	next  *hashSetWithSampleNode[T]
}
type HashSetWithSample[T comparable] struct {
	v     []*hashSetWithSampleNode[T]
	items []T
}

func NewHashSetWithSample[T comparable]() *HashSetWithSample[T] {
	return &HashSetWithSample[T]{v: []*hashSetWithSampleNode[T]{nil}}
}

func (h *HashSetWithSample[T]) Add(item T) {
	i := math.Hash(item) % uint(len(h.v))
	for tmp := h.v[i]; tmp != nil; tmp = tmp.next {
		if tmp.item == item {
			return
		}
	}
	h.v[i] = &hashSetWithSampleNode[T]{item, len(h.items), h.v[i]}
	h.items = append(h.items, item)
	if len(h.items)*4 > len(h.v) {
		keyValues := make([]Pair[T, int], len(h.items))
		for i, item := range h.items {
			keyValues[i] = MakePair(item, i)
		}
		h.v = make([]*hashSetWithSampleNode[T], len(h.v)*2)
		for _, p := range keyValues {
			i := math.Hash(p.First) % uint(len(h.v))
			h.v[i] = &hashSetWithSampleNode[T]{p.First, p.Second, h.v[i]}
		}
	}
}

func (h *HashSetWithSample[T]) Remove(item T) {
	i := math.Hash(item) % uint(len(h.v))
	tmp := &h.v[i]
	for (*tmp).item != item {
		tmp = &(*tmp).next
	}
	indexToRemove := (*tmp).index
	i = math.Hash(h.items[len(h.items)-1]) % uint(len(h.v))
	tmp2 := h.v[i]
	for tmp2.item != h.items[len(h.items)-1] {
		tmp2 = tmp2.next
	}
	tmp2.index = indexToRemove
	*tmp = (*tmp).next
	h.items[indexToRemove], h.items[len(h.items)-1] = h.items[len(h.items)-1], h.items[indexToRemove]
	h.items = h.items[:len(h.items)-1]
}

func (h HashSetWithSample[T]) Contains(item T) bool {
	i := math.Hash(item) % uint(len(h.v))
	for tmp := h.v[i]; tmp != nil; tmp = tmp.next {
		if tmp.item == item {
			return true
		}
	}
	return false
}

func (h HashSetWithSample[T]) Items() []T {
	return h.items
}

func (h HashSetWithSample[T]) Sample() T {
	return h.items[0]
}

func (h HashSetWithSample[T]) String() string {
	return fmt.Sprint(h.items)
}

func (h HashSetWithSample[T]) Len() int {
	return len(h.items)
}
