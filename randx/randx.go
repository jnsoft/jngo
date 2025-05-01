package randx

import (
	"math/rand"
	"sync"
	"time"
)

var (
	globalRand *rand.Rand
	once       sync.Once
	mu         sync.Mutex
)

func Get() *rand.Rand {
	once.Do(func() {
		globalRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	})
	return globalRand
}

func Seed(seed int64) {
	mu.Lock()
	defer mu.Unlock()

	once.Do(func() {
		globalRand = rand.New(rand.NewSource(seed))
	})

	globalRand.Seed(seed)
}

func Bool() bool {
	return Intn(2) == 0
}

// Wrapper for Intn
func Int(min, max int) int {
	return Intn(max-min+1) + min
}

// Wrapper for Float64
func Random() float64 {
	return Float64()
}

// Random int in [0, n).
func Intn(n int) int {
	mu.Lock()
	defer mu.Unlock()
	return Get().Intn(n)
}

// Random float64 in [0.0, 1.0).
func Float64() float64 {
	mu.Lock()
	defer mu.Unlock()
	return Get().Float64()
}

// normally distributed float64 with mean 0 and std dev 1.
func NormFloat64() float64 {
	mu.Lock()
	defer mu.Unlock()
	return Get().NormFloat64()
}

// non-negative pseudo-random 63-bit integer as an int64.
func Int63() int64 {
	mu.Lock()
	defer mu.Unlock()
	return Get().Int63()
}

// Perm returns a random permutation of the integers [0, n).
func Perm(n int) []int {
	mu.Lock()
	defer mu.Unlock()
	return Get().Perm(n)
}

// Shuffles a slice using Fisher-Yates.
func Shuffle(sliceLen int, swap func(i, j int)) {
	mu.Lock()
	defer mu.Unlock()
	Get().Shuffle(sliceLen, swap)
}

// Wrapper for Shuffle
func Shuffle2[T any](arr []T) {
	if len(arr) == 0 {
		return
	}
	Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
}

func GetRandomName(length int) string {
	if length < 1 {
		return ""
	}
	const upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const lower = "abcdefghijklmnopqrstuvwxyz"

	var res = make([]byte, length)
	res[0] = upper[Intn(len(upper))]
	for i := 1; i < length; i++ {
		res[i] = lower[Intn(len(lower))]
	}

	return string(res)
}
