package trie

import (
	"math/bits"
	"unsafe"

	coll "github.com/lorenzotinfena/goji/collections"
	"github.com/lorenzotinfena/goji/utils/constraints"
	"github.com/lorenzotinfena/goji/utils/slices"
)

// Implementations of amt-like data structures:
// 							amt | hamt
// keys ordered				yes | no
// multiple keys support	yes	| no
// with value associated	no	| yes
// For other variants, you can write them by your own!

type nodeBitwiseTrieWithBitmap[T comparable] struct {
	elements coll.MultiSet[T]

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
	bitmap uint
	next   []*nodeBitwiseTrieWithBitmap[T]
}

// Array mapped trie
// https://en.wikipedia.org/wiki/Bitwise_trie_with_bitmap
// Insert with the same structures will be overwritten
type BitwiseTrieWithBitmap[T comparable, S constraints.Unsigned] struct {
	root         nodeBitwiseTrieWithBitmap[T]
	getStructure func(T) []S
	length       int
}

func NewBitwiseTrieWithBitmap[T comparable, S constraints.Unsigned](getStructure func(T) []S) *BitwiseTrieWithBitmap[T, S] {
	return &BitwiseTrieWithBitmap[T, S]{
		root:         nodeBitwiseTrieWithBitmap[T]{next: make([]*nodeBitwiseTrieWithBitmap[T], 0)},
		getStructure: getStructure,
		length:       0,
	}
}

func (t *BitwiseTrieWithBitmap[T, S]) Insert(element T) {
	// get size of structure
	var tmp S
	B := (unsafe.Sizeof(tmp) * 8) - 6

	n := &t.root
	for _, s := range t.getStructure(element) {
		mask := uint(63 << B)
		for b := B; b > 0; b -= 6 {
			data := (mask & uint(s)) >> b
			pos := bits.OnesCount(n.bitmap << (63 - data))
			if n.bitmap&(1<<data) != 0 {
				n = n.next[pos]
			} else {
				next := &nodeBitwiseTrieWithBitmap[T]{next: make([]*nodeBitwiseTrieWithBitmap[T], 0)}
				n.next = slices.Insert(n.next, pos, next)
				n.bitmap |= (1 << data)
			}
			mask >>= 6
		}
	}
	n.elements.Add(element)
	t.length++
}

func (t *BitwiseTrieWithBitmap[T, S]) Len() int {
	return t.length
}
