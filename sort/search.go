package sort

func LowerBound[T any](v []T, element T, prior func(a, b T) bool) int {
	l := 0
	r := len(v)
	for l != r {
		middle := (l + r) / 2
		if prior(v[middle], element) {
			l = middle + 1
		} else {
			r = middle
		}
	}
	return l
}

func UpperBound[T any](v []T, element T, prior func(a, b T) bool) int {
	l := 0
	r := len(v)
	for l != r {
		middle := (l + r) / 2
		if prior(element, v[middle]) {
			r = middle
		} else {
			l = middle + 1
		}
	}
	return l
}
