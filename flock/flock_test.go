package flock

import (
	"fmt"
	"testing"

	. "github.com/jnsoft/jngo/testhelper"
)

func TestFlock(t *testing.T) {
	t.Run("Test a flock of functions", func(t *testing.T) {

		fac := Y(func(rec func(int) int) func(int) int {
			return Cond(
				func(x int) bool { return x == 0 },
				K(1),
				B(S(func(x, y int) int { return x * y }, rec), C(B, func(x int) int { return x - 1 })),
			)
		})

		n := 5
		fmt.Printf("Factorial of %d is %d\n", n, fac(n))

	})
}

func TestBool(t *testing.T) {
	t.Run("Test boolean functions", func(t *testing.T) {

		val2 := True(true)(false)

		AssertEqual(t, true, val2)

	})
}
