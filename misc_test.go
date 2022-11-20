package misc

import (
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
