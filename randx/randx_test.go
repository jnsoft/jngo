package randx

import (
	"sync"
	"testing"

	. "github.com/jnsoft/jngo/testhelper"
)

func TestRandx(t *testing.T) {

	t.Run("Test Seed Determinism", func(t *testing.T) {

		Seed(42)

		val1 := Intn(100)
		val2 := Float64()
		val3 := NormFloat64()
		perm := Perm(5)

		Seed(42)

		AssertEqual(t, Intn(100), val1)
		AssertEqual(t, Float64(), val2)
		AssertEqual(t, NormFloat64(), val3)
		CollectionAssertEqual(t, Perm(5), perm)
	})

	t.Run("Test Thread Safety", func(t *testing.T) {

		Seed(100)

		var wg sync.WaitGroup
		errors := make(chan error, 100)

		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				defer func() {
					if r := recover(); r != nil {
						errors <- r.(error)
					}
				}()
				Intn(1000)
				Float64()
				NormFloat64()
				Shuffle(10, func(i, j int) {})
				Perm(10)
			}()
		}

		wg.Wait()
		close(errors)

		for err := range errors {
			t.Errorf("thread safety error: %v", err)
		}
	})

	t.Run("Test Shuffle", func(t *testing.T) {
		arr := make([]int, 1000)
		for i := range arr {
			arr[i] = i
		}
		Seed(42)
		Shuffle2(arr)

		unchanged := true
		for i, v := range arr {
			if v != i {
				unchanged = false
				break
			}
		}

		AssertFalse(t, unchanged)

	})

	t.Run("Test Random name", func(t *testing.T) {
		name := GetRandomName(10)
		AssertEqual(t, len(name), 10)
	})

}
