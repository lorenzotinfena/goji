package trie

import (
	"math/bits"

	coll "github.com/lorenzotinfena/goji/collections"
	ll "github.com/lorenzotinfena/goji/collections/linkedlist"
	"github.com/lorenzotinfena/goji/math"
	"github.com/lorenzotinfena/goji/utils/slices"
)

// Implementations of amt-like data structures:
// 							amt | hamt
// keys ordered				yes | no
// multiple keys support	yes	| no
// with value associated	no	| yes
// For other variants, you can write them by your own!

type node[K comparable, V any] struct {
	bitmap    uint
	hash      uint
	keyValues ll.SinglyLinkedList[coll.Pair[K, V]]
	next      []*node[K, V]
}

// Hash array mapped trie
// https://en.wikipedia.org/wiki/Hash_array_mapped_trie
type HAMT[K comparable, V any] struct {
	root   *node[K, V]
	length int
}

func NewHAMT[K comparable, V any]() *HAMT[K, V] {
	return &HAMT[K, V]{
		root:   nil,
		length: 0,
	}
}

var bitsPrefixSum = [...]uint64{0, 6, 12, 18, 24, 30, 36, 42, 48, 54, 60}

func (t *HAMT[K, V]) Set(key K, value V) {
	element := math.Hash(key)
	if t.length == 0 {
		keyValues := ll.NewSinglyLinkedList[coll.Pair[K, V]](func(a, b coll.Pair[K, V]) bool { return a.First == b.First })
		keyValues.InsertLast(coll.MakePair(key, value))
		t.root = &node[K, V]{bitmap: 0, hash: element, keyValues: *keyValues, next: make([]*node[K, V], 0)}
		t.length++
		return
	}

	n := t.root

	for i := 0; i < len(bitsPrefixSum); i++ {
		if n.hash == element {
			break
		}
		data := (element << bitsPrefixSum[i]) >> (64 - 6)
		if n.bitmap&(1<<data) != 0 {
			n = n.next[bits.OnesCount64(uint64(n.bitmap<<(64-data)))]
		} else {
			keyValues := ll.NewSinglyLinkedList[coll.Pair[K, V]](func(a, b coll.Pair[K, V]) bool { return a.First == b.First })
			keyValues.InsertLast(coll.MakePair(key, value))
			next := &node[K, V]{bitmap: 0, hash: element, keyValues: *keyValues, next: make([]*node[K, V], 0)}
			pos := bits.OnesCount64(uint64(n.bitmap << (63 - data)))
			n.next = slices.Insert(n.next, pos, next)
			n.bitmap |= (1 << data)
			t.length++
			return
		}
	}
	n.keyValues.Remove(coll.MakePair[K, V](key, value))
	n.keyValues.InsertLast(coll.MakePair[K, V](key, value))
}

func (t *HAMT[K, V]) Get(key K) (V, bool) {
	if t.length == 0 {
		var foo V
		return foo, false
	}

	n := t.root
	element := math.Hash(key)
	for i := 0; i < len(bitsPrefixSum); i++ {
		if n.hash == element {
			break
		}
		data := (element << bitsPrefixSum[i]) >> (64 - 6)
		if n.bitmap&(1<<data) != 0 {
			n = n.next[bits.OnesCount(n.bitmap<<(64-data))]
		} else {
			var foo V
			return foo, false
		}
	}
	var foo V
	pair, present := n.keyValues.GetElementEqualsTo(coll.MakePair(key, foo))
	return pair.Second, present
}

func (t *HAMT[K, V]) Len() int {
	return t.length
}

func (t *HAMT[K, V]) Keys() []K {
	res := make([]K, t.Len())
	i := 0
	var visit func(n *node[K, V])
	visit = func(n *node[K, V]) {
		for _, el := range n.keyValues.ToSlice() {
			res[i] = el.First
			i++
		}
		for _, el := range n.next {
			visit(el)
		}
	}
	if t.root != nil {
		visit(t.root)
	}
	return res
}
