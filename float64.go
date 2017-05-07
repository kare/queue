package queue // import "kkn.fi/queue"

// Float64Queue is a FIFO queue data structure.
type Float64Queue []float64

// NewFloat64 creates an empty float64 queue.
func NewFloat64() *Float64Queue {
	return &Float64Queue{}
}

// Enqueue adds a value to queue.
func (q *Float64Queue) Enqueue(value float64) {
	(*q) = append([]float64{value}, (*q)...)
}

// Dequeue removes the least recently added value.
func (q *Float64Queue) Dequeue() float64 {
	length := len((*q)) - 1
	value := (*q)[length]
	(*q) = (*q)[:length]
	return value
}

// Peek returns the least recently added value.
func (q *Float64Queue) Peek() float64 {
	return (*q)[len(*q)-1]
}

// Len returns number of values in the queue.
func (q *Float64Queue) Len() int {
	return len(*q)
}

// IsEmpty returns true if the queue is empty.
func (q *Float64Queue) IsEmpty() bool {
	return len(*q) == 0
}
