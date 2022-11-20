package inthelper

import (
	"testing"

	. "github.com/jnsoft/jngo/testhelper"
)

func TestSum(t *testing.T) {
	t.Run("sum of ints", func(t *testing.T) {
		ns := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		tot := Sum(ns)
		AssertEqual(t, tot, 45)
	})
}
