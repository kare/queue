package queue // import "kkn.fi/queue"

// Uint8Queue is a FIFO queue data structure.
type Uint8Queue []uint8

// NewUint8 creates an empty uint8 queue.
func NewUint8() *Uint8Queue {
	return &Uint8Queue{}
}

// Enqueue adds a value to queue.
func (q *Uint8Queue) Enqueue(value uint8) {
	(*q) = append([]uint8{value}, (*q)...)
}

// Dequeue removes the least recently added value.
func (q *Uint8Queue) Dequeue() uint8 {
	length := len((*q)) - 1
	value := (*q)[length]
	(*q) = (*q)[:length]
	return value
}

// Peek returns the least recently added value.
func (q *Uint8Queue) Peek() uint8 {
	return (*q)[len(*q)-1]
}

// Len returns number of values in the queue.
func (q *Uint8Queue) Len() int {
	return len(*q)
}

// IsEmpty returns true if the queue is empty.
func (q *Uint8Queue) IsEmpty() bool {
	return len(*q) == 0
}
