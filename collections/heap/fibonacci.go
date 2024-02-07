package heap

import (
	"math/bits"

	"github.com/lorenzotinfena/goji/collections"
)

type FibonacciHeapNode[T comparable] struct {
	value                      T
	degree                     int
	marked                     bool
	parent, right, left, child *FibonacciHeapNode[T]
}

type FibonacciHeap[T comparable] struct {
	prior    func(T, T) bool
	size     int
	nodesMap collections.MultiMap[T, *FibonacciHeapNode[T]]
	rootlist *FibonacciHeapNode[T] // The first node is always the first to be popped
}

func NewFibonacciHeap[T comparable](prior func(T, T) bool, nodesMap collections.MultiMap[T, *FibonacciHeapNode[T]]) *FibonacciHeap[T] {
	return &FibonacciHeap[T]{
		prior:    prior,
		nodesMap: nodesMap,
		rootlist: nil,
	}
}

func (fb *FibonacciHeap[T]) Len() int {
	return fb.size
}

func (fb *FibonacciHeap[T]) Push(value T) {
	fb.size++
	node := &FibonacciHeapNode[T]{value: value}

	fb.nodesMap.Add(value, node)

	node.left = node
	node.right = node

	if fb.rootlist == nil {
		fb.rootlist = node
	} else {
		fb.appendSiblings(node)
	}

	if fb.prior(value, fb.rootlist.value) {
		fb.rootlist = node
	}
}

func (fb *FibonacciHeap[T]) Preview() T {
	return fb.rootlist.value
}

// Assumptions:
// - *base != nil
func (fb *FibonacciHeap[T]) appendSibling(base **FibonacciHeapNode[T], toappend *FibonacciHeapNode[T]) {
	toappend.parent = (*base).parent
	toappend.marked = false
	toappend.right = *base
	toappend.left = (*base).left
	(*base).left.right = toappend
	(*base).left = toappend
}

// Assumptions:
// - *base != nil
func (fb *FibonacciHeap[T]) appendSiblings(toappend *FibonacciHeapNode[T]) {
	if toappend == nil {
		return
	}

	tmp := toappend
	for {
		tmp.parent = fb.rootlist.parent
		tmp.marked = false
		tmp = tmp.right
		if tmp == toappend {
			break
		}
	}

	toappend.left.right = fb.rootlist
	fb.rootlist.left.right = toappend
	fb.rootlist.left, toappend.left = toappend.left, fb.rootlist.left
}

func (fb *FibonacciHeap[T]) Contains(value T) bool {
	return fb.nodesMap.Contains(value)
}

func (fb *FibonacciHeap[T]) mark(node *FibonacciHeapNode[T]) {
	for node.marked {
		parent := node.parent
		parent.degree--
		if parent.child == node {
			if parent.degree == 0 {
				parent.child = nil
			} else {
				parent.child = node.right
			}
		}
		node.left.right = node.right
		node.right.left = node.left
		node.left = node
		node.right = node
		fb.appendSibling(&fb.rootlist, node)
		node = parent
	}
	if node.parent != nil {
		node.marked = true
	}
}

func (fb *FibonacciHeap[T]) Remove(value T) {
	node := fb.nodesMap.GetOne(value)
	// Handle case where value is the only minimum
	if node == fb.rootlist {
		fb.Pop()
		return
	}

	fb.nodesMap.Remove(value, node)
	if node.parent == nil {
		node.left.right = node.right
		node.right.left = node.left
	} else {
		node.parent.degree--
		if node.parent.child == node {
			if node.parent.degree == 0 {
				node.parent.child = nil
			} else {
				node.parent.child = node.right
			}
		}
		node.left.right = node.right
		node.right.left = node.left
		fb.mark(node.parent)
	}

	fb.appendSiblings(node.child)

	fb.size--
}

// Assumptions:
// - prior(old, new) = false
func (fb *FibonacciHeap[T]) DecreaseKey(old, new T) {
	node := fb.nodesMap.GetOne(old)
	fb.nodesMap.Remove(old, node)
	fb.nodesMap.Add(new, node)

	node.value = new
	if node.parent != nil {
		node.parent.degree--
		if node.parent.child == node {
			if node.parent.degree == 0 {
				node.parent.child = nil
			} else {
				node.parent.child = node.right
			}
		}
		node.left.right = node.right
		node.right.left = node.left
		fb.mark(node.parent)
		fb.appendSibling(&fb.rootlist, node)
	}
	if fb.prior(new, fb.rootlist.value) {
		fb.rootlist = node
	}
}

func (fb *FibonacciHeap[T]) Pop() T {
	toreturn := fb.rootlist.value
	fb.nodesMap.Remove(toreturn, fb.rootlist)

	fb.size--
	if fb.size == 0 {
		fb.rootlist = nil
		return toreturn
	}

	supp := fb.rootlist.child
	if supp != nil {
		for {
			supp = supp.right
			if supp == fb.rootlist.child {
				break
			}
		}
	}

	fb.appendSiblings(fb.rootlist.child)

	fb.rootlist.left.right = fb.rootlist.right
	fb.rootlist.right.left = fb.rootlist.left

	// Consolidate

	roots := []*FibonacciHeapNode[T]{}
	supp = fb.rootlist.right
	for {
		roots = append(roots, supp)
		supp = supp.right
		if supp == fb.rootlist.right {
			break
		}
	}
	for _, r := range roots {
		r.right = r
		r.left = r
	}
	degrees := make([]*FibonacciHeapNode[T], 2+(64-bits.LeadingZeros(uint(fb.size))-1))
	for _, supp := range roots {
		for degrees[supp.degree] != nil {
			// Here degrees[supp.degree] should keep the lowest value and supp the highest value
			if fb.prior(supp.value, degrees[supp.degree].value) {
				supp, degrees[supp.degree] = degrees[supp.degree], supp
			}

			// Merge trees
			if degrees[supp.degree].child == nil {
				supp.parent = degrees[supp.degree]
				degrees[supp.degree].child = supp
			} else {
				fb.appendSibling(&degrees[supp.degree].child, supp)
			}

			supp = degrees[supp.degree]
			degrees[supp.degree] = nil
			supp.degree++
		}
		degrees[supp.degree] = supp
	}

	// Set new first value to be popped
	degreesTmp := []*FibonacciHeapNode[T]{}
	for _, node := range degrees {
		if node != nil {
			degreesTmp = append(degreesTmp, node)
		}
	}
	degrees = degreesTmp
	degrees[0].left = degrees[len(degrees)-1]
	degrees[len(degrees)-1].right = degrees[0]
	for i := 1; i < len(degrees); i++ {
		degrees[i-1].right = degrees[i]
		degrees[i].left = degrees[i-1]
	}
	fb.rootlist = degrees[0]
	for i := 1; i < len(degrees); i++ {
		if fb.prior(degrees[i].value, fb.rootlist.value) {
			fb.rootlist = degrees[i]
		}
	}
	return toreturn
}
