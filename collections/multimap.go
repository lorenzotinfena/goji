package collections

type MultiMap[K any, V any] interface {
	Add(K, V)
	GetOne(K) V
	Contains(K) bool
	Remove(K, V)
}
