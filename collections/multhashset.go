package collections

import (
	"fmt"
)

type MultiHashSet[T comparable] struct {
	m      HashMap[T, int]
	length int
}

func NewMultiHashSet[T comparable]() *MultiHashSet[T] {
	return &MultiHashSet[T]{m: *NewHashMap[T, int]()}
}

func (s *MultiHashSet[T]) Add(element T) {
	if !s.m.Contains(element) {
		s.m.Set(element, 1)
	} else {
		s.m.Set(element, s.m.Get(element)+1)
	}
	s.length++
}

func (s *MultiHashSet[T]) Remove(element T) {
	if !s.m.Contains(element) || s.m.Get(element) == 1 {
		s.m.Remove(element)
	} else {
		s.m.Set(element, s.m.Get(element)-1)
	}
	s.length--
}

func (s *MultiHashSet[T]) MultiplicityOf(element T) int {
	if s.m.Contains(element) {
		return s.m.Get(element)
	} else {
		return 0
	}
}

func (s MultiHashSet[T]) Contains(element T) bool {
	return s.m.Contains(element)
}

func (s MultiHashSet[T]) ItemsWithoutMultiplicity() []T {
	return s.m.Keys()
}

func (s MultiHashSet[T]) Items() []T {
	keys := make([]T, 0, s.length)
	for _, k := range s.m.Keys() {
		i := s.m.Get(k)
		for i > 0 {
			keys = append(keys, k)
			i--
		}
	}
	return keys
}

func (s MultiHashSet[T]) String() string {
	return fmt.Sprint(s.Items())
}

func (s MultiHashSet[T]) Len() int {
	return s.length
}

func (s MultiHashSet[T]) LenWithoutMultiplicity() int {
	return s.m.Len()
}
