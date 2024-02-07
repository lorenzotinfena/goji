package collections

import "fmt"

func MakePair[A, B any](first A, second B) Pair[A, B] {
	return Pair[A, B]{
		First:  first,
		Second: second,
	}
}

type Pair[A, B any] struct {
	First  A
	Second B
}

func (p Pair[A, B]) String() string {
	return "(" + fmt.Sprint(p.First) + ", " + fmt.Sprint(p.Second) + ")"
}

func MakeTriple[A, B, C any](first A, second B, third C) Triple[A, B, C] {
	return Triple[A, B, C]{
		First:  first,
		Second: second,
		Third:  third,
	}
}

type Triple[A, B, C any] struct {
	First  A
	Second B
	Third  C
}

func (t *Triple[A, B, C]) String() string {
	return fmt.Sprint(t.First, t.Second, t.Third)
}
