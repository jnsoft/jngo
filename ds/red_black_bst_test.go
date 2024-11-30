package ds

import (
	"testing"

	. "github.com/jnsoft/jngo/testhelper"
)

func TestRedBlackBST(t *testing.T) {
	t.Run("test int, string", func(t *testing.T) {
		tree := &RedBlackTree[int, string]{}

		AssertEqual(t, tree, nil)
	})
}
