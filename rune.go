package queue // import "kkn.fi/queue"

// RuneQueue is a FIFO queue data structure.
type RuneQueue []rune

// NewRune creates an empty rune queue.
func NewRune() *RuneQueue {
	return &RuneQueue{}
}

// Enqueue adds a value to queue.
func (q *RuneQueue) Enqueue(value rune) {
	(*q) = append([]rune{value}, (*q)...)
}

// Dequeue removes the least recently added value.
func (q *RuneQueue) Dequeue() rune {
	length := len((*q)) - 1
	value := (*q)[length]
	(*q) = (*q)[:length]
	return value
}

// Peek returns the least recently added value.
func (q *RuneQueue) Peek() rune {
	return (*q)[len(*q)-1]
}

// Len returns number of values in the queue.
func (q *RuneQueue) Len() int {
	return len(*q)
}

// IsEmpty returns true if the queue is empty.
func (q *RuneQueue) IsEmpty() bool {
	return len(*q) == 0
}
