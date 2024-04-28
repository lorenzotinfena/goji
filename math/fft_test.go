package math_test

import (
	"testing"

	"github.com/lorenzotinfena/goji/math"
	"github.com/stretchr/testify/assert"
)

func Test1MultiplyPolynomials(t *testing.T) {
	p1 := []int{1, 2}
	p2 := []int{1, 2}
	n := 1
	for n < len(p1)+len(p2) {
		n *= 2
	}
	res := make([]int, n)

	for i := 0; i < len(p1); i++ {
		for j := 0; j < len(p2); j++ {
			res[i+j] += p1[i] * p2[j]
		}
	}
	assert.Equal(t, math.MultiplyPolynomials(p1, p2), res)
}

func Test2MultiplyPolynomials(t *testing.T) {
	p1 := []int{1, 2, -3, 2, 5, -2}
	p2 := []int{8, 2, 2, 5, -2, 1, 3, -3, 7}
	n := 1
	for n < len(p1)+len(p2) {
		n *= 2
	}
	res := make([]int, n)

	for i := 0; i < len(p1); i++ {
		for j := 0; j < len(p2); j++ {
			res[i+j] += p1[i] * p2[j]
		}
	}
	assert.Equal(t, math.MultiplyPolynomials(p1, p2), res)
}
