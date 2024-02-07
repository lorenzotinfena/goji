package collections

type Set[T any] interface {
	Add(T)
	Remove(T)
	Contains(T) bool
}
