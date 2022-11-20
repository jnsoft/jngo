package queue

import (
	"testing"

	. "github.com/jnsoft/jngo/testhelper"
)

func TestStack(t *testing.T) {
	t.Run("integer queue", func(t *testing.T) {
		q := new(Queue[int])

		// check stack is empty
		AssertTrue(t, q.IsEmpty())

		// enqueue item, then check it's not empty
		q.Enqueue(1)
		AssertFalse(t, q.IsEmpty())

		// enqueue another item, dequeue first item
		q.Enqueue(2)
		value := q.Dequeue()
		AssertEqual(t, value, 1)
		value = q.Dequeue()
		AssertEqual(t, value, 2)
		AssertTrue(t, q.IsEmpty())

		// peek item
		q.Enqueue(1)
		value = q.Peek()
		AssertFalse(t, q.IsEmpty())
		AssertEqual(t, value, 1)
		value = q.Dequeue()
		AssertTrue(t, q.IsEmpty())
		AssertEqual(t, value, 1)

		// can get the numbers we put in as numbers, not untyped interface{}
		q.Enqueue(1)
		q.Enqueue(2)
		fst := q.Dequeue()
		scd := q.Dequeue()
		AssertEqual(t, fst+scd, 3)
	})
}
