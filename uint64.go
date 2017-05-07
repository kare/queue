package queue // import "kkn.fi/queue"

// Uint64Queue is a FIFO queue data structure.
type Uint64Queue []uint64

// NewUint64 creates an empty uint64 queue.
func NewUint64() *Uint64Queue {
	return &Uint64Queue{}
}

// Enqueue adds a value to queue.
func (q *Uint64Queue) Enqueue(value uint64) {
	(*q) = append([]uint64{value}, (*q)...)
}

// Dequeue removes the least recently added value.
func (q *Uint64Queue) Dequeue() uint64 {
	length := len((*q)) - 1
	value := (*q)[length]
	(*q) = (*q)[:length]
	return value
}

// Peek returns the least recently added value.
func (q *Uint64Queue) Peek() uint64 {
	return (*q)[len(*q)-1]
}

// Len returns number of values in the queue.
func (q *Uint64Queue) Len() int {
	return len(*q)
}

// IsEmpty returns true if the queue is empty.
func (q *Uint64Queue) IsEmpty() bool {
	return len(*q) == 0
}
