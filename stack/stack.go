package stack

import (
	"fmt"
	"strings"
)

type (
	Stack[T any] struct {
		head   *node[T]
		length int
	}

	node[T any] struct {
		value T
		prev  *node[T]
	}
)

func New[T any]() *Stack[T] { return &Stack[T]{nil, 0} }

func (s *Stack[T]) Len() int { return s.length }

func (s *Stack[T]) IsEmpty() bool { return s.length == 0 }

func (s *Stack[T]) Peek() T {
	if s.length == 0 {
		panic("stack underflow")
	}

	return s.head.value
}

func (s *Stack[T]) Pop() T {
	if s.length == 0 {
		panic("stack underflow")
	}

	n := s.head
	s.head = n.prev
	s.length--
	return n.value
}

func (s *Stack[T]) Push(value T) {
	n := &node[T]{value, s.head}
	s.head = n
	s.length++
}

func (s *Stack[T]) ToArray() []T {
	result := make([]T, 0, s.length)
	for s.length > 0 {
		value := s.Pop()
		result = append(result, value)
	}
	return result
}

// yield
func (s *Stack[T]) ToChannel() <-chan T {
	ch := make(chan T)

	go func() {
		defer close(ch) // Close the channel when done
		current := s.head
		for current != nil {
			ch <- current.value
			current = current.prev
		}
	}()

	return ch
}

func (s *Stack[T]) String() string {
	if s == nil || s.head == nil {
		return "nil"
	}
	cur := s.head
	var sb strings.Builder
	for cur.prev != nil {
		sb.WriteString(fmt.Sprintf("%v->", cur.value))
		cur = cur.prev
	}
	sb.WriteString(fmt.Sprintf("%v", cur.value))
	return sb.String()
}
