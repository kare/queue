package queue

import "errors"

// ErrEmptyQueue error signals that the queue is empty.
var ErrEmptyQueue = errors.New("queue is empty")

// T is a generic FIFO queue data structure.
type T[E any] []E

// New creates an empty queue.
func New[E any]() *T[E] {
	return &T[E]{}
}

// Enqueue adds a value to queue.
func (q *T[E]) Enqueue(value E) {
	(*q) = append(*q, value)
}

// Dequeue removes the least recently added value.
// If called on an empty queue, will return zero value and ErrEmptyQueue.
func (q *T[E]) Dequeue() (E, error) {
	var zero E
	if len(*q) == 0 {
		return zero, ErrEmptyQueue
	}
	value := (*q)[0]
	(*q)[0] = zero
	(*q) = (*q)[1:]
	return value, nil
}

// Peek returns the least recently added value.
// If called on an empty queue, will return zero value and ErrEmptyQueue.
func (q *T[E]) Peek() (E, error) {
	if len(*q) == 0 {
		var zero E
		return zero, ErrEmptyQueue
	}
	return (*q)[0], nil
}

// Len returns number of values in the queue.
func (q *T[_]) Len() int {
	return len(*q)
}

// IsEmpty returns true if the queue is empty.
func (q *T[_]) IsEmpty() bool {
	return len(*q) == 0
}

// Slice returns a copy of the queue values as a slice.
func (q *T[E]) Slice() []E {
	r := make([]E, 0, len(*q))
	return append(r, []E(*q)...)
}
