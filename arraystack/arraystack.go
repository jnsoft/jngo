package arraystack

type ArrayStack[T any] struct {
	data []T
}

func (s *ArrayStack[T]) Push(value T) {
	s.data = append(s.data, value)
}

func (s *ArrayStack[T]) Pop() (T, bool) {
	if len(s.data) == 0 {
		var zero T
		return zero, false
	}

	value := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return value, true
}

func (s *ArrayStack[T]) Peek() (T, bool) {
	if len(s.data) == 0 {
		var zero T
		return zero, false
	}

	return s.data[len(s.data)-1], true
}

func (s *ArrayStack[T]) IsEmpty() bool {
	return len(s.data) == 0
}
