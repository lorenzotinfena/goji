package linkedlist

import (
	"fmt"

	"github.com/lorenzotinfena/goji/utils"
)

type SinglyLinkedListNode[T any] struct {
	Value T
	Next  *SinglyLinkedListNode[T]
}

type SinglyLinkedList[T any] struct {
	first  *SinglyLinkedListNode[T]
	last   *SinglyLinkedListNode[T]
	length int
	equals func(T, T) bool
}

// equals can be nil
func NewSinglyLinkedList[T any](equals func(T, T) bool) *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{
		first:  nil,
		last:   nil,
		length: 0,
		equals: equals,
	}
}
func (l *SinglyLinkedList[T]) Len() int { return l.length }

func (l *SinglyLinkedList[T]) First() T { return l.first.Value }

func (l *SinglyLinkedList[T]) Last() T { return l.last.Value }

func (l *SinglyLinkedList[T]) FirstNode() *SinglyLinkedListNode[T] { return l.first }

func (l *SinglyLinkedList[T]) LastNode() *SinglyLinkedListNode[T] { return l.last }

func (l *SinglyLinkedList[T]) InsertFirst(value T) {
	if l.length == 0 {
		nodeToInsert := &SinglyLinkedListNode[T]{
			Value: value,
			Next:  l.first,
		}
		l.first = nodeToInsert
		l.last = nodeToInsert
	} else {
		l.first = &SinglyLinkedListNode[T]{
			Value: value,
			Next:  l.first,
		}
	}
	l.length++
}

func (l *SinglyLinkedList[T]) InsertLast(value T) {
	if l.length == 0 {
		nodeToInsert := &SinglyLinkedListNode[T]{
			Value: value,
			Next:  l.first,
		}
		l.first = nodeToInsert
		l.last = nodeToInsert
	} else {
		l.last.Next = &SinglyLinkedListNode[T]{
			Value: value,
			Next:  nil,
		}
		l.last = l.last.Next
	}
	l.length++
}

func (l *SinglyLinkedList[T]) InsertAfter(node *SinglyLinkedListNode[T], value T) {
	if node == l.last {
		l.InsertLast(value)
		return
	}

	toAdd := &SinglyLinkedListNode[T]{
		Value: value,
		Next:  node.Next,
	}
	node.Next = toAdd
	l.length++
}

// merge another ll after the last element
func (l *SinglyLinkedList[T]) MergeEnd(ll *SinglyLinkedList[T]) {
	if l.Len() != 0 {
		if ll.Len() != 0 {
			l.length += ll.Len()
			l.last.Next = ll.first
			l.last = ll.last
		}
	} else {
		*l = *ll
	}
}

// index <= length
func (l *SinglyLinkedList[T]) InsertAt(index int, value T) {
	if index == 0 {
		l.InsertFirst(value)
		return
	}
	if index == l.length {
		l.InsertLast(value)
		return
	}

	n := l.first
	for index > 1 {
		n = n.Next
		index--
	}
	n.Next = &SinglyLinkedListNode[T]{
		Value: value,
		Next:  n.Next,
	}
	l.length++
}

// Assumptions:
// - equals != nil
func (l *SinglyLinkedList[T]) Contains(value T) bool {
	tmp := l.first
	for i := 0; i < l.length; i++ {
		if l.equals(tmp.Value, value) {
			return true
		}
		tmp = tmp.Next
	}
	return false
}

func (l *SinglyLinkedList[T]) GetElementEqualsTo(value T) (T, bool) {
	tmp := l.first
	for i := 0; i < l.length; i++ {
		if l.equals(tmp.Value, value) {
			return tmp.Value, true
		}
		tmp = tmp.Next
	}
	return value, false
}

func (l *SinglyLinkedList[T]) Clear() {
	l.first = nil
	l.last = nil
	l.length = 0
}

// index < length
func (l *SinglyLinkedList[T]) GetElementAt(index int) T {
	n := l.first
	for index > 0 {
		n = n.Next
		index--
	}
	return n.Value
}

// index < length
func (l *SinglyLinkedList[T]) SetElementAt(index int, value T) {
	n := l.first
	for index > 0 {
		n = n.Next
		index--
	}
	n.Value = value
}

func (l *SinglyLinkedList[T]) RemoveFirst() T {
	tmp := l.first
	l.first = l.first.Next
	l.length--
	if l.first == nil {
		l.last = nil
	}
	return tmp.Value
}

func (l *SinglyLinkedList[T]) RemoveLast() (value T) {
	value = l.last.Value
	if l.length == 1 {
		l.first = nil
		l.last = nil
		l.length = 0
		return
	}

	tmp := l.first
	for i := 2; i < l.length; i++ {
		tmp = tmp.Next
	}
	l.last = tmp
	l.length--
	return
}

// index < length
func (l *SinglyLinkedList[T]) RemoveAt(index int) T {
	if index == 0 {
		return l.RemoveFirst()
	}
	if index == l.length-1 {
		return l.RemoveLast()
	}

	tmp := l.first
	for i := 1; i < index; i++ {
		tmp = tmp.Next
	}
	res := tmp.Next.Value
	tmp.Next = tmp.Next.Next
	l.length--
	return res
}

func (l *SinglyLinkedList[T]) Remove(element T) bool {
	if l.Len() == 0 {
		return false
	}

	if l.equals(l.First(), element) {
		l.RemoveFirst()
		return true
	}
	if l.equals(l.Last(), element) {
		l.RemoveLast()
		return true
	}
	tmp := l.first
	for i := 2; i < l.Len(); i++ {
		if l.equals(tmp.Next.Value, element) {
			tmp.Next = tmp.Next.Next
			return true
		}
		tmp = tmp.Next
	}
	return false
}

func (l *SinglyLinkedList[T]) ToSlice() (res []T) {
	res = make([]T, 0, l.length)
	tmp := l.first
	for i := 0; i < l.length; i++ {
		res = append(res, tmp.Value)
		tmp = tmp.Next
	}
	return
}

func (l *SinglyLinkedList[T]) GetIterator() utils.Iterator[T] {
	return &singlyLinkedListIterator[T]{
		current: l.first,
	}
}

func (it SinglyLinkedList[T]) String() string {
	return fmt.Sprint(it.ToSlice())
}

type singlyLinkedListIterator[T any] struct {
	current *SinglyLinkedListNode[T]
}

func (it *singlyLinkedListIterator[T]) HasNext() bool {
	return it.current != nil
}

func (it *singlyLinkedListIterator[T]) Next() T {
	tmp := it.current.Value
	it.current = it.current.Next
	return tmp
}
