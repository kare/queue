package queue

import "testing"

func TestQueueLen(t *testing.T) {
	q := NewInt()
	if !q.IsEmpty() {
		t.Error("Newly created queue's size should be 0. Got %d.", q.Len())
	}
	q.Enqueue(1)
	if q.Len() != 1 {
		t.Errorf("Expected len 1. Got %d", q.Len())
	}
	q.Enqueue(2)
	if q.Len() != 2 {
		t.Errorf("Expected len 2. Got %d.", q.Len())
	}
	r := q.Dequeue()
	if r != 1 {
		t.Errorf("Expected to get 1 from dequeue. Got %d.", r)
	}
	r = q.Dequeue()
	if r != 2 {
		t.Errorf("Expected to get 2 from dequeue. Got %d.", r)
	}
	if q.Len() != 0 {
		t.Errorf("Expected len 0. Got %d.", q.Len())
	}
}
