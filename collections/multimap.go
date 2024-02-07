package collections

import "github.com/lorenzotinfena/goji/utils"

type MultiMap[K any, V any] interface {
	Add(K, V)
	GetOne(K) V
	Contains(K) bool
	Remove(K, V)
	Iterator(K) utils.Iterator[V]
}
