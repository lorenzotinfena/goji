package sort

import (
	"unsafe"

	"github.com/lorenzotinfena/goji/collections"
	"github.com/lorenzotinfena/goji/utils/constraints"
)

// Non generic implementation:
/*
func ReverseRadixSort(v []int) {
	var rec func(b, L, R int)
	rec = func(b, L, R int) {
		l := L
		r := R
		for l <= r {
			if v[l]&(1<<b) == 0 {
				l++
			} else {
				v[l], v[r] = v[r], v[l]
				r--
			}
		}
		// The loop finish always in a situation like:
		// ...0000011111...
		// ...L---rl---R...
		if b > 0 {
			if L < r {
				rec(b-1, L, r)
			}
			if l < R {
				rec(b-1, l, R)
			}
		}
	}
	rec(63, 0, len(v)-1)
}
*/
// Assumptions:
// - getStructure always returns slices of same length
func MSDRadixSort[T any, S constraints.Unsigned](v []T, getStructure func(T) []S) {
	if len(v) <= 1 {
		return
	}

	data := make([]*collections.Pair[T, []S], len(v))
	foo := make([]collections.Pair[T, []S], len(v))
	for i := range v {
		foo[i] = collections.MakePair[T, []S](v[i], getStructure(v[i]))
		data[i] = &foo[i]
	}

	// get size of structure
	J := (int(unsafe.Sizeof(v[0])) * 8) - 1
	B := len(data[0].Second) - 1
	b := 0
	j := J

	var rec func(L, R int)
	rec = func(L, R int) {
		l := L
		r := R
		for l <= r {
			if data[l].Second[b]&(1<<j) == 0 {
				l++
			} else {
				data[l], data[r] = data[r], data[l]
				r--
			}
		}
		// The loop finish always in a situation like:
		// ...0000011111...
		// ...L---rl---R...
		if j > 0 {
			j--
			if L < r {
				rec(L, r)
			}
			if l < R {
				rec(l, R)
			}
			j++
		} else {
			j = J
			if b < B {
				b++
				if L < r {
					rec(L, r)
				}
				if l < R {
					rec(l, R)
				}
				b--
			}
			j = 0
		}
	}
	rec(0, len(data)-1)
	for i := range data {
		v[i] = data[i].First
	}
}
