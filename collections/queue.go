package collections

import (
	ll "github.com/lorenzotinfena/goji/collections/linkedlist"
)

type Queue[T any] struct {
	l ll.SinglyLinkedList[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		l: *ll.NewSinglyLinkedList[T](nil),
	}
}
func (q *Queue[T]) Len() int        { return q.l.Len() }
func (q *Queue[T]) Enqueue(value T) { q.l.InsertLast(value) }
func (q *Queue[T]) Dequeue() T      { return q.l.RemoveFirst() }
func (q *Queue[T]) Preview() T      { return q.l.First() }
func (q Queue[T]) String() string   { return q.l.String() }
