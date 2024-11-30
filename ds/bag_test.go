package ds

import (
	"testing"

	. "github.com/jnsoft/jngo/testhelper"
)

func TestBag(t *testing.T) {
	t.Run("test strings", func(t *testing.T) {
		bag := NewBag[string]()
		bag.Add("apple")
		bag.Add("apple")
		bag.Add("orange")
		apples := bag.Count("apple")
		AssertEqual(t, apples, 2)
		cnt := bag.Size()
		AssertEqual(t, cnt, 3)
		bag.Remove("apple")
		apples = bag.Count("apple")
		AssertEqual(t, apples, 1)
	})

	t.Run("test integers", func(t *testing.T) {
		bag := NewBag[int]()
		bag.Add(1)
		bag.Add(1)
		bag.Add(2)
		ones := bag.Count(1)
		AssertEqual(t, ones, 2)
		cnt := bag.Size()
		AssertEqual(t, cnt, 3)
		bag.Remove(1)
		ones = bag.Count(1)
		AssertEqual(t, ones, 1)
	})

	t.Run("test random", func(t *testing.T) {
		bag := NewBag[int]()
		bag.Add(1)
		bag.Add(1)
		bag.Add(2)
		_, ok := bag.PeekRandom()
		AssertTrue(t, ok)
		cnt := bag.Size()
		AssertEqual(t, cnt, 3)
		bag.ExtractRandom()
		bag.ExtractRandom()
		bag.ExtractRandom()
		cnt = bag.Size()
		AssertEqual(t, cnt, 0)
		_, ok = bag.ExtractRandom()
		AssertFalse(t, ok)
	})
}
