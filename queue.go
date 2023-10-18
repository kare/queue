package queue

import "errors"

// ErrEmptyQueue error signals that the queue is empty.
var ErrEmptyQueue = errors.New("queue: queue is empty")

// Simple is a generic FIFO queue data structure.
type Simple[E any] struct {
	queue []E
}

// NewSimple creates an empty queue.
func NewSimple[E any]() *Simple[E] {
	return &Simple[E]{}
}

// Enqueue adds a value to queue.
func (q *Simple[E]) Enqueue(value E) {
	q.queue = append(q.queue, value)
}

// Dequeue removes the least recently added value.
// If called on an empty queue, will return zero value and ErrEmptyQueue.
func (q *Simple[E]) Dequeue() (E, error) {
	var zero E
	if len(q.queue) == 0 {
		return zero, ErrEmptyQueue
	}
	value := q.queue[0]
	q.queue[0] = zero
	q.queue = q.queue[1:]
	return value, nil
}

// Peek returns the least recently added value.
// If called on an empty queue, will return zero value and ErrEmptyQueue.
func (q *Simple[E]) Peek() (E, error) {
	if len(q.queue) == 0 {
		var zero E
		return zero, ErrEmptyQueue
	}
	return q.queue[0], nil
}

// Len returns number of values in the queue.
func (q *Simple[_]) Len() int {
	return len(q.queue)
}

// IsEmpty returns true if the queue is empty.
func (q *Simple[_]) IsEmpty() bool {
	return len(q.queue) == 0
}

// Slice returns a copy of the queue values as a slice.
func (q *Simple[E]) Slice() []E {
	r := make([]E, 0, len(q.queue))
	return append(r, []E(q.queue)...)
}
