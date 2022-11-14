package queue

import (
	"fmt"
	"strings"
)

type (
	Queue[T any] struct {
		head, end *node[T]
		length    int
	}
	node[T any] struct {
		value T
		next  *node[T]
	}
)

func New[T any]() *Queue[T] { return &Queue[T]{nil, nil, 0} }

func (s *Queue[T]) Len() int { return s.length }

func (s *Queue[T]) Peek() T {
	if s.length == 0 {
		panic("queue underflow")
	}

	return s.head.value
}

func (q *Queue[T]) Dequeue() T {
	if q.length == 0 {
		panic("queue underflow")
	}
	n := q.head
	q.head = q.head.next
	q.length--
	if q.length == 0 {
		q.end = nil
	}
	return n.value
}

func (q *Queue[T]) Enqueue(value T) {
	n := &node[T]{value, nil}
	if q.length == 0 {
		q.head = n
	} else {
		q.end.next = n
	}
	q.end = n
	q.length++
}

func (s *Queue[T]) String() string {
	if s == nil || s.head == nil {
		return "nil"
	}
	cur := s.head
	var sb strings.Builder
	for cur.next != nil {
		sb.WriteString(fmt.Sprintf("%v->", cur.value))
		cur = cur.next
	}
	sb.WriteString(fmt.Sprintf("%v", cur.value))
	return sb.String()
}
