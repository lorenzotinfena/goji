package collections_test

import (
	"testing"

	"github.com/lorenzotinfena/goji/collections"
)

func TestArrayDeque(t *testing.T) {
	deque := collections.NewArrayDeque[int]()
	deque.InsertLast(1)
	deque.InsertLast(1)
	deque.InsertFirst(1)
	deque.InsertFirst(1)
	deque.RemoveFirst()
	deque.RemoveFirst()
	deque.RemoveFirst()
	deque.RemoveFirst()
	deque.InsertFirst(1)
	deque.InsertLast(1)
	deque.InsertFirst(1)
	deque.InsertLast(1)
	deque.InsertLast(1)
	deque.RemoveLast()
	deque.RemoveLast()
	deque.RemoveLast()
	deque.RemoveLast()
	deque.RemoveLast()
}
