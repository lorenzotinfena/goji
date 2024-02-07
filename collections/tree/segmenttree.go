package tree

import "fmt"

type SegmentTree[E any, Q comparable] struct {
	query         func(element E) Q
	mergeForQuery func(q1 Q, q2 Q) Q

	NumberElements int
	Elements       []E
	Segments       []Q
}

// Pass nil as update if you never call Update method
// Assumptions:
//   - len(elements) > 0
func NewSegmentTree[E any, Q comparable](
	elements []E,
	query func(element E) Q,
	mergeForQuery func(q1 Q, q2 Q) Q,
	mergeSegments func(left, right Q) Q,
) *SegmentTree[E, Q] {
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

	return &SegmentTree[E, Q]{
		query:         query,
		mergeForQuery: mergeForQuery,

		NumberElements: len(elements),
		Elements:       elements,
		Segments:       segments,
	}
}

// Performs a query to [start, end]
// Assumptions:
//   - start <= end
//   - start and end are valid
func (s *SegmentTree[E, Q]) Query(start int, end int) Q {
	var queryRecRight func(i, l, r int) Q
	queryRecRight = func(i, l, r int) Q {
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
	return queryRec(0, 0, len(s.Elements)-1)
}

// Assumptions:
//   - index is valid
func (s *SegmentTree[E, Q]) Update(index int, update func(Q) Q) {
	var updateRec func(i, l, r int)
	updateRec = func(i, l, r int) {
		s.Segments[i] = update(s.Segments[i])
		if l == r {
			return
		}
		m := (l + r) / 2
		if index <= m {
			updateRec(2*i+1, l, m)
		} else {
			updateRec(2*i+2, m+1, r)
		}
	}
	updateRec(0, 0, len(s.Elements)-1)
}

func (s *SegmentTree[E, Q]) Len() int {
	return len(s.Elements)
}

func (s *SegmentTree[E, Q]) String() string {
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

// Assumptions:
//   - index is valid
func (s *SegmentTree[E, Q]) UpdateRangeWithoutPropagation(start, end int, update func(old Q) Q, mergeSegments func(left, right *Q, old Q) Q) {
	var updateRecRight func(i, l, r int)
	updateRecRight = func(i, l, r int) {
		if r == end {
			s.Segments[i] = update(s.Segments[i])
			if l != r {
				s.Segments[i] = mergeSegments(&s.Segments[2*i+1], &s.Segments[2*i+2], s.Segments[i])
			} else {
				s.Segments[i] = mergeSegments(nil, nil, s.Segments[i])
			}
		}
		m := (l + r) / 2
		if end <= m {
			updateRecRight(2*i+1, l, m)
		} else {
			s.Segments[2*i+2] = update(s.Segments[2*i+2])
			updateRecRight(2*i+2, m+1, r)
		}
		s.Segments[i] = mergeSegments(&s.Segments[2*i+1], &s.Segments[2*i+2], s.Segments[i])
	}
	var updateRecLeft func(i, l, r int)
	updateRecLeft = func(i, l, r int) {
		if l == start {
			s.Segments[i] = update(s.Segments[i])
			if l != r {
				s.Segments[i] = mergeSegments(&s.Segments[2*i+1], &s.Segments[2*i+2], s.Segments[i])
			} else {
				s.Segments[i] = mergeSegments(nil, nil, s.Segments[i])
			}
			return
		}
		m := (l + r) / 2
		if start >= m+1 {
			updateRecLeft(2*i+2, m+1, r)
		} else {
			s.Segments[2*i+1] = update(s.Segments[2*i+1])
			updateRecLeft(2*i+1, l, m)
		}
		s.Segments[i] = mergeSegments(&s.Segments[2*i+1], &s.Segments[2*i+2], s.Segments[i])
	}
	var updateRec func(i, l, r int)
	updateRec = func(i, l, r int) {
		if start == l && end == r {
			s.Segments[i] = update(s.Segments[i])
			if l != r {
				s.Segments[i] = mergeSegments(&s.Segments[2*i+1], &s.Segments[2*i+2], s.Segments[i])
			} else {
				s.Segments[i] = mergeSegments(nil, nil, s.Segments[i])
			}
			return
		}
		m := (l + r) / 2
		if end <= m {
			updateRec(2*i+1, l, m)
		} else if start >= m+1 {
			updateRec(2*i+2, m+1, r)
		} else {
			updateRecLeft(2*i+1, l, m)
			updateRecRight(2*i+2, m+1, r)
		}
		s.Segments[i] = mergeSegments(&s.Segments[2*i+1], &s.Segments[2*i+2], s.Segments[i])
	}
	updateRec(0, 0, s.NumberElements-1)
}

func (s *SegmentTree[E, Q]) RightMost(left int, merge func(Q, Q) Q, predicate func(Q) bool) int {
	var recRight func(i, l, r int, q Q) (int, Q)
	recRight = func(i, l, r int, q Q) (int, Q) {
		if l == r {
			tmp := merge(q, s.Segments[i])
			if predicate(tmp) {
				return l + 1, tmp
			} else {
				return l, q
			}
		}
		m := (l + r) / 2
		tmp := merge(q, s.Segments[2*i+1])
		if predicate(tmp) {
			return recRight(2*i+2, m+1, r, tmp)
		} else {
			return recRight(2*i+1, l, m, q)
		}
	}
	var rec func(i, l, r int) (int, Q)
	rec = func(i, l, r int) (int, Q) {
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
	rightMost, _ := rec(0, 0, len(s.Elements)-1)
	return rightMost
}

func (s *SegmentTree[E, Q]) LeftMost(right int, merge func(Q, Q) Q, predicate func(Q) bool) int {
	var recLeft func(i, l, r int, q Q) (int, Q)
	recLeft = func(i, l, r int, q Q) (int, Q) {
		if l == r {
			tmp := merge(q, s.Segments[i])
			if predicate(tmp) {
				return l - 1, tmp
			} else {
				return l, q
			}
		}
		m := (l + r) / 2
		tmp := merge(q, s.Segments[2*i+2])
		if predicate(tmp) {
			return recLeft(2*i+1, l, m, tmp)
		} else {
			return recLeft(2*i+2, m+1, r, q)
		}
	}
	var rec func(i, l, r int) (int, Q)
	rec = func(i, l, r int) (int, Q) {
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
	rightMost, _ := rec(0, 0, len(s.Elements)-1)
	return rightMost
}
