package job

import (
	"errors"
	"testing"
)

func makeTestJob(id string) Job {
	return NewJob(
		id,
		"embedding-generation",
		"Generate embeddings",
		"document-123",
		5,
	)
}

func TestNewBoundedQueue(t *testing.T) {
	q, err := NewBoundedQueue(3)

	if err != nil {
		t.Fatal(err)
	}

	if q.Capacity() != 3 {
		t.Fatalf("Expected capacity 3 , but got %d", q.Capacity())
	}
}

func TestNewBoundedQueueRejectsInvalidCapacity(t *testing.T) {
	_ , err := NewBoundedQueue(0)

	if err == nil {
		t.Fatalf("Expected error for invalid capacity, but got nil")
	}
}

func TestQueueFIFO(t *testing.T) {
	q, err := NewBoundedQueue(3)

	if err != nil {
		t.Fatal(err)
	}

	_ = q.EnQueue(makeTestJob("job-1"))
	_ = q.EnQueue(makeTestJob("job-2"))
	_ = q.EnQueue(makeTestJob("job-3"))

	first, ok := q.DeQueue()

	if !ok || first.ID != "job-1" {
		t.Fatalf("Expected job-1 , but got %v", first)
	}

	second, ok := q.DeQueue()

	if !ok || second.ID != "job-2" {
		t.Fatalf("expected job-002 as the second job, but got %v", second)
	}
}

func TestQueueFull(t *testing.T) {
	q, _ := NewBoundedQueue(1)

	if err := q.EnQueue(makeTestJob("job-001")); err != nil {
		t.Fatal(err)
	}

	err := q.EnQueue(makeTestJob("job-002"))

	if !errors.Is(err, ErrQueueFull) {
		t.Fatalf("expected queue full error, got %v", err)
	}
}

func TestDeQueueEmptyQueue(t *testing.T) {

	q, err := NewBoundedQueue(1)
	if err != nil {
		t.Fatal(err)
	}

	_, ok := q.DeQueue()

	if ok {
		t.Fatal("expected dequeue from empty queue to return false")
	}
}

func EnqueueThreeJobsDequeueTwo(t *testing.T) {
	q, err := NewBoundedQueue(3)
	if err != nil {
		t.Fatal(err)
	}

	_ = q.EnQueue(makeTestJob("job-1"))
	_ = q.EnQueue(makeTestJob("job-2"))
	_ = q.EnQueue(makeTestJob("job-3"))

	j, ok := q.DeQueue()
	if !ok {
		t.Fatal("expected dequeue to succeed")
	}

	if j.ID != "job-1" {
		t.Fatalf("expected job-1 , but got %s", j.ID)
	}

	j2, ok := q.DeQueue()
	if !ok {
		t.Fatal("expected dequeue to succeed")
	}

	if j2.ID != "job-2" {
		t.Fatalf("expected job-2 , but got %s", j2.ID)
	}

}
