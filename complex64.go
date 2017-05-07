package queue // import "kkn.fi/queue"

// Complex64Queue is a FIFO queue data structure.
type Complex64Queue []complex64

// NewComplex64 creates an empty complex64 queue.
func NewComplex64() *Complex64Queue {
	return &Complex64Queue{}
}

// Enqueue adds a value to queue.
func (q *Complex64Queue) Enqueue(value complex64) {
	(*q) = append([]complex64{value}, (*q)...)
}

// Dequeue removes the least recently added value.
func (q *Complex64Queue) Dequeue() complex64 {
	length := len((*q)) - 1
	value := (*q)[length]
	(*q) = (*q)[:length]
	return value
}

// Peek returns the least recently added value.
func (q *Complex64Queue) Peek() complex64 {
	return (*q)[len(*q)-1]
}

// Len returns number of values in the queue.
func (q *Complex64Queue) Len() int {
	return len(*q)
}

// IsEmpty returns true if the queue is empty.
func (q *Complex64Queue) IsEmpty() bool {
	return len(*q) == 0
}
