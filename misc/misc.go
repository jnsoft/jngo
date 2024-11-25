package misc

import "math"

type Number interface {
	int | int64 | float64
}

func Sequence[T Number](min, max, step T) []T {
	if step != 0 && (max-min)*step < 0 {
		panic("incorrect stepsize")
	}

	if step == 0 {
		return []T{0}
	}

	size := int(math.Floor(float64((max-min)/step)) + 1)
	seq := make([]T, size)
	var cnt T = 0
	for i := range seq {
		seq[i] = min + cnt*step
		cnt++
	}
	return seq
}

func FilterByArray[A any, B bool](arr []A, filter []B) []A {
	if len(arr) != len(filter) {
		panic("array lengths must be equal")
	}
	filtered := make([]A, 0)
	for i := range filter {
		if filter[i] {
			filtered = append(filtered, arr[i])
		}
	}
	return filtered
}

func Map[T any, M any](arr []T, f func(T) M) []M {
	res := make([]M, len(arr))
	for i, e := range arr {
		res[i] = f(e)
	}
	return res
}

func Filter[A any](arr []A, f func(A) bool) []A {
	filtered := make([]A, 0)
	for _, v := range arr {
		if f(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func Reduce[A, B any](collection []A, accumulator func(B, A) B, initialValue B) B {
	var result = initialValue
	for _, x := range collection {
		result = accumulator(result, x)
	}
	return result
}

// same as Reduce
func Fold[A, B any](arr []A, f func(B, A) B, initval B) B { return Reduce(arr, f, initval) }

func Find[A any](items []A, predicate func(A) bool) (value A, found bool) {
	for _, v := range items {
		if predicate(v) {
			return v, true
		}
	}
	return
}
