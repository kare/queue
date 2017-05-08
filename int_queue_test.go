package queue

import "testing"

func TestQueue(t *testing.T) {
	q := NewInt()
	if !q.IsEmpty() {
		t.Errorf("created queue's size should be 0, got %d", q.Len())
	}
	q.Enqueue(1)
	if q.Len() != 1 {
		t.Errorf("expected len 1, got %d", q.Len())
	}
	q.Enqueue(2)
	if q.Len() != 2 {
		t.Errorf("expected len 2, got %d", q.Len())
	}
	r, _ := q.Dequeue()
	if r != 1 {
		t.Errorf("expected to get 1 from dequeue, got %d", r)
	}
	r, _ = q.Dequeue()
	if r != 2 {
		t.Errorf("expected to get 2 from dequeue, got %d", r)
	}
	if q.Len() != 0 {
		t.Errorf("expected len 0, got %d", q.Len())
	}
}

func TestQueueDequeueEmpty(t *testing.T) {
	q := NewInt()
	zero, err := q.Dequeue()
	if zero != 0 {
		t.Errorf("expected 0, got %d", zero)
	}
	if err.Error() != "queue is empty" {
		t.Errorf("expected empty queue, got '%v'", err.Error())
	}
}

func TestQueuePeekValue(t *testing.T) {
	q := NewInt()
	q.Enqueue(1)
	q.Enqueue(2)
	one, err := q.Peek()
	if one != 1 {
		t.Errorf("expected peek to return 1, but got %d", one)
	}
	if err != nil {
		t.Errorf("expected nil error, but got '%v'", err)
	}
}

func TestQueuePeekEmpty(t *testing.T) {
	q := NewInt()
	zero, err := q.Peek()
	if zero != 0 {
		t.Errorf("expected peek to return zero value, got %d", zero)
	}
	if err.Error() != "queue is empty" {
		t.Errorf("expected error queue is empty, got '%v'", err.Error())
	}
}
