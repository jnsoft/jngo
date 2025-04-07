package combinators

import (
	"testing"

	. "github.com/jnsoft/jngo/testhelper"
)

func TestFlock(t *testing.T) {

	t.Run("Test Identity", func(t *testing.T) {

		output := I("42")
		output2 := I(true)
		output3 := I([]int{1, 2, 3})

		AssertEqual(t, output, "42")
		AssertEqual(t, output2, true)
		CollectionAssertEqual(t, output3, []int{1, 2, 3})

		increment := func(x int) int {
			return x + 1
		}

		actualValue := I(increment)(5)
		AssertEqual(t, actualValue, 6)

	})
	t.Run("Test Mockingbird", func(t *testing.T) {

		double := func(x int) int {
			return x * 2
		}

		doubleTwice := M(double)

		output := doubleTwice(3)

		AssertEqual(t, output, 3*2*2)

	})

	t.Run("Test B combinator", func(t *testing.T) {

		double := func(x int) int {
			return x * 2
		}

		increment := func(x int) int {
			return x + 1
		}

		composedFunction := B(double, increment)

		output := composedFunction(3)

		AssertEqual(t, output, (3+1)*2)

	})

	t.Run("Test Y combinator", func(t *testing.T) {

		factorial := func(recurse func(int) int) func(int) int {
			return func(x int) int {
				if x == 0 {
					return 1
				}
				return x * recurse(x-1)
			}
		}

		factorialFunction := Y(factorial)

		output := factorialFunction(5)

		AssertEqual(t, output, 1*2*3*4*5)

	})

	t.Run("Test S combinator", func(t *testing.T) {

		add := func(a, b int) int {
			return a + b
		}

		double := func(x int) int {
			return x * 2
		}

		composedFunction := S(add, double)

		output := composedFunction(5)

		AssertEqual(t, output, 5+5*2)

	})
}

func TestBool(t *testing.T) {
	t.Run("Test boolean functions", func(t *testing.T) {

		val2 := True(true)(false)

		AssertEqual(t, true, val2)

	})
}

/*

fac := Y(func(rec func(int) int) func(int) int {
			return Cond(
				func(x int) bool { return x == 0 }, // Base case: x == 0
				K(1), // Constant value 1 when x == 0
				B(
					S(func(x, y int) int { return x * y },
						C(rec, func(x int) int { return x - 1 })), // Use `C` to flip the arguments for `rec(x - 1)`
					func(x int) int { return x }, // Pass `x` as is
				),
			)
		})

		n := 5
		fmt.Printf("Factorial of %d is %d\n", n, fac(n))

*/
