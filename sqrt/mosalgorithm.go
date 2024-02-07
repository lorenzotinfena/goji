package sqrt

import (
	"math"

	"github.com/lorenzotinfena/goji/collections"
	ll "github.com/lorenzotinfena/goji/collections/linkedlist"
	"github.com/lorenzotinfena/goji/sort"
	"github.com/lorenzotinfena/goji/utils"
)

// Use this if in SqrtDecompsitionSimple, the mergeQ function is hard, but an expandQ function is easy
func MoSAlgorithm[E any, Q any](
	querySingleElement func(element E) Q,
	expandQ func(*Q, E),
	clone func(Q) Q,
	elements []E,
	queries []struct {
		left  uint64
		right uint64
	},
) (res []Q) {
	// Initialization
	res = make([]Q, len(queries))
	sqrt := math.Sqrt(float64(len(elements)))
	blockSize := uint64(sqrt)
	blocks := make([]*ll.SinglyLinkedList[struct {
		left  uint64
		right uint64
		index uint64
	}], uint64(math.Ceil(float64(len(elements))/float64(blockSize))))
	for i := range blocks {
		blocks[i] = ll.NewSinglyLinkedList[struct {
			left  uint64
			right uint64
			index uint64
		}](utils.Equalize[struct {
			left  uint64
			right uint64
			index uint64
		}]())
	}
	for i, v := range queries {
		blocks[v.left/blockSize].InsertLast(struct {
			left  uint64
			right uint64
			index uint64
		}{
			left:  v.left,
			right: v.right,
			index: uint64(i),
		})
	}
	blockSorted := make([]*collections.Queue[struct {
		left  uint64
		right uint64
		index uint64
	}], len(blocks))
	for i := range blocks {
		tmp := blocks[i].ToSlice()
		sort.SelectionSort(tmp, func(a, b struct {
			left  uint64
			right uint64
			index uint64
		},
		) bool {
			return a.right < b.right
		})
		blockSorted[i] = collections.NewQueue[struct {
			left  uint64
			right uint64
			index uint64
		}]()
		for _, tmp2 := range tmp {
			blockSorted[i].Enqueue(tmp2)
		}
	}
	// Main
	for i := uint64(0); i < uint64(len(blockSorted)); i++ {
		block := blockSorted[i]
		if block.Len() == 0 {
			continue
		}
		for block.Len() > 0 {
			if block.Preview().right/blockSize == i {
				q := block.Dequeue()
				res[q.index] = querySingleElement(elements[q.left])
				q.left++
				for q.left < q.right {
					expandQ(&res[q.index], elements[q.left])
				}
			} else {
				break
			}
		}
		middleIndex := (i+1)*blockSize - 1
		rightRes := querySingleElement(elements[middleIndex])
		rightIndex := middleIndex + 1 // now is the first index of the next block
		for block.Len() > 0 {
			q := block.Dequeue()
			for rightIndex < q.right { // Expand right side
				expandQ(&rightRes, elements[rightIndex])
				rightIndex++
			}
			res[q.index] = clone(rightRes) // Expand left side
			for q.left < middleIndex {
				expandQ(&res[q.index], elements[q.left])
				q.left++
			}
		}
	}
	return
}
