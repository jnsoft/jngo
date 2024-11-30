package bag

import (
	"math/rand"
)

// Bag or multiset, similar to a set, but it allows duplicate elements

type Bag[T comparable] struct {
	elements  map[T]int
	totalSize int
}

func NewBag[T comparable]() *Bag[T] {
	return &Bag[T]{elements: make(map[T]int)}
}

func (b *Bag[T]) Add(element T) {
	b.elements[element]++
	b.totalSize++
}

func (b *Bag[T]) Remove(element T) {
	if count, found := b.elements[element]; found {
		if count > 1 {
			b.elements[element]--
		} else {
			delete(b.elements, element)
		}
		b.totalSize--
	}
}

func (b *Bag[T]) Count(element T) int {
	return b.elements[element]
}

func (b *Bag[T]) Size() int {
	return b.totalSize
}

func (b *Bag[T]) PeekRandom() (T, bool) {
	if b.totalSize == 0 {
		var zero T
		return zero, false
	}
	i := rand.Intn(b.totalSize)
	for element, count := range b.elements {
		if i < count {
			return element, true
		}
		i -= count
	}
	var zero T
	return zero, false
}

func (b *Bag[T]) ExtractRandom() (T, bool) {
	element, found := b.PeekRandom()
	if found {
		b.Remove(element)
	}
	return element, found
}
