package math

import (
	"math"
)

func FastFourierTransform(p []complex128, inverse bool) []complex128 {
	n := len(p)
	if n == 1 {
		return p
	}
	p0 := make([]complex128, n/2)
	p1 := make([]complex128, n/2)

	for i := 0; i < n/2; i++ {
		p0[i] = p[2*i]
		p1[i] = p[2*i+1]
	}

	y0 := FastFourierTransform(p0, inverse)
	y1 := FastFourierTransform(p1, inverse)

	y := make([]complex128, n)

	w := complex(1, 0)
	var angle float64
	if inverse {
		angle = -(2 * math.Pi) / float64(n)
	} else {
		angle = (2 * math.Pi) / float64(n)
	}
	w1 := complex(math.Cos(angle), math.Sin(angle))
	for i := 0; i < n/2; i++ {
		y[i] = y0[i] + w*y1[i]
		y[i+n/2] = y0[i] - w*y1[i]

		if inverse {
			y[i] /= 2
			y[i+n/2] /= 2
		}
		w *= w1
	}
	return y
}

func MultiplyPolynomials(p1, p2 []int) []int {
	n := 1
	for n < len(p1)+len(p2) {
		n *= 2
	}
	p1c := make([]complex128, n)
	p2c := make([]complex128, n)
	for i := 0; i < len(p1); i++ {
		p1c[i] = complex(float64(p1[i]), 0)
	}
	for i := 0; i < len(p2); i++ {
		p2c[i] = complex(float64(p2[i]), 0)
	}
	y1 := FastFourierTransform(p1c, false)
	y2 := FastFourierTransform(p2c, false)
	y := make([]complex128, n)
	for i := 0; i < n; i++ {
		y[i] = y1[i] * y2[i]
	}
	pc := FastFourierTransform(y, true)
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = int(math.Round(real(pc[i])))
	}
	return p
}
