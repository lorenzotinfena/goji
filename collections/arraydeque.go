package collections

type ArrayDeque[T any] struct {
	v     []T
	start int
	end   int
}

func NewArrayDeque[T any]() *ArrayDeque[T] {
	return &ArrayDeque[T]{
		[]T{}, -1, -1,
	}
}

func (arr *ArrayDeque[T]) doubleSize() {
	v1 := make([]T, len(arr.v)*2)
	j := 0
	for i := arr.start; i != arr.end; i = (i + 1) % len(arr.v) {
		v1[j] = arr.v[i]
		j++
	}
	v1[j] = arr.v[arr.end]
	arr.start = 0
	arr.end = len(arr.v) - 1
	arr.v = v1
}

func (arr *ArrayDeque[T]) Len() int {
	if arr.start <= arr.end {
		return arr.end - arr.start + 1
	} else {
		return arr.end + len(arr.v) - arr.start + 1
	}
}

func (arr *ArrayDeque[T]) Get(index int) T {
	return arr.v[(arr.start+index)%len(arr.v)]
}

func (arr *ArrayDeque[T]) Set(index int, element T) {
	arr.v[(arr.start+index)%len(arr.v)] = element
}

func (arr *ArrayDeque[T]) InsertFirst(element T) {
	if len(arr.v) == 0 {
		arr.v = []T{element}
		arr.start = 0
		arr.end = 0
		return
	}
	if (arr.end+1)%len(arr.v) == arr.start {
		arr.doubleSize()
	}
	if arr.start == 0 {
		arr.start = len(arr.v) - 1
	} else {
		arr.start = (arr.start - 1) % len(arr.v)
	}
	arr.v[arr.end] = element
}

func (arr *ArrayDeque[T]) InsertLast(element T) {
	if len(arr.v) == 0 {
		arr.v = []T{element}
		arr.start = 0
		arr.end = 0
		return
	}
	if (arr.end+1)%len(arr.v) == arr.start {
		arr.doubleSize()
	}

	arr.end = (arr.end + 1) % len(arr.v)
	arr.v[arr.end] = element
}

func (arr *ArrayDeque[T]) shrink() {
	if 4*arr.Len() <= len(arr.v) {
		nextV := make([]T, len(arr.v))
		j := 0
		for i := arr.start; i <= arr.end; i = (i + 1) % len(arr.v) {
			nextV[j] = arr.v[i]
			j++
		}
		arr.v = nextV
	}
}

func (arr *ArrayDeque[T]) RemoveFirst() {
	if arr.end == 0 {
		arr.end = len(arr.v) - 1
	} else {
		arr.end--
	}
	arr.shrink()
}

func (arr *ArrayDeque[T]) RemoveLast() {
	if arr.start == len(arr.v)-1 {
		arr.start = 0
	} else {
		arr.start++
	}
	arr.shrink()
}

func (arr *ArrayDeque[T]) Slice() []T {
	nextV := make([]T, len(arr.v))
	j := 0
	for i := arr.start; i <= arr.end; i = (i + 1) % len(arr.v) {
		nextV[j] = arr.v[i]
		j++
	}
	return nextV
}

func (arr *ArrayDeque[T]) Insert(element T, index int) {
	if index == 0 {
		arr.InsertLast(element)
		return
	}
	if (arr.end+1)%len(arr.v) == arr.start {
		arr.doubleSize()
	}
	i := (arr.start + index) % len(arr.v)
	for ; i <= arr.end; i = (i + 1) % len(arr.v) {
		element, arr.v[i] = arr.v[i], element
	}
	arr.v[i] = element
}
