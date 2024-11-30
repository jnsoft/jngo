package red_black_dst

import (
	"fmt"
	"testing"

	. "github.com/jnsoft/jngo/testhelper"
)

func TestRedBlackBST(t *testing.T) {
	t.Run("test new tree", func(t *testing.T) {
		tree := NewRedBlackTree[int, string]()
		AssertTrue(t, tree.IsEmpty())
	})

	t.Run("test int, string", func(t *testing.T) {
		n := 1000
		tree := NewRedBlackTree[int, string]()
		for i := 0; i < n; i++ {
			tree.Put(i, fmt.Sprint(i))
		}

		min, err := tree.Min()
		AssertTrue(t, err == nil)
		AssertEqual(t, min, 0)

		max, err := tree.Max()
		AssertTrue(t, err == nil)
		AssertEqual(t, max, n-1)
	})

	t.Run("test string, int", func(t *testing.T) {
		tree := NewRedBlackTree[string, int]()

		tree.Put("a", 0)
		tree.Put("aa", 1)
		tree.Put("aaa", 2)
		tree.Put("b", 3)
		tree.Put("c", 4)
		tree.Put("abc", 5)
		tree.Put("bcc", 6)

		min, err := tree.Min()
		AssertTrue(t, err == nil)
		AssertEqual(t, min, "a")

		max, err := tree.Max()
		AssertTrue(t, err == nil)
		println(max)
		AssertEqual(t, max, "c")
	})
}
