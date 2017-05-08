// Code generated by "gends"; DO NOT EDIT.

package queue // import "kkn.fi/queue"

// Float64Queue is a FIFO queue data structure.
type Float64Queue []float64

// NewFloat64 creates an empty float64 queue.
func NewFloat64() *Float64Queue {
	return &Float64Queue{}
}

// Enqueue adds a value to queue.
func (q *Float64Queue) Enqueue(value float64) {
	(*q) = append([]float64{value}, (*q)...)
}

// Dequeue removes the least recently added value.
// If called on an empty queue, will return zero value and ErrEmptyQueue.
func (q *Float64Queue) Dequeue() (float64, error) {
	if len(*q) == 0 {
		return 0, ErrEmptyQueue
	}
	length := len(*q) - 1
	value := (*q)[length]
	(*q) = (*q)[:length]
	return value, nil
}

// Peek returns the least recently added value.
// If called on an empty queue, will return zero value and ErrEmptyQueue.
func (q *Float64Queue) Peek() (float64, error) {
	if len(*q) == 0 {
		return 0, ErrEmptyQueue
	}
	return (*q)[len(*q)-1], nil
}

// Len returns number of values in the queue.
func (q *Float64Queue) Len() int {
	return len(*q)
}

// IsEmpty returns true if the queue is empty.
func (q *Float64Queue) IsEmpty() bool {
	return len(*q) == 0
}
