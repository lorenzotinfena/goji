package tree

import "fmt"

// The main difference is pendingUpdates which holds for each segment its pending updates,
// When mergeUpdates is set all pending updates for each node are continuously merged,
// Otherwise they are all queued
type LazySegmentTree[Q comparable] struct {
	mergeForQuery func(q1 Q, q2 Q) Q
	mergeUpdates  func(func(Q) Q, func(Q) Q) func(Q) Q

	NumberElements int
	PendingUpdates [][]func(Q) Q
	Segments       []Q
}

// Pass nil as mergeUpdates if you never call UpdateRange or if you can't merge updates
// Assumptions:
// - len(elements) > 0
// - query != nil
// - mergeUpdates != nil
func NewLazySegmentTree[E any, Q comparable](
	elements []E,
	query func(element E) Q,
	mergeForQuery func(q1 Q, q2 Q) Q,
	mergeSegments func(q1 Q, q2 Q) Q,
	mergeUpdates func(func(Q) Q, func(Q) Q) func(Q) Q,
) *LazySegmentTree[Q] {
	segments := make([]Q, 4*len(elements))

	var build func(i, l, r int)
	build = func(i, l, r int) {
		if l == r {
			segments[i] = query(elements[l])
			return
		}
		m := (l + r) / 2
		build(2*i+1, l, m)
		build(2*i+2, m+1, r)
		segments[i] = mergeSegments(segments[2*i+1], segments[2*i+2])
	}
	build(0, 0, len(elements)-1)

	pendingUpdates := make([][]func(Q) Q, 4*len(elements))
	for i := 0; i < 4*len(elements); i++ {
		pendingUpdates[i] = make([]func(Q) Q, 0)
	}

	return &LazySegmentTree[Q]{
		mergeForQuery: mergeForQuery,
		mergeUpdates:  mergeUpdates,

		NumberElements: len(elements),
		PendingUpdates: pendingUpdates,
		Segments:       segments,
	}
}

func (s *LazySegmentTree[Q]) insertPendingUpdate(i int, update func(Q) Q) {
	if s.mergeUpdates == nil {
		s.PendingUpdates[i] = append(s.PendingUpdates[i], update)
	} else {
		if len(s.PendingUpdates[i]) == 0 {
			s.PendingUpdates[i] = append(s.PendingUpdates[i], update)
		} else {
			s.PendingUpdates[i][0] = s.mergeUpdates(s.PendingUpdates[i][0], update)
		}
	}
}

func (s *LazySegmentTree[Q]) push(i, l, r int) {
	for _, f := range s.PendingUpdates[i] {
		s.Segments[i] = f(s.Segments[i])
	}
	if l != r {
		for _, f := range s.PendingUpdates[i] {
			s.insertPendingUpdate(2*i+1, f)
			s.insertPendingUpdate(2*i+2, f)
		}
	}
	s.PendingUpdates[i] = make([]func(Q) Q, 0)
}

// Performs a query to [start, end]
// Assumptions:
//   - start <= end
//   - start and end are valid
func (s *LazySegmentTree[Q]) Query(start int, end int) Q {
	var queryRecRight func(i, l, r int) Q
	queryRecRight = func(i, l, r int) Q {
		s.push(i, l, r)
		if r == end {
			return s.Segments[i]
		}
		m := (l + r) / 2
		if end <= m {
			return queryRecRight(2*i+1, l, m)
		} else {
			return s.mergeForQuery(
				s.Segments[2*i+1],
				queryRecRight(2*i+2, m+1, r),
			)
		}
	}
	var queryRecLeft func(i, l, r int) Q
	queryRecLeft = func(i, l, r int) Q {
		s.push(i, l, r)
		m := (l + r) / 2
		if start >= m+1 {
			return queryRecLeft(2*i+2, m+1, r)
		} else {
			if l == start {
				return s.Segments[i]
			}
			return s.mergeForQuery(
				queryRecLeft(2*i+1, l, m),
				s.Segments[2*i+2],
			)
		}
	}
	var queryRec func(i, l, r int) Q
	queryRec = func(i, l, r int) Q {
		s.push(i, l, r)
		if l == r {
			return s.Segments[i]
		}
		m := (l + r) / 2
		if end <= m {
			return queryRec(2*i+1, l, m)
		} else if start >= m+1 {
			return queryRec(2*i+2, m+1, r)
		} else {
			return s.mergeForQuery(
				queryRecLeft(2*i+1, l, m),
				queryRecRight(2*i+2, m+1, r),
			)
		}
	}
	return queryRec(0, 0, s.NumberElements-1)
}

