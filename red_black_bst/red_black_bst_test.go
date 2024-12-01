package red_black_dst

import (
	"fmt"
	"testing"

	. "github.com/jnsoft/jngo/testhelper"
)

func TestRedBlackBST(t *testing.T) {

	n := 1000
	t.Run("test new tree", func(t *testing.T) {
		tree := NewRedBlackTree[int, string]()
		AssertTrue(t, tree.IsEmpty())
	})

	t.Run("test int, string", func(t *testing.T) {
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

	t.Run("validate tree", func(t *testing.T) {
		tree := NewRedBlackTree[int, string]()
		for i := 0; i < n; i++ {
			tree.Put(i, fmt.Sprint(i))
		}

		isValid, err := ValidateRedBlackTree(tree)
		AssertTrue(t, isValid)
		AssertTrue(t, err == nil)
	})

	t.Run("get keys", func(t *testing.T) {
		tree := NewRedBlackTree[int, string]()
		for i := 0; i < n; i++ {
			tree.Put(i, fmt.Sprint(i))
		}

		keys := tree.Keys()
		AssertEqual(t, len(keys), n)
	})

	t.Run("copy tree", func(t *testing.T) {
		tree := NewRedBlackTree[int, string]()
		for i := 0; i < n; i++ {
			tree.Put(i, fmt.Sprint(i))
		}

		tree2 := tree.Copy()
		AssertEqual(t, tree2.Size(), tree.Size())

		min1, err := tree.Min()
		AssertTrue(t, err == nil)
		min2, err := tree2.Min()
		AssertEqual(t, min1, min2)
		AssertTrue(t, err == nil)
	})

	t.Run("cieling", func(t *testing.T) {
		tree := NewRedBlackTree[int, string]()
		for i := 0; i < n; i++ {
			tree.Put(i, fmt.Sprint(i))
		}

		tree.Delete(400)

		f, found := tree.Ceiling(400)
		AssertTrue(t, found)
		AssertEqual(t, f, 401)
	})

	t.Run("floor", func(t *testing.T) {
		tree := NewRedBlackTree[int, string]()
		for i := 0; i < n; i++ {
			tree.Put(i, fmt.Sprint(i))
		}

		tree.Delete(400)

		f, found := tree.Floor(400)
		AssertTrue(t, found)
		AssertEqual(t, f, 399)
	})

	t.Run("select", func(t *testing.T) {
		tree := NewRedBlackTree[int, string]()
		for i := 0; i < n; i++ {
			tree.Put(i, fmt.Sprint(i))
		}

		tree.Delete(400)

		f, found := tree.Select(400)
		AssertTrue(t, found)
		AssertEqual(t, f, 401)
	})

	t.Run("rank", func(t *testing.T) {
		tree := NewRedBlackTree[int, string]()
		for i := 0; i < n; i++ {
			tree.Put(i, fmt.Sprint(i))
		}

		tree.Delete(400)
		tree.Delete(200)

		rank := tree.Rank(400)
		AssertEqual(t, rank, 399)
	})
}
