package slices

// Using code from: golang.org/x/exp/slices

func Insert[T []E, E any](sli T, index int, element E) T {
	n := len(sli)
	if index == n {
		return append(sli, element)
	}
	if n == cap(sli) {
		s2 := make(T, n+1, n*2)
		copy(s2[:index], sli[:index])
		copy(s2[index+1:], sli[index:])
		s2[index] = element
		return s2
	}
	sli = sli[:n+1]
	copy(sli[index+1:], sli[index:])
	sli[index] = element
	return sli
}

func Clone[T any](v []T) []T {
	v1 := make([]T, len(v))
	copy(v1, v)
	return v1
}

func Map[Dom any, Codom any](v []Dom, f func(Dom) Codom) []Codom {
	mapped := make([]Codom, len(v))
	for i, x := range v {
		mapped[i] = f(x)
	}
	return mapped
}
