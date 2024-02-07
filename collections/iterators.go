package collections

type dfsIterator[T comparable] struct {
	getNexts  func(T) []T
	toAnalyze Stack[T]
	visited   HashSet[T]
}

func (it *dfsIterator[T]) HasNext() bool {
	return it.toAnalyze.Len() != 0
}

func (it *dfsIterator[T]) Next() T {
	cur := it.toAnalyze.Pop()

	it.visited.Add(cur)
	nexts := it.getNexts(cur)
	for _, v := range nexts {
		if !it.visited.Contains(v) {
			it.toAnalyze.Push(v)
		}
	}
	for it.toAnalyze.Len() != 0 && it.visited.Contains(it.toAnalyze.Preview()) {
		it.toAnalyze.Pop()
	}
	return cur
}

func NewIteratorDFS[T comparable](root T, getNexts func(T) []T) *dfsIterator[T] {
	toAnalyze := *NewStack[T]()
	toAnalyze.Push(root)
	visited := *NewHashSet[T]()

	return &dfsIterator[T]{
		getNexts:  getNexts,
		toAnalyze: toAnalyze,
		visited:   visited,
	}
}

type bfsIterator[T comparable] struct {
	getNexts  func(T) []T
	toAnalyze Queue[T]
	visited   HashSet[T]
}

func (it *bfsIterator[T]) HasNext() bool {
	return it.toAnalyze.Len() != 0
}

func (it *bfsIterator[T]) Next() T {
	cur := it.toAnalyze.Dequeue()
	nexts := it.getNexts(cur)
	for _, v := range nexts {
		if !it.visited.Contains(v) {
			it.toAnalyze.Enqueue(v)
			it.visited.Add(v)
		}
	}
	return cur
}

func NewIteratorBFS[T comparable](root T, getNexts func(T) []T) *bfsIterator[T] {
	toAnalyze := *NewQueue[T]()
	toAnalyze.Enqueue(root)
	visited := *NewHashSet[T]()
	visited.Add(root)

	return &bfsIterator[T]{
		getNexts:  getNexts,
		toAnalyze: toAnalyze,
		visited:   *NewHashSet[T](),
	}
}
