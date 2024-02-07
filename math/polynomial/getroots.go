package polynomial

import (
	"math"
)

// Return real and complex root of real polynomial with given coefficients in reverse order with their degree, ending with coefficient with degree 0.
// Every n degree polynomial has n roots
// Is assumed that length of coefficients is at least 1
//
// For polynomial of degree:
//
// - 1: Opposite of coefficient
// - 2: Quadratic formula
func ComputeRoots(coefficients ...float64) (realRoots []float64, complexRoots []complex128) {
	realRoots = make([]float64, 0, len(coefficients))
	complexRoots = make([]complex128, 0, len(coefficients))
	degree := len(coefficients) - 1
	switch degree {
	case 1:
		realRoots = append(realRoots, -coefficients[0])
	case 2:
		discriminant := math.Pow(coefficients[1], 2) - 4*coefficients[0]*coefficients[2]
		firstPart := -coefficients[1] / (2 * coefficients[0])

		if discriminant != 0 {
			if discriminant > 0 {
				secondPart := math.Sqrt(discriminant) / (2 * coefficients[0])
				realRoots = append(realRoots, firstPart+secondPart)
				realRoots = append(realRoots, firstPart-secondPart)
			} else {
				secondPart := math.Sqrt(-discriminant) / (2 * coefficients[0])
				complexRoots = append(complexRoots, complex(firstPart, secondPart))
				complexRoots = append(complexRoots, complex(firstPart, -secondPart))
			}
		} else {
			realRoots = append(realRoots, firstPart)
			realRoots = append(realRoots, firstPart)
		}
	}
	return
}
