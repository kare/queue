package queue // import "kkn.fi/queue"

// Float32Queue is a FIFO queue data structure.
type Float32Queue []float32

// NewFloat32 creates an empty float32 queue.
func NewFloat32() *Float32Queue {
	return &Float32Queue{}
}

// Enqueue adds a value to queue.
func (q *Float32Queue) Enqueue(value float32) {
	(*q) = append([]float32{value}, (*q)...)
}

// Dequeue removes the least recently added value.
func (q *Float32Queue) Dequeue() float32 {
	length := len((*q)) - 1
	value := (*q)[length]
	(*q) = (*q)[:length]
	return value
}

// Peek returns the least recently added value.
func (q *Float32Queue) Peek() float32 {
	return (*q)[len(*q)-1]
}

// Len returns number of values in the queue.
func (q *Float32Queue) Len() int {
	return len(*q)
}

// IsEmpty returns true if the queue is empty.
func (q *Float32Queue) IsEmpty() bool {
	return len(*q) == 0
}
