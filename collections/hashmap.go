package collections

import (
	"fmt"

	"github.com/lorenzotinfena/goji/math"
)

type hashMapNode[K comparable, V any] struct {
	key   K
	value V
	next  *hashMapNode[K, V]
}
type HashMap[K comparable, V any] struct {
	v      []*hashMapNode[K, V]
	length int
}

func NewHashMap[K comparable, V any]() *HashMap[K, V] {
	return &HashMap[K, V]{v: []*hashMapNode[K, V]{nil}}
}

func (h *HashMap[K, V]) Set(key K, value V) {
	i := math.Hash(key) % uint(len(h.v))
	for tmp := h.v[i]; tmp != nil; tmp = tmp.next {
		if tmp.key == key {
			tmp.value = value
			return
		}
	}
	h.v[i] = &hashMapNode[K, V]{key, value, h.v[i]}
	h.length++
	if h.length*4 > len(h.v) {
		keyValues := make([]Pair[K, V], h.length)
		j := 0
		for _, node := range h.v {
			for node != nil {
				keyValues[j] = MakePair(node.key, node.value)
				j++
				node = node.next
			}
		}
		h.v = make([]*hashMapNode[K, V], len(h.v)*2)
		for _, p := range keyValues {
			i := math.Hash(p.First) % uint(len(h.v))
			h.v[i] = &hashMapNode[K, V]{p.First, p.Second, h.v[i]}
		}
	}
}

func (h *HashMap[K, V]) Remove(key K) {
	i := math.Hash(key) % uint(len(h.v))
	tmp := &h.v[i]
	for (*tmp).key != key {
		tmp = &(*tmp).next
	}
	*tmp = (*tmp).next
	h.length--
}

func (h HashMap[K, V]) Get(key K) V {
	i := math.Hash(key) % uint(len(h.v))
	tmp := h.v[i]
	for {
		if tmp.key == key {
			return tmp.value
		}
		tmp = tmp.next
	}
}

func (h HashMap[K, V]) Contains(key K) bool {
	i := math.Hash(key) % uint(len(h.v))
	for tmp := h.v[i]; tmp != nil; tmp = tmp.next {
		if tmp.key == key {
			return true
		}
	}
	return false
}

func (h HashMap[K, V]) Keys() []K {
	keys := make([]K, h.length)
	j := 0
	for _, node := range h.v {
		for node != nil {
			keys[j] = node.key
			j++
			node = node.next
		}
	}
	return keys
}

func (h HashMap[K, V]) String() string {
	res := make([]string, 0)
	for _, k := range h.Keys() {
		res = append(res, MakePair(k, h.Get(k)).String())
	}
	return fmt.Sprint(res)
}

func (h HashMap[K, V]) Len() int {
	return h.length
}
