package queue // import "kkn.fi/queue"

// Int32Queue is a FIFO queue data structure.
type Int32Queue []int32

// NewInt32 creates an empty int32 queue.
func NewInt32() *Int32Queue {
	return &Int32Queue{}
}

// Enqueue adds a value to queue.
func (q *Int32Queue) Enqueue(value int32) {
	(*q) = append([]int32{value}, (*q)...)
}

// Dequeue removes the least recently added value.
func (q *Int32Queue) Dequeue() int32 {
	length := len((*q)) - 1
	value := (*q)[length]
	(*q) = (*q)[:length]
	return value
}

// Peek returns the least recently added value.
func (q *Int32Queue) Peek() int32 {
	return (*q)[len(*q)-1]
}

// Len returns number of values in the queue.
func (q *Int32Queue) Len() int {
	return len(*q)
}

// IsEmpty returns true if the queue is empty.
func (q *Int32Queue) IsEmpty() bool {
	return len(*q) == 0
}
