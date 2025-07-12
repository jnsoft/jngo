package pqueue

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"
)

const (
	EMPTY_QUEUE = "queue is empty"
)

type PriorityQueue[T any] struct {
	pq   []T
	N    int
	less func(i, j T) bool
	mu   sync.RWMutex
}

func NewPriorityQueue[T any](less func(i, j T) bool) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		pq:   make([]T, 1), // index 0 is unused
		less: less,
	}
}

func (pq *PriorityQueue[T]) IsEmpty() bool {
	pq.mu.RLock()
	defer pq.mu.RUnlock()
	return pq.N == 0
}

func (pq *PriorityQueue[T]) Size() int {
	pq.mu.RLock()
	defer pq.mu.RUnlock()
	return pq.N
}

func (pq *PriorityQueue[T]) Peek() (T, error) {
	pq.mu.RLock()
	defer pq.mu.RUnlock()
	if pq.N == 0 { // calling IsEmpty causing deadlock
		return *new(T), errors.New(EMPTY_QUEUE)
	}
	return pq.pq[1], nil
}

func (pq *PriorityQueue[T]) Enqueue(x T) {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	pq.N++
	if pq.N >= len(pq.pq) {
		pq.pq = append(pq.pq, x)
	} else {
		pq.pq[pq.N] = x
	}
	pq.swim(pq.N)
}

func (pq *PriorityQueue[T]) Dequeue() (T, error) {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	if pq.N == 0 {
		return *new(T), errors.New(EMPTY_QUEUE)
	}
	pq.exch(1, pq.N)
	min := pq.pq[pq.N]
	pq.N--
	pq.sink(1)
	pq.pq = pq.pq[:pq.N+1] // no loitering
	return min, nil
}

func (pq *PriorityQueue[T]) PrettyPrint() string {
	var sb strings.Builder
	fst := true
	for item := range pq.GetEnumerator() {
		if !fst {
			sb.WriteString("->")
		} else {
			fst = false
		}
		sb.WriteString(fmt.Sprintf("%v", item))
	}
	return sb.String()
}

func (pq *PriorityQueue[T]) Get(index int) T {
	pq.mu.RLock()
	defer pq.mu.RUnlock()
	return pq.pq[index+1]
}

func (pq *PriorityQueue[T]) Items() []T {
	pq.mu.RLock()
	defer pq.mu.RUnlock()
	return append([]T(nil), pq.pq[1:pq.N+1]...)
}

func (pq *PriorityQueue[T]) GetEnumerator() <-chan T {
	ch := make(chan T)
	go func() {
		pq.mu.RLock()
		defer pq.mu.RUnlock()
		tempPQ := NewPriorityQueue(pq.less)
		tempPQ.pq = append(tempPQ.pq, pq.pq[1:pq.N+1]...)
		tempPQ.N = pq.N

		for !tempPQ.IsEmpty() {
			v, err := tempPQ.Dequeue()
			if err == nil {
				ch <- v
			} else {
				log.Printf("PriorityQueue->GetEnumerator: %v", err)
			}
		}
		close(ch)
	}()
	return ch
}

func (pq *PriorityQueue[T]) greater(i, j int) bool {
	return pq.less(pq.pq[j], pq.pq[i])
}

func (pq *PriorityQueue[T]) exch(i, j int) {
	pq.pq[i], pq.pq[j] = pq.pq[j], pq.pq[i]
}

// heap helper functions

func (pq *PriorityQueue[T]) swim(k int) {
	for k > 1 && pq.greater(k/2, k) {
		pq.exch(k, k/2)
		k = k / 2
	}
}

func (pq *PriorityQueue[T]) sink(k int) {
	for 2*k <= pq.N {
		j := 2 * k
		if j < pq.N && pq.greater(j, j+1) {
			j++
		}
		if !pq.greater(k, j) {
			break
		}
		pq.exch(k, j)
		k = j
	}
}
