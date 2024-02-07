package heap

type BinaryHeap[T any] struct {
	v     []T
	prior func(T, T) bool
}

// Note for function Prior: It's a strict order relation
func NewBinaryHeapFromSlice[T any](
	v []T,
	prior func(T, T) bool,
) (h *BinaryHeap[T]) {
	h = &BinaryHeap[T]{
		v:     v,
		prior: prior,
	}
	if h.Len() > 1 {
		for i := (int64(h.Len()) - 2) / 2; i >= 0; i-- {
			h.heapifyDown(int64(i))
		}
	}
	return
}

func NewBinaryHeap[T any](Prior func(T, T) bool) *BinaryHeap[T] {
	return &BinaryHeap[T]{v: make([]T, 0), prior: Prior}
}

func (h *BinaryHeap[T]) Len() uint64 {
	return uint64(len(h.v))
}

func (h *BinaryHeap[T]) Push(value T) {
	h.v = append(h.v, value)
	h.heapifyUp(int64(h.Len() - 1))
}

func (h *BinaryHeap[T]) Pop() (res T) {
	res = h.v[0]
	h.v[0] = h.v[h.Len()-1]
	h.v = h.v[:h.Len()-1]
	h.heapifyDown(0)
	return
}

func (h *BinaryHeap[T]) heapifyDown(index int64) bool {
	origin := index
	for {
		j := index*2 + 2
		if j < int64(h.Len()) {
			if h.prior(h.v[j-1], h.v[j]) {
				j--
			}
		} else {
			j--
			if j >= int64(h.Len()) {
				break
			}
		}
		if h.prior(h.v[j], h.v[index]) {
			h.v[j], h.v[index] = h.v[index], h.v[j]
			index = j
		} else {
			break
		}
	}
	return origin != index
}

func (h *BinaryHeap[T]) heapifyUp(index int64) {
	for {
		if index == 0 {
			break
		}
		parent := (index - 1) / 2
		if h.prior(h.v[parent], h.v[index]) {
			break
		}
		h.v[index], h.v[parent] = h.v[parent], h.v[index]
		index = parent
	}
}

func (h *BinaryHeap[T]) Preview() T {
	return h.v[0]
}

func (h BinaryHeap[T]) String() string {
	return "" // #TODO
}
