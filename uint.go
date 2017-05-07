package queue // import "kkn.fi/queue"

// UintQueue is a FIFO queue data structure.
type UintQueue []uint

// NewUint creates an empty uint queue.
func NewUint() *UintQueue {
	return &UintQueue{}
}

// Enqueue adds a value to queue.
func (q *UintQueue) Enqueue(value uint) {
	(*q) = append([]uint{value}, (*q)...)
}

// Dequeue removes the least recently added value.
func (q *UintQueue) Dequeue() uint {
	length := len((*q)) - 1
	value := (*q)[length]
	(*q) = (*q)[:length]
	return value
}

// Peek returns the least recently added value.
func (q *UintQueue) Peek() uint {
	return (*q)[len(*q)-1]
}

// Len returns number of values in the queue.
func (q *UintQueue) Len() int {
	return len(*q)
}

// IsEmpty returns true if the queue is empty.
func (q *UintQueue) IsEmpty() bool {
	return len(*q) == 0
}
