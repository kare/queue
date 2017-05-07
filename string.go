package queue // import "kkn.fi/queue"

// StringQueue is a FIFO queue data structure.
type StringQueue []string

// NewString creates an empty string queue.
func NewString() *StringQueue {
	return &StringQueue{}
}

// Enqueue adds a value to queue.
func (q *StringQueue) Enqueue(value string) {
	(*q) = append([]string{value}, (*q)...)
}

// Dequeue removes the least recently added value.
func (q *StringQueue) Dequeue() string {
	length := len((*q)) - 1
	value := (*q)[length]
	(*q) = (*q)[:length]
	return value
}

// Peek returns the least recently added value.
func (q *StringQueue) Peek() string {
	return (*q)[len(*q)-1]
}

// Len returns number of values in the queue.
func (q *StringQueue) Len() int {
	return len(*q)
}

// IsEmpty returns true if the queue is empty.
func (q *StringQueue) IsEmpty() bool {
	return len(*q) == 0
}
