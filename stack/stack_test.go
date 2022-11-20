package stack

import (
	"testing"

	. "github.com/jnsoft/jngo/testhelper"
)

func TestStack(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		s := new(Stack[int])

		// check stack is empty
		AssertTrue(t, s.IsEmpty())

		// add a thing, then check it's not empty
		s.Push(123)
		AssertFalse(t, s.IsEmpty())

		// add another thing, pop it back again
		s.Push(456)
		value := s.Pop()
		AssertEqual(t, value, 456)
		value = s.Pop()
		AssertEqual(t, value, 123)
		AssertTrue(t, s.IsEmpty())

		// can get the numbers we put in as numbers, not untyped interface{}
		s.Push(1)
		s.Push(2)
		fst := s.Pop()
		scd := s.Pop()
		AssertEqual(t, fst+scd, 3)
	})
}
