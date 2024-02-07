package utils

import (
	"github.com/lorenzotinfena/goji/utils/constraints"
)

func Prioritize[T constraints.Ordered]() func(a, b T) bool {
	return func(a, b T) bool {
		return a < b
	}
}

func Equalize[T comparable]() func(a, b T) bool {
	return func(a, b T) bool {
		return a == b
	}
}

func Multiplize[T constraints.Integer | constraints.Float]() func(a, b T) T {
	return func(a, b T) T {
		return a * b
	}
}
