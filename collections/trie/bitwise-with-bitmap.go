package trie

import (
	"math/bits"
	"unsafe"

	"github.com/lorenzotinfena/goji/utils/constraints"
	"github.com/lorenzotinfena/goji/utils/slices"
)

// Implementations of amt-like data structures:
// 							amt | hamt
// keys ordered				yes | no
// multiple keys support	yes	| no
// with value associated	no	| yes
// For other variants, you can write them by your own!

type nodeBitwiseTrieWithBitmap[S constraints.Unsigned] struct {
	// segment tree that counts the elements for each of the 5 bits from the most significant
	// But it sufficient to counts only the 0 bits
	//
	// The first value of the segtree rapresent the count of elements with first bit set to 0, which is 3
	// If I wanna search for the i-th element I will go left if i<count
	// Take the following vertically elements rapresented as bits:
	// 0 0 0 1
	// 0 0 1 0
	// 0 0 1 0
	// 1 1 1 1
	// 1 1 1 1
	// So If Im looking for the 2-th element I will go left, and then right until the end
	//sizeSegmentTree [63]uint
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
