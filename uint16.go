// Code generated by "gends"; DO NOT EDIT.

package queue // import "kkn.fi/queue"

// Uint16Queue is a FIFO queue data structure.
type Uint16Queue []uint16

// NewUint16 creates an empty uint16 queue.
func NewUint16() *Uint16Queue {
	return &Uint16Queue{}
}

// Enqueue adds a value to queue.
func (q *Uint16Queue) Enqueue(value uint16) {
	(*q) = append([]uint16{value}, (*q)...)
}

// Dequeue removes the least recently added value.
// If called on an empty queue, will return zero value and ErrEmptyQueue.
func (q *Uint16Queue) Dequeue() (uint16, error) {
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
func (q *Uint16Queue) Peek() (uint16, error) {
	if len(*q) == 0 {
		return 0, ErrEmptyQueue
	}
	return (*q)[len(*q)-1], nil
}

// Len returns number of values in the queue.
func (q *Uint16Queue) Len() int {
	return len(*q)
}

// IsEmpty returns true if the queue is empty.
func (q *Uint16Queue) IsEmpty() bool {
	return len(*q) == 0
}
