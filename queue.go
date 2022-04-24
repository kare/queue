package queue

import "errors"

// ErrEmptyQueue error signals that the queue is empty.
var ErrEmptyQueue = errors.New("queue is empty")

// Q is a generic FIFO queue data structure.
type Q[T any] []T

// New creates an empty queue.
func New[T any]() *Q[T] {
	return &Q[T]{}
}

// Enqueue adds a value to queue.
func (q *Q[T]) Enqueue(value T) {
	(*q) = append(*q, value)
}

// Dequeue removes the least recently added value.
// If called on an empty queue, will return zero value and ErrEmptyQueue.
func (q *Q[T]) Dequeue() (T, error) {
	var zero T
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
func (q *Q[T]) Peek() (T, error) {
	if len(*q) == 0 {
		var zero T
		return zero, ErrEmptyQueue
	}
	return (*q)[0], nil
}

// Len returns number of values in the queue.
func (q *Q[_]) Len() int {
	return len(*q)
}

// IsEmpty returns true if the queue is empty.
func (q *Q[_]) IsEmpty() bool {
	return len(*q) == 0
}

// Slice returns a copy of the queue values as a slice.
func (q *Q[T]) Slice() []T {
	r := make([]T, 0, len(*q))
	return append(r, []T(*q)...)
}
