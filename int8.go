package queue // import "kkn.fi/queue"

// Int8Queue is a FIFO queue data structure.
type Int8Queue []int8

// NewInt8 creates an empty int8 queue.
func NewInt8() *Int8Queue {
	return &Int8Queue{}
}

// Enqueue adds a value to queue.
func (q *Int8Queue) Enqueue(value int8) {
	(*q) = append([]int8{value}, (*q)...)
}

// Dequeue removes the least recently added value.
func (q *Int8Queue) Dequeue() int8 {
	length := len((*q)) - 1
	value := (*q)[length]
	(*q) = (*q)[:length]
	return value
}

// Peek returns the least recently added value.
func (q *Int8Queue) Peek() int8 {
	return (*q)[len(*q)-1]
}

// Len returns number of values in the queue.
func (q *Int8Queue) Len() int {
	return len(*q)
}

// IsEmpty returns true if the queue is empty.
func (q *Int8Queue) IsEmpty() bool {
	return len(*q) == 0
}
