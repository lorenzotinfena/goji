package utils

type Iterator[T any] interface {
	HasNext() bool
	Next() T
}
