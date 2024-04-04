package linkedlist

import (
	"fmt"

	"github.com/lorenzotinfena/goji/utils"
)

type CircularDoublyLinkedListNode[T any] struct {
	Value      T
	Prev, Next *CircularDoublyLinkedListNode[T]
}

type CircularDoublyLinkedList[T any] struct {
	first  *CircularDoublyLinkedListNode[T]
	length int
	equals func(T, T) bool
}

// equals can be nil
func NewCircularDoublyLinkedList[T any](equals func(T, T) bool) *CircularDoublyLinkedList[T] {
	return &CircularDoublyLinkedList[T]{
		first:  nil,
		length: 0,
		equals: equals,
	}
}

func (l *CircularDoublyLinkedList[T]) Len() int { return l.length }

func (l *CircularDoublyLinkedList[T]) First() T { return l.first.Value }

func (l *CircularDoublyLinkedList[T]) Last() T { return l.first.Prev.Value }

func (l *CircularDoublyLinkedList[T]) FirstNode() *CircularDoublyLinkedListNode[T] { return l.first }

func (l *CircularDoublyLinkedList[T]) LastNode() *CircularDoublyLinkedListNode[T] {
	return l.first.Prev
}

func (l *CircularDoublyLinkedList[T]) InsertFirst(value T) {
	l.InsertLast(value)
	l.first = l.first.Prev
}

func (l *CircularDoublyLinkedList[T]) InsertLast(value T) {
	if l.length == 0 {
		l.first = &CircularDoublyLinkedListNode[T]{
			Value: value,
			Prev:  nil,
			Next:  nil,
		}
		l.first.Next = l.first
		l.first.Prev = l.first
		l.length++
		return
	}

	node := &CircularDoublyLinkedListNode[T]{
		Value: value,
		Prev:  l.first.Prev,
		Next:  l.first,
	}
	node.Prev.Next = node
	l.first.Prev = node
	l.length++
}

func (l *CircularDoublyLinkedList[T]) InserBefore(node *CircularDoublyLinkedListNode[T], value T) {
	if node == l.first {
		l.InsertFirst(value)
		return
	}

	toAdd := &CircularDoublyLinkedListNode[T]{
		Value: value,
		Prev:  node.Prev,
		Next:  node,
	}
	node.Prev.Next = toAdd
	node.Prev = toAdd
	l.length++
}

func (l *CircularDoublyLinkedList[T]) InsertAfter(node *CircularDoublyLinkedListNode[T], value T) {
	toAdd := &CircularDoublyLinkedListNode[T]{
		Value: value,
		Prev:  node,
		Next:  node.Next,
	}
	node.Next.Prev = toAdd
	node.Next = toAdd
	l.length++
}

// merge another ll after the last element
func (l *CircularDoublyLinkedList[T]) MergeEnd(ll *CircularDoublyLinkedList[T]) {
	if ll.Len() == 0 {
		return
	}
	l.length += ll.Len()
	if l.Len() == 0 {
		*l = *ll
		return
	}

	l.first.Prev.Next = ll.first
	ll.first.Prev.Next = l.first
	l.first.Prev, ll.first.Prev = ll.first.Prev, l.first.Prev
}

// index <= length
func (l *CircularDoublyLinkedList[T]) InsertAt(index int, value T) {
	if index == 0 {
		l.InsertFirst(value)
		return
	}

	var node *CircularDoublyLinkedListNode[T]
	if index <= l.Len()/2 {
		node = l.first
		for index > 1 {
			node = node.Next
			index--
		}
	} else {
		node = l.first.Prev
		for index < l.Len() {
			node = node.Prev
			index++
		}
	}

	toAdd := &CircularDoublyLinkedListNode[T]{
		Value: value,
		Prev:  node,
		Next:  node.Next,
	}
	node.Next.Prev = toAdd
	node.Next = toAdd
	l.length++
}

// Assumptions:
// - equals != nil
func (l *CircularDoublyLinkedList[T]) Contains(value T) bool {
	tmp := l.first
	for i := 0; i < l.Len(); i++ {
		if l.equals(tmp.Value, value) {
			return true
		}
		tmp = tmp.Next
	}
	return false
}

