package queue // import "kkn.fi/queue"

// Uint32Queue is a FIFO queue data structure.
type Uint32Queue []uint32

// NewUint32 creates an empty uint32 queue.
func NewUint32() *Uint32Queue {
	return &Uint32Queue{}
}

// Enqueue adds a value to queue.
func (q *Uint32Queue) Enqueue(value uint32) {
	(*q) = append([]uint32{value}, (*q)...)
}

// Dequeue removes the least recently added value.
func (q *Uint32Queue) Dequeue() uint32 {
	length := len((*q)) - 1
	value := (*q)[length]
	(*q) = (*q)[:length]
	return value
}

// Peek returns the least recently added value.
func (q *Uint32Queue) Peek() uint32 {
	return (*q)[len(*q)-1]
}

// Len returns number of values in the queue.
func (q *Uint32Queue) Len() int {
	return len(*q)
}

// IsEmpty returns true if the queue is empty.
func (q *Uint32Queue) IsEmpty() bool {
	return len(*q) == 0
}
