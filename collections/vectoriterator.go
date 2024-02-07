package collections

type VectorIterator[T any] struct {
	v []T
	i int
}

func NewVectorIterator[T any](v []T) *VectorIterator[T] {
	return &VectorIterator[T]{v, 0}
}

func (h *VectorIterator[T]) HasNext() bool {
	return h.i < len(h.v)
}

func (h *VectorIterator[T]) Next() T {
	tmp := h.v[h.i]
	h.i++
	return tmp
}