func (l *CircularDoublyLinkedList[T]) GetElementEqualsTo(value T) (T, bool) {
	tmp := l.first
	for i := 0; i < l.Len(); i++ {
		if l.equals(tmp.Value, value) {
			return tmp.Value, true
		}
		tmp = tmp.Next
	}
	return value, false
}

func (l *CircularDoublyLinkedList[T]) Clear() {
	l.first = nil
	l.length = 0
}

// index < length
func (l *CircularDoublyLinkedList[T]) GetNodeAt(index int) *CircularDoublyLinkedListNode[T] {
	n := l.first
	for index > 0 {
		n = n.Next
		index--
	}
	return n
}

// index < length
func (l *CircularDoublyLinkedList[T]) GetElementAt(index int) T {
	n := l.first
	for index > 0 {
		n = n.Next
		index--
	}
	return n.Value
}

// index < length
func (l *CircularDoublyLinkedList[T]) SetElementAt(index int, value T) {
	var node *CircularDoublyLinkedListNode[T]
	if index <= l.Len()/2 {
		node = l.first
		for index > 0 {
			node = node.Next
			index--
		}
	} else {
		node = l.first
		for index < l.Len() {
			node = node.Prev
			index++
		}
	}
	node.Value = value
}

func (l *CircularDoublyLinkedList[T]) RemoveFirst() T {
	value := l.first.Value
	if l.Len() == 1 {
		l.first = nil
	} else {
		l.first.Prev.Next = l.first.Next
		l.first.Next.Prev = l.first.Prev
	}
	l.length--
	return value
}

func (l *CircularDoublyLinkedList[T]) RemoveLast() (value T) {
	value = l.first.Prev.Value
	if l.Len() == 1 {
		l.first = nil
	} else {

		l.first.Prev.Prev.Next = l.first
		l.first.Prev = l.first.Prev.Prev
	}
	l.length--
	return
}

// index < length
func (l *CircularDoublyLinkedList[T]) RemoveAt(index int) T {
	if index == 0 {
		return l.RemoveFirst()
	}
	var node *CircularDoublyLinkedListNode[T]
	if index <= l.Len()/2 {
		node = l.first
		for index > 0 {
			node = node.Next
			index--
		}
	} else {
		node = l.first
		for index < l.Len() {
			node = node.Prev
			index++
		}
	}
	l.length--
	node.Next = node.Prev
	node.Prev.Next = node.Next
	return node.Value
}

func (l *CircularDoublyLinkedList[T]) Remove(element T) bool {
	if l.Len() == 0 {
		return false
	}
	if l.equals(l.first.Value, element) {
		l.RemoveFirst()
		return true
	}
	tmp := l.first.Next
	for i := 1; i < l.Len(); i++ {
		if l.equals(tmp.Value, element) {
			l.length--
			return true
		}
		tmp = tmp.Next
	}
	return false
}

func (l *CircularDoublyLinkedList[T]) ToSlice() []T {
	res := make([]T, l.length)
	tmp := l.first
	for i := 0; i < l.length; i++ {
		res[i] = tmp.Value
		tmp = tmp.Next
	}
	return res
}

func (l *CircularDoublyLinkedList[T]) GetIterator() utils.Iterator[T] {
	return &circularDoublyLinkedListIterator[T]{
		current:   l.first,
		remaining: l.Len(),
	}
}

func (it CircularDoublyLinkedList[T]) String() string {
	return fmt.Sprint(it.ToSlice())
}

type circularDoublyLinkedListIterator[T any] struct {
	current   *CircularDoublyLinkedListNode[T]
	remaining int
}

func (it *circularDoublyLinkedListIterator[T]) HasNext() bool {
	return it.remaining != 0
}

func (it *circularDoublyLinkedListIterator[T]) Next() T {
	tmp := it.current.Value
	it.current = it.current.Next
	it.remaining--
	return tmp
}
