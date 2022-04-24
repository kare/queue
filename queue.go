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
	(*q) = append([]T{value}, (*q)...)
}

// Dequeue removes the least recently added value.
// If called on an empty queue, will return zero value and ErrEmptyQueue.
func (q *Q[T]) Dequeue() (T, error) {
	if len(*q) == 0 {
		var zero T
		return zero, ErrEmptyQueue
	}
	length := len(*q) - 1
	value := (*q)[length]
	(*q) = (*q)[:length]
	return value, nil
}

// Peek returns the least recently added value.
// If called on an empty queue, will return zero value and ErrEmptyQueue.
func (q *Q[T]) Peek() (T, error) {
	if len(*q) == 0 {
		var zero T
		return zero, ErrEmptyQueue
	}
	return (*q)[len(*q)-1], nil
}

// Len returns number of values in the queue.
func (q *Q[_]) Len() int {
	return len(*q)
}

// IsEmpty returns true if the queue is empty.
func (q *Q[_]) IsEmpty() bool {
	return len(*q) == 0
}
