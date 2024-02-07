package sort_test

import (
	"fmt"
	"math/rand"
	so "sort"
	"testing"

	"github.com/lorenzotinfena/goji/sort"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slices"
)

func gen(n int) []uint {
	v := make([]uint, n)
	for i, e := range rand.Perm(n) {
		v[i] = uint(e)
	}
	return v
}

func TestStandardSort(t *testing.T) {
	v := gen(100)
	so.Slice(v, func(i, j int) bool {
		return v[i] < v[j]
	})
	assert.True(t, slices.IsSorted(v))
}

func TestLSDRadixSort(t *testing.T) {
	v := gen(100)
	sort.LSDRadixSort(v, func(a uint) []uint { return []uint{uint(a)} })
	fmt.Printf("v: %v\n", v)
	assert.True(t, slices.IsSorted(v))
}

func TestMSDRadixSort(t *testing.T) {
	v := gen(100)
	sort.MSDRadixSort(v, func(a uint) []uint { return []uint{uint(a)} })
	assert.True(t, slices.IsSorted(v))
}
