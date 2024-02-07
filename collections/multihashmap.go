package collections

import "github.com/lorenzotinfena/goji/utils"

type MultiHashMap[K comparable, V comparable] struct {
	m *HashMap[K, *HashSetWithSample[V]]
}

func NewMultiHashMap[K comparable, V comparable]() *MultiHashMap[K, V] {
	return &MultiHashMap[K, V]{m: NewHashMap[K, *HashSetWithSample[V]]()}
}

func (h *MultiHashMap[K, V]) Add(key K, value V) {
	if h.m.Contains(key) {
		h.m.Get(key).Add(value)
	} else {
		tmp := NewHashSetWithSample[V]()
		tmp.Add(value)
		h.m.Set(key, tmp)
	}
}

func (h *MultiHashMap[K, V]) GetOne(key K) V {
	return h.m.Get(key).Sample()
}

func (h *MultiHashMap[K, V]) Contains(key K) bool {
	return h.m.Contains(key)
}

func (h *MultiHashMap[K, V]) Iterator(key K) utils.Iterator[V] {
	return NewVectorIterator(h.m.Get(key).Items())
}

func (h *MultiHashMap[K, V]) Remove(key K, value V) {
	tmp := h.m.Get(key)
	tmp.Remove(value)
	if tmp.Len() == 0 {
		h.m.Remove(key)
	}
}
