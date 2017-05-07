package queue // import "kkn.fi/queue"

// Int16Queue is a FIFO queue data structure.
type Int16Queue []int16

// NewInt16 creates an empty int16 queue.
func NewInt16() *Int16Queue {
	return &Int16Queue{}
}

// Enqueue adds a value to queue.
func (q *Int16Queue) Enqueue(value int16) {
	(*q) = append([]int16{value}, (*q)...)
}

// Dequeue removes the least recently added value.
func (q *Int16Queue) Dequeue() int16 {
	length := len((*q)) - 1
	value := (*q)[length]
	(*q) = (*q)[:length]
	return value
}

// Peek returns the least recently added value.
func (q *Int16Queue) Peek() int16 {
	return (*q)[len(*q)-1]
}

// Len returns number of values in the queue.
func (q *Int16Queue) Len() int {
	return len(*q)
}

// IsEmpty returns true if the queue is empty.
func (q *Int16Queue) IsEmpty() bool {
	return len(*q) == 0
}
