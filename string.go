// Code generated by "gends"; DO NOT EDIT.

package queue // import "kkn.fi/queue"

// String is a FIFO queue data structure.
type String []string

// NewString creates an empty string queue.
func NewString() *String {
	return &String{}
}

// Enqueue adds a value to queue.
func (q *String) Enqueue(value string) {
	(*q) = append([]string{value}, (*q)...)
}

// Dequeue removes the least recently added value.
// If called on an empty queue, will return zero value and ErrEmptyQueue.
func (q *String) Dequeue() (string, error) {
	if len(*q) == 0 {
		return "", ErrEmptyQueue
	}
	length := len(*q) - 1
	value := (*q)[length]
	(*q) = (*q)[:length]
	return value, nil
}

// Peek returns the least recently added value.
// If called on an empty queue, will return zero value and ErrEmptyQueue.
func (q *String) Peek() (string, error) {
	if len(*q) == 0 {
		return "", ErrEmptyQueue
	}
	return (*q)[len(*q)-1], nil
}

// Len returns number of values in the queue.
func (q *String) Len() int {
	return len(*q)
}

// IsEmpty returns true if the queue is empty.
func (q *String) IsEmpty() bool {
	return len(*q) == 0
}
