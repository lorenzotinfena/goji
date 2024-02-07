package sort_test

import (
	so "sort"
	"testing"

	"github.com/lorenzotinfena/goji/sort"
)

const n = 4000000

func BenchmarkStandardSort(b *testing.B) {
	v := gen(n)
	so.Slice(v, func(i, j int) bool {
		return v[i] < v[j]
	})
}

func BenchmarkLSDRadixSort(b *testing.B) {
	v := gen(n)
	sort.LSDRadixSort(v, func(a uint) []uint { return []uint{uint(a)} })
}

func BenchmarkMSDRadixSort(b *testing.B) {
	v := gen(n)
	sort.MSDRadixSort(v, func(a uint) []uint { return []uint{uint(a)} })
}
