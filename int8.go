// Code generated by "gends"; DO NOT EDIT.

package queue // import "kkn.fi/queue"

// Int8 is a FIFO queue data structure.
type Int8 []int8

// NewInt8 creates an empty int8 queue.
func NewInt8() *Int8 {
	return &Int8{}
}

// Enqueue adds a value to queue.
func (q *Int8) Enqueue(value int8) {
	(*q) = append([]int8{value}, (*q)...)
}

// Dequeue removes the least recently added value.
// If called on an empty queue, will return zero value and ErrEmptyQueue.
func (q *Int8) Dequeue() (int8, error) {
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
func (q *Int8) Peek() (int8, error) {
	if len(*q) == 0 {
		return 0, ErrEmptyQueue
	}
	return (*q)[len(*q)-1], nil
}

// Len returns number of values in the queue.
func (q *Int8) Len() int {
	return len(*q)
}

// IsEmpty returns true if the queue is empty.
func (q *Int8) IsEmpty() bool {
	return len(*q) == 0
}
