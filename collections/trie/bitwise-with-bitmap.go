package trie

import (
	"math/bits"
	"unsafe"

	"github.com/lorenzotinfena/goji/utils/constraints"
	"github.com/lorenzotinfena/goji/utils/slices"
)

type nodeBitwiseTrieWithBitmap[S constraints.Unsigned] struct {
	end    bool // denote if there is an element till here
	bitmap uint
	next   []*nodeBitwiseTrieWithBitmap[S]
}

// Array mapped trie
// https://en.wikipedia.org/wiki/Bitwise_trie_with_bitmap
type BitwiseTrieWithBitmap[S constraints.Unsigned] struct {
	root                 nodeBitwiseTrieWithBitmap[S]
	length               int
	sizeStructureElement int
	mask                 S
}

func NewBitwiseTrieWithBitmap[S constraints.Unsigned]() *BitwiseTrieWithBitmap[S] {
	var tmp S
	return &BitwiseTrieWithBitmap[S]{
		root:                 nodeBitwiseTrieWithBitmap[S]{next: make([]*nodeBitwiseTrieWithBitmap[S], 0)},
		length:               0,
		sizeStructureElement: int(unsafe.Sizeof(tmp) * 8),
		mask:                 S(63),
	}
}

func (t *BitwiseTrieWithBitmap[S]) Insert(element []S) {
	n := &t.root
	for _, s := range element {
		for b := t.sizeStructureElement; b > 0; b -= 6 {
			data := t.mask & s
			pos := bits.OnesCount(n.bitmap >> (64 - data))
			if n.bitmap&((1<<63)>>data) != 0 {
				n = n.next[pos]
			} else {
				next := &nodeBitwiseTrieWithBitmap[S]{next: make([]*nodeBitwiseTrieWithBitmap[S], 0)}
				n.next = slices.Insert(n.next, pos, next)
				n.bitmap |= ((1 << 63) >> data)
				n.end = true
			}
			s >>= 6
		}
	}
	n.end = true
	t.length++
}

func (t *BitwiseTrieWithBitmap[S]) Contains(element []S) bool {
	n := &t.root
	for _, s := range element {
		for b := t.sizeStructureElement; b > 0; b -= 6 {
			data := t.mask & s
			pos := bits.OnesCount(n.bitmap >> (64 - data))
			if n.bitmap&((1<<63)>>data) != 0 {
				n = n.next[pos]
			} else {
				return false
			}
			s >>= 6
		}
	}
	return n.end
}

func (t *BitwiseTrieWithBitmap[S]) Len() int {
	return t.length
}
