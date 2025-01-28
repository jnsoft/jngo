package pqueue

import (
	"testing"

	. "github.com/jnsoft/jngo/testhelper"
)

func TestQueue(t *testing.T) {
	t.Run("integer min priority queue", func(t *testing.T) {
		//q := new(PriorityQueue[int])

		q := NewPriorityQueue[int](func(i, j int) bool { return i < j })

		// check stack is empty
		AssertTrue(t, q.IsEmpty())
		AssertEqual(t, q.Size(), 0)

		// enqueue item, then check it's not empty
		q.Enqueue(3)
		q.Enqueue(1)
		q.Enqueue(2)

		AssertFalse(t, q.IsEmpty())
		AssertEqual(t, q.Size(), 3)

		// peek first item
		value, err := q.Peek()
		AssertNil(t, err)
		AssertEqual(t, value, 1)

		// dequeue first item
		value, err = q.Dequeue()
		AssertNil(t, err)
		AssertEqual(t, value, 1)

		// dequeue second item
		value, err = q.Dequeue()
		AssertNil(t, err)
		AssertEqual(t, value, 2)

		// dequeue third item, check queue is empty
		value, err = q.Dequeue()
		AssertNil(t, err)
		AssertEqual(t, value, 3)
		AssertTrue(t, q.IsEmpty())

		// can get the numbers we put in as numbers, not untyped interface{}
		q.Enqueue(1)
		q.Enqueue(2)
		fst, _ := q.Dequeue()
		scd, _ := q.Dequeue()
		AssertEqual(t, fst+scd, 3)

		// string representation
		q.Enqueue(1)
		q.Enqueue(2)
		q.Enqueue(3)
		string_rep := q.PrettyPrint()
		AssertEqual(t, string_rep, "1->2->3")
	})
}