// Assumptions:
//   - start and end are valid
//   - update != nil
func (s *LazySegmentTree[Q]) UpdateRange(start, end int, update func(Q) Q, mergeSegments func(left, right, old Q) Q) {
	var updateRecRight func(i, l, r int)
	updateRecRight = func(i, l, r int) {
		if r == end {
			s.insertPendingUpdate(i, update)
			s.push(i, l, r)
			return
		}

		s.push(i, l, r)
		m := (l + r) / 2
		if end <= m {
			updateRecRight(2*i+1, l, m)
			s.push(2*i+2, m+1, r)
		} else {
			s.insertPendingUpdate(2*i+1, update)
			s.push(2*i+1, l, m)
			updateRecRight(2*i+2, m+1, r)
		}
		s.Segments[i] = mergeSegments(s.Segments[2*i+1], s.Segments[2*i+2], s.Segments[i])
	}
	var updateRecLeft func(i, l, r int)
	updateRecLeft = func(i, l, r int) {
		if l == start {
			s.insertPendingUpdate(i, update)
			s.push(i, l, r)
			return
		}
		s.push(i, l, r)
		m := (l + r) / 2
		if start >= m+1 {
			updateRecLeft(2*i+2, m+1, r)
			s.push(2*i+1, l, m)
		} else {
			s.insertPendingUpdate(2*i+2, update)
			s.push(2*i+2, m+1, r)
			updateRecLeft(2*i+1, l, m)
		}
		s.Segments[i] = mergeSegments(s.Segments[2*i+1], s.Segments[2*i+2], s.Segments[i])
	}
	var updateRec func(i, l, r int)
	updateRec = func(i, l, r int) {
		if start == l && end == r {
			s.insertPendingUpdate(i, update)
			s.push(i, l, r)
			return
		}
		s.push(i, l, r)
		m := (l + r) / 2
		if end <= m {
			updateRec(2*i+1, l, m)
			s.push(2*i+2, m+1, r)
		} else if start >= m+1 {
			updateRec(2*i+2, m+1, r)
			s.push(2*i+1, l, m)
		} else {
			updateRecLeft(2*i+1, l, m)
			updateRecRight(2*i+2, m+1, r)
		}
		s.Segments[i] = mergeSegments(s.Segments[2*i+1], s.Segments[2*i+2], s.Segments[i])
	}
	updateRec(0, 0, s.NumberElements-1)
}

func (s *LazySegmentTree[Q]) Len() int {
	return s.NumberElements
}

func (s *LazySegmentTree[Q]) String() string {
	var rec func(l, r int) *treeNode[string]
	rec = func(l, r int) *treeNode[string] {
		node := &treeNode[string]{
			Value: fmt.Sprint(l) + "━━━" + fmt.Sprint(r) + "\n" + fmt.Sprint(s.Query(l, r)),
		}
		if l != r {
			m := (l + r) / 2
			node.Children = []*treeNode[string]{rec(l, m), rec(m+1, r)}
		}
		return node
	}
	return rec(0, s.Len()-1).String()
}

func (s *LazySegmentTree[Q]) RightMost(left int, merge func(Q, Q) Q, predicate func(Q) bool) int {
	var recRight func(i, l, r int, q Q) (int, Q)
	recRight = func(i, l, r int, q Q) (int, Q) {
		s.push(i, l, r)
		if l == r {
			tmp := merge(q, s.Segments[i])
			if predicate(tmp) {
				return l + 1, tmp
			} else {
				return l, q
			}
		}
		m := (l + r) / 2
		s.push(2*i+1, l, m)
		tmp := merge(q, s.Segments[2*i+1])
		if predicate(tmp) {
			return recRight(2*i+2, m+1, r, tmp)
		} else {
			return recRight(2*i+1, l, m, q)
		}
	}
	var rec func(i, l, r int) (int, Q)
	rec = func(i, l, r int) (int, Q) {
		s.push(i, l, r)
		if l == r {
			if predicate(s.Segments[i]) {
				return l + 1, s.Segments[i]
			} else {
				var tmp Q
				return l, tmp
			}
		}
		m := (l + r) / 2
		if left <= m {
			rightMostAsFar, q := rec(2*i+1, l, m)
			if rightMostAsFar == m+1 {
				s.push(2*i+2, m+1, r)
				tmp := merge(q, s.Segments[2*i+2])
				if predicate(tmp) {
					return r + 1, tmp
				} else {
					return recRight(2*i+2, m+1, r, q)
				}
			} else {
				return rightMostAsFar, q
			}
		} else {
			return rec(2*i+2, m+1, r)
		}
	}
	rightMost, _ := rec(0, 0, s.NumberElements-1)
	return rightMost
}

func (s *LazySegmentTree[Q]) LeftMost(right int, merge func(Q, Q) Q, predicate func(Q) bool) int {
	var recLeft func(i, l, r int, q Q) (int, Q)
	recLeft = func(i, l, r int, q Q) (int, Q) {
		s.push(i, l, r)
		if l == r {
			tmp := merge(q, s.Segments[i])
			if predicate(tmp) {
				return l - 1, tmp
			} else {
				return l, q
			}
		}
		m := (l + r) / 2
		s.push(2*i+2, m+1, r)
		tmp := merge(q, s.Segments[2*i+2])
		if predicate(tmp) {
			return recLeft(2*i+1, l, m, tmp)
		} else {
			return recLeft(2*i+2, m+1, r, q)
		}
	}
	var rec func(i, l, r int) (int, Q)
	rec = func(i, l, r int) (int, Q) {
		s.push(i, l, r)
		if l == r {
			if predicate(s.Segments[i]) {
				return l - 1, s.Segments[i]
			} else {
				var tmp Q
				return l, tmp
			}
		}
		m := (l + r) / 2
		if right > m {
			leftMostAsFar, q := rec(2*i+2, m+1, r)
			if leftMostAsFar == m {
				s.push(2*i+1, l, m)
				tmp := merge(q, s.Segments[2*i+1])
				if predicate(tmp) {
					return l - 1, tmp
				} else {
					return recLeft(2*i+1, l, m, q)
				}
			} else {
				return leftMostAsFar, q
			}
		} else {
			return rec(2*i+1, l, m)
		}
	}
	rightMost, _ := rec(0, 0, s.NumberElements-1)
	return rightMost
}
