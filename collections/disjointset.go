package collections

type DisjointSet struct {
	parentsOrSizes []int
}

func NewDisjointSet(n int) *DisjointSet {
	parentsOrSizes := make([]int, n)
	for i := 0; i < n; i++ {
		parentsOrSizes[i] = -1
	}
	return &DisjointSet{parentsOrSizes}
}

func (ds DisjointSet) Root(item int) int {
	for ds.parentsOrSizes[item] >= 0 {
		item = ds.parentsOrSizes[item]
	}
	return item
}

func (ds DisjointSet) Size(root int) int {
	return -ds.parentsOrSizes[root]
}

func (ds *DisjointSet) Merge(root1, root2 int) {
	if ds.parentsOrSizes[root1] > ds.parentsOrSizes[root2] {
		root1, root2 = root2, root1
	}
	ds.parentsOrSizes[root1] += ds.parentsOrSizes[root2]
	ds.parentsOrSizes[root2] = root1
}
