package queue // import "kkn.fi/queue"

// Complex128Queue is a FIFO queue data structure.
type Complex128Queue []complex128

// NewComplex128 creates an empty complex128 queue.
func NewComplex128() *Complex128Queue {
	return &Complex128Queue{}
}

// Enqueue adds a value to queue.
func (q *Complex128Queue) Enqueue(value complex128) {
	(*q) = append([]complex128{value}, (*q)...)
}

// Dequeue removes the least recently added value.
func (q *Complex128Queue) Dequeue() complex128 {
	length := len((*q)) - 1
	value := (*q)[length]
	(*q) = (*q)[:length]
	return value
}

// Peek returns the least recently added value.
func (q *Complex128Queue) Peek() complex128 {
	return (*q)[len(*q)-1]
}

// Len returns number of values in the queue.
func (q *Complex128Queue) Len() int {
	return len(*q)
}

// IsEmpty returns true if the queue is empty.
func (q *Complex128Queue) IsEmpty() bool {
	return len(*q) == 0
}
