package queue // import "kkn.fi/queue"

// Int64Queue is a FIFO queue data structure.
type Int64Queue []int64

// NewInt64 creates an empty int64 queue.
func NewInt64() *Int64Queue {
	return &Int64Queue{}
}

// Enqueue adds a value to queue.
func (q *Int64Queue) Enqueue(value int64) {
	(*q) = append([]int64{value}, (*q)...)
}

// Dequeue removes the least recently added value.
func (q *Int64Queue) Dequeue() int64 {
	length := len((*q)) - 1
	value := (*q)[length]
	(*q) = (*q)[:length]
	return value
}

// Peek returns the least recently added value.
func (q *Int64Queue) Peek() int64 {
	return (*q)[len(*q)-1]
}

// Len returns number of values in the queue.
func (q *Int64Queue) Len() int {
	return len(*q)
}

// IsEmpty returns true if the queue is empty.
func (q *Int64Queue) IsEmpty() bool {
	return len(*q) == 0
}
