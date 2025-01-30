package misc

import (
	"math"
	"math/rand"
)

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64
}

type Ordered interface {
	Number | ~string
}

func Max[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T Ordered](a, b T) T {
	if a <= b {
		return a
	}
	return b
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

func Zip[T1 any, T2 any, R any](arr1 []T1, arr2 []T2, f func(T1, T2) R) []R {
	length := len(arr1)
	if len(arr2) < length {
		length = len(arr2)
	}
	res := make([]R, length)
	for i := 0; i < length; i++ {
		res[i] = f(arr1[i], arr2[i])
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

// / length ...int means it's a variadic parameter, it's optional
func SubArray[T any](data []T, index int, length ...int) []T {
	l := len(data) - index
	if len(length) > 0 && length[0] != -1 {
		l = length[0]
	}
	result := make([]T, l)
	copy(result, data[index:index+l])
	return result
}

func GetRandomValues[T any](arr []T, n int, repetitions bool) []T {
	result := make([]T, 0, n)

	if repetitions {
		for i := 0; i < n; i++ {
			result = append(result, arr[rand.Intn(len(arr))])
		}
	} else {
		if n > len(arr) {
			n = len(arr)
		}
		indices := rand.Perm(len(arr))
		for i := 0; i < n; i++ {
			result = append(result, arr[indices[i]])
		}
	}
	return result
}

func GetElementsByIndexes[T any](arr []T, ixs []int) []T {
	result := make([]T, 0, len(ixs))
	for _, ix := range ixs {
		if ix >= 0 && ix < len(arr) {
			result = append(result, arr[ix])
		}
	}
	return result
}

func GetRandomBytes(n int) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = byte(rand.Intn(256))
	}
	return b
}
