package misc

import (
	"strings"
	"testing"

	. "github.com/jnsoft/jngo/testhelper"
)

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
