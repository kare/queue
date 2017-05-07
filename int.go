package queue // import "kkn.fi/queue"

// IntQueue is a FIFO queue data structure.
type IntQueue []int

// NewInt creates an empty int queue.
func NewInt() *IntQueue {
	return &IntQueue{}
}

// Enqueue adds a value to queue.
func (q *IntQueue) Enqueue(value int) {
	(*q) = append([]int{value}, (*q)...)
}

// Dequeue removes the least recently added value.
func (q *IntQueue) Dequeue() int {
	length := len((*q)) - 1
	value := (*q)[length]
	(*q) = (*q)[:length]
	return value
}

// Peek returns the least recently added value.
func (q *IntQueue) Peek() int {
	return (*q)[len(*q)-1]
}

// Len returns number of values in the queue.
func (q *IntQueue) Len() int {
	return len(*q)
}

// IsEmpty returns true if the queue is empty.
func (q *IntQueue) IsEmpty() bool {
	return len(*q) == 0
}
