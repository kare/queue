package queue // import "kkn.fi/queue"

// ByteQueue is a FIFO queue data structure.
type ByteQueue []byte

// NewByte creates an empty byte queue.
func NewByte() *ByteQueue {
	return &ByteQueue{}
}

// Enqueue adds a value to queue.
func (q *ByteQueue) Enqueue(value byte) {
	(*q) = append([]byte{value}, (*q)...)
}

// Dequeue removes the least recently added value.
func (q *ByteQueue) Dequeue() byte {
	length := len((*q)) - 1
	value := (*q)[length]
	(*q) = (*q)[:length]
	return value
}

// Peek returns the least recently added value.
func (q *ByteQueue) Peek() byte {
	return (*q)[len(*q)-1]
}

// Len returns number of values in the queue.
func (q *ByteQueue) Len() int {
	return len(*q)
}

// IsEmpty returns true if the queue is empty.
func (q *ByteQueue) IsEmpty() bool {
	return len(*q) == 0
}
