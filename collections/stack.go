package collections

import (
	ll "github.com/lorenzotinfena/goji/collections/linkedlist"
)

type Stack[T any] struct {
	l ll.SinglyLinkedList[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		l: *ll.NewSinglyLinkedList[T](nil),
	}
}

func (s *Stack[T]) Len() int      { return s.l.Len() }

func (s *Stack[T]) Push(value T)  { s.l.InsertFirst(value) }

func (s *Stack[T]) Pop() T        { return s.l.RemoveFirst() }

func (s *Stack[T]) Preview() T    { return s.l.First() }

func (s Stack[T]) String() string { return s.l.String() }
