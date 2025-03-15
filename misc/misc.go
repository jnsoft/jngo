package misc

import (
	"math"
	"math/rand"
	"strings"
	"reflect"
	"bytes"
    	"fmt"
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

func Contains[T any](arr []T, target T) bool {
    for _, value := range arr {
        if reflect.DeepEqual(value, target) {
            return true
        }
    }
    return false
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

func HasDuplicates[T comparable](arr []T) bool {
	seen := make(map[T]struct{})
	for _, v := range arr {
		if _, exists := seen[v]; exists {
			return true // dup found
		}
		seen[v] = struct{}{}
	}
	return false
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

func RangeSlice(a, b int) []int {
    size := b - a + 1
    slice := make([]int, size)
    for i := range slice {
        slice[i] = a + i
    }
    return slice
}

func SplitStrings(input []string, delimiter string) [][]string {
	var result [][]string
	for _, str := range input {
		splitStr := strings.Split(str, delimiter)
		result = append(result, splitStr)
	}
	return result
}

func IsPalindrome(s string) bool {
	normalized := strings.ToLower(strings.ReplaceAll(s, " ", ""))
	for i := 0; i < len(normalized)/2; i++ {
		if normalized[i] != normalized[len(normalized)-1-i] {
			return false
		}
	}
	return true
}

func Permutations(arr []any) [][]any {
    var result [][]any
    generatePermutations(arr, 0, &result)
    return result
}

func generatePermutations(arr []any, start int, result *[][]any) {
    if start == len(arr)-1 {
        // Append a copy of the current permutation to the result
        temp := make([]any, len(arr))
        copy(temp, arr)
        *result = append(*result, temp)
        return
    }

    for i := start; i < len(arr); i++ {
        // Swap current element with the starting element
        arr[start], arr[i] = arr[i], arr[start]

        // Recursively generate permutations for the remaining elements
        generatePermutations(arr, start+1, result)

        // Swap back to restore the original state
        arr[start], arr[i] = arr[i], arr[start]
    }


// The Hamming Distance measures the minimum number of substitutions required to change one string into the other
func HammingDistance_string(s1, s2 string) (int) {
	if len(s1) != len(s2) {
		return -1
	}

	distance := 0
	for i := range s1 {
		if s1[i] != s2[i] {
			distance++
		}
	}
	return distance
}

// Edit distance/Hamming distance between two strings.The Hamming distance is here defined as the number of differing bits
func HammingDistance(var1, var2 []byte) int {
    arr1 := bytesToBits(var1)
    arr2 := bytesToBits(var2)

    diff := 0

    if len(arr1) >= len(arr2) {
        for i := 0; i < len(arr2); i++ {
            if arr1[i] != arr2[i] {
                diff++
            }
        }
        // Account for extra bits in the longer slice
        diff += len(arr1) - len(arr2)
    } else {
        for i := 0; i < len(arr1); i++ {
            if arr1[i] != arr2[i] {
                diff++
            }
        }
        // Account for extra bits in the longer slice
        diff += len(arr2) - len(arr1)
    }

    return diff
}

func bytesToBits(data []byte) []bool {
    bits := []bool{}
    for _, b := range data {
        for i := 7; i >= 0; i-- { 
            bits = append(bits, (b&(1<<i)) != 0)
        }
    }
    return bits
}
