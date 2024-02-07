// Package sqrt contains algorithms and data structures that contains a √n in their complexity
package sqrt

import "math"

// Sqrt (or Square Root) Decomposition is a technique used for query an array and perform updates
// Inside this package is described its most simple data structure, you can find more at: https://cp-algorithms.com/data_structures/sqrt_decomposition.html
//
// Formally, You can use SqrtDecomposition only if:
//
// Given a function $Query:E_1,...,E_n\rightarrow Q$
//
// if $\exist unionQ:Q,Q\rightarrow Q$
//
// s.t.
//
// - $\forall n\in \N > 1, 1\le i<n, E_1,..., E_n\in E \\ query(E_1,..., E_n)=unionQ(query(E_1,..., E_i), query(E_{i+1},...,E_n))$
//
// - (Only if you want use $update$ function)
// $\forall n\in \N > 0, E_1,..., E_n\in E \\ query(E_1,...,E_{new},..., E_n)=updateQ(query(E_1,...,E_{old},...,E_n), indexof(E_{old}), E_{new})$
type SqrtDecomposition[E any, Q any] struct {
	query  func(element E) Q
	merge  func(q1 Q, q2 Q) Q
	update func(oldQ Q, oldE E, newE E) (newQ Q)

	elements  []E
	blocks    []Q
	blockSize int
}

// Create a new SqrtDecomposition instance with the parameters as specified by SqrtDecomposition comment
// Pass nil as update if you never call Update method
// Assumptions:
//   - len(elements) > 0
func NewSqrtDecomposition[E any, Q any](
	elements []E,
	query func(element E) Q,
	merge func(q1 Q, q2 Q) Q,
	update func(oldQ Q, oldE E, newE E) (newQ Q),
) *SqrtDecomposition[E, Q] {
	sqrtDec := &SqrtDecomposition[E, Q]{
		query:    query,
		merge:    merge,
		update:   update,
		elements: elements,
	}
	sqrt := math.Sqrt(float64(len(sqrtDec.elements)))
	blockSize := int(sqrt)
	numBlocks := int(math.Ceil(float64(len(elements)) / float64(blockSize)))
	sqrtDec.blocks = make([]Q, numBlocks)
	for i := 0; i < len(elements); i++ {
		if i%blockSize == 0 {
			sqrtDec.blocks[i/blockSize] = sqrtDec.query(elements[i])
		} else {
			sqrtDec.blocks[i/blockSize] = sqrtDec.merge(sqrtDec.blocks[i/blockSize], sqrtDec.query(elements[i]))
		}
	}
	sqrtDec.blockSize = blockSize
	return sqrtDec
}

// Performs a query from index start to index end (non included)
// Assumptions:
//   - start < end
//   - start and end are valid
func (s *SqrtDecomposition[E, Q]) Query(start int, end int) Q {
	firstIndexNextBlock := ((start / s.blockSize) + 1) * s.blockSize
	q := s.query(s.elements[start])
	if firstIndexNextBlock > end { // if in same block
		start++
		for start < end {
			q = s.merge(q, s.query(s.elements[start]))
			start++
		}
	} else {
		// left side
		start++
		for start < firstIndexNextBlock {
			q = s.merge(q, s.query(s.elements[start]))
			start++
		}

		// middle part
		endBlock := end / s.blockSize
		for i := firstIndexNextBlock / s.blockSize; i < endBlock; i++ {
			q = s.merge(q, s.blocks[i])
		}

		// right part
		for i := endBlock * s.blockSize; i < end; i++ {
			q = s.merge(q, s.query(s.elements[i]))
		}
	}
	return q
}

// Assumptions:
//   - index is valid
func (s *SqrtDecomposition[E, Q]) Update(index int, newE E) {
	i := index / s.blockSize
	s.blocks[i] = s.update(s.blocks[i], s.elements[index], newE)
	s.elements[index] = newE
}
