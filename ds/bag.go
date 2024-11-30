package ds

// Bag or multiset, similar to a set, but it allows duplicate elements

type Bag[T comparable] struct {
	elements map[T]int
}

func NewBag[T comparable]() *Bag[T] {
	return &Bag[T]{elements: make(map[T]int)}
}

func (b *Bag[T]) Add(element T) {
	b.elements[element]++
}

func (b *Bag[T]) Remove(element T) {
	if count, found := b.elements[element]; found {
		if count > 1 {
			b.elements[element]--
		} else {
			delete(b.elements, element)
		}
	}
}

func (b *Bag[T]) Count(element T) int {
	return b.elements[element]
}

func (b *Bag[T]) Size() int {
	size := 0
	for _, count := range b.elements {
		size += count
	}
	return size
}
