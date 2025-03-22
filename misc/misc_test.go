package misc

import (
	"fmt"
	"strings"
	"testing"

	. "github.com/jnsoft/jngo/testhelper"
)

func TestSequence(t *testing.T) {
	t.Run("generate ints", func(t *testing.T) {
		ns := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		is := Sequence(0, 9, 1)
		CollectionAssertEqual(t, is, ns)
	})

	t.Run("generate floats", func(t *testing.T) {
		fs := Sequence[float64](0.0001, 1, 0.0001)
		AssertEqual(t, len(fs), 10000)
		AssertEqual(t, fs[0], 0.0001)
		AssertEqual(t, fs[9999], 1)
	})

	t.Run("negative sequence", func(t *testing.T) {
		fs := Sequence[float64](1, 0, -0.1)
		AssertEqual(t, len(fs), 10+1)
		AssertEqual(t, fs[0], 1)
		AssertEqual(t, fs[10], 0)
	})

	t.Run("short sequence", func(t *testing.T) {
		fs := Sequence(0, 0, 100)
		AssertEqual(t, len(fs), 1)
		AssertEqual(t, fs[0], 0)
	})

	t.Run("no step sequence", func(t *testing.T) {
		fs := Sequence(0, 0, 0)
		AssertEqual(t, len(fs), 1)
		AssertEqual(t, fs[0], 0)
	})
}

func TestFilterByArray(t *testing.T) {
	t.Run("filter three", func(t *testing.T) {
		ns := []int{1, 2, 3, 4, 5, 6}
		even := FilterByArray(ns, []bool{false, true, false, true, false, false})
		CollectionAssertEqual(t, even, []int{2, 4})
	})
}

func TestMap(t *testing.T) {
	t.Run("square", func(t *testing.T) {
		ns := []int{1, 2, 3, 4}
		sqs := Map(ns, func(x int) int {
			return x * x
		})
		CollectionAssertEqual(t, sqs, []int{1, 4, 9, 16})
	})
}

func TestZip(t *testing.T) {
	t.Run("string zip", func(t *testing.T) {
		arr1 := []int{1, 2, 3}
		arr2 := []string{"a", "b", "c"}
		result := Zip(arr1, arr2, func(a int, b string) string {
			return fmt.Sprintf("%d%s", a, b)
		})
		AssertEqual(t, len(result), 3)
		CollectionAssertEqual(t, result, []string{"1a", "2b", "3c"})
	})
}

func TestFilter(t *testing.T) {
	t.Run("even numbers", func(t *testing.T) {
		ns := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
		even := Filter(ns, func(i int) bool {
			return i%2 == 0
		})
		CollectionAssertEqual(t, even, []int{2, 4, 6, 8, 0})
	})

	t.Run("concatenate strings", func(t *testing.T) {
		concatenate := func(x, y string) string {
			return x + y
		}

		AssertEqual(t, Reduce([]string{"a", "b", "c"}, concatenate, ""), "abc")
	})
}

func TestReduce(t *testing.T) {
	t.Run("multiplication of all elements", func(t *testing.T) {
		multiply := func(x, y int) int {
			return x * y
		}
		AssertEqual(t, Reduce([]int{1, 2, 3}, multiply, 1), 6)
	})

	t.Run("concatenate strings", func(t *testing.T) {
		concatenate := func(x, y string) string {
			return x + y
		}

		AssertEqual(t, Reduce([]string{"a", "b", "c"}, concatenate, ""), "abc")
	})
}

func TestFind(t *testing.T) {
	t.Run("find first even number", func(t *testing.T) {
		numbers := []int{1, 3, 4, 5, 6, 7, 8, 9, 10}

		firstEvenNumber, found := Find(numbers, func(x int) bool {
			return x%2 == 0
		})
		AssertTrue(t, found)
		AssertEqual(t, firstEvenNumber, 4)
	})

	t.Run("Find the best fruit", func(t *testing.T) {
		// arrange
		best_fruit := "Banana"

		fruits := []string{
			"Apple",
			"Banana",
		}

		// act

		best, found := Find(fruits, func(s string) bool {
			return strings.Contains(s, best_fruit)
		})

		AssertTrue(t, found)
		AssertEqual(t, best, best_fruit)
	})
}

func TestSubArray(t *testing.T) {
	t.Run("test sub array", func(t *testing.T) {
		numbers := []int{1, 3, 4, 5, 6, 7, 8, 9, 10}

		sub_arr1 := SubArray(numbers, 8)
		sub_arr2 := SubArray(numbers, 1, 3)

		numbers[8] = 0
		numbers[1] = 0

		CollectionAssertEqual(t, sub_arr1, []int{10})
		CollectionAssertEqual(t, sub_arr2, []int{3, 4, 5})
	})
}

func TestReverse(t *testing.T) {
	t.Run("Test Reverse", func(t *testing.T) {
		start := []int{1, 2, 3, 4}
		expected := []int{4, 3, 2, 1}
		Reverse[int](start, 0, len(start)-1)
		CollectionAssertEqual(t, start, expected)
	})

}

func TestRotate(t *testing.T) {
	t.Run("Test Rotate", func(t *testing.T) {
		start := []int{1, 2, 3, 4, 5, 6}
		start2 := Copy(start)
		expected := []int{5, 6, 1, 2, 3, 4}
		Rotate(start, 2)
		Rotate(start2, 8)
		CollectionAssertEqual(t, start, expected)
		CollectionAssertEqual(t, start2, expected)
	})

	t.Run("Test negative rotate", func(t *testing.T) {
		start := []int{1, 2, 3, 4, 5, 6}
		start2 := Copy(start)
		expected := []int{3, 4, 5, 6, 1, 2}
		Rotate(start, -2)
		Rotate(start2, -8)
		CollectionAssertEqual(t, start, expected)
		CollectionAssertEqual(t, start2, expected)
	})

}
