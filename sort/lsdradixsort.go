package sort

import (
	"unsafe"

	"github.com/lorenzotinfena/goji/collections"
	"github.com/lorenzotinfena/goji/utils/constraints"
)

// Assumptions:
// - getStructure always returns slices of same length
func LSDRadixSort[T any, S constraints.Unsigned](v []T, getStructure func(T) []S) {
	if len(v) <= 1 {
		return
	}

	data := make([]*collections.Pair[T, []S], len(v))
	foo := make([]collections.Pair[T, []S], len(v))
	for i := range v {
		foo[i] = collections.MakePair[T, []S](v[i], getStructure(v[i]))
		data[i] = &foo[i]
	}

	res := make([]*collections.Pair[T, []S], len(v))

	// get size of structure
	J := (int(unsafe.Sizeof(v[0])) * 8) - 1

	for b := len(data[0].Second) - 1; b >= 0; b-- {
		for j := 0; j <= J; j++ {

			// empty lists
			zero := 0
			one := 0

			for _, item := range data {
				if item.Second[b]&(1<<j) == 0 {
					zero++
				} else {
					one++
				}
			}

			one += zero

			for i := len(data) - 1; i >= 0; i-- {
				if data[i].Second[b]&(1<<j) == 0 {
					zero--
					res[zero] = data[i]
				} else {
					one--
					res[one] = data[i]
				}
			}
			res, data = data, res
		}
	}
	for i := range data {
		v[i] = data[i].First
	}
}
