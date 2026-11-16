package job

import "errors"


var (
	ErrInvalidCapacity = errors.New("queue cpacity must be greater than zero")
	ErrQueueFull       = errors.New("queue is full")
)

// Bounded Queue stores jobs in FIFO Order

type BoundedQueue struct {
	items    []Job
	capacity int
}

// NewBoundedQueue creates a new BoundedQueue with the specified capacity.

func NewBoundedQueue(capacity int) (*BoundedQueue, error) {
	if capacity <= 0 {
		return nil, ErrInvalidCapacity
	}

	return &BoundedQueue{
		items:    make([]Job, 0 ,  capacity),
		capacity: capacity,
	}, nil
}

// EnQueue adds a job to the end of the queue. If the queue is full, it returns an error.

func (q *BoundedQueue) EnQueue(j Job) error {
	if len(q.items) >= q.capacity {
		return ErrQueueFull
	}

	q.items = append(q.items, j)

	

	return nil
}

// DeQueue removes the oldest Job

func (q *BoundedQueue) DeQueue() (Job, bool) {
	if len(q.items) == 0 {
		return Job{}, false
	}

	j := q.items[0]

	copy(q.items, q.items[1:])
	// Clear the last element to release references.
	q.items[len(q.items)-1] = Job{}
	q.items = q.items[:len(q.items)-1]

	return j, true
}

func (q *BoundedQueue) Len() int {
	return len(q.items)
}

func (q *BoundedQueue) Capacity() int {
	return q.capacity
}

func (q *BoundedQueue) isEmpty() bool {
	return len(q.items) == 0
}

func (q *BoundedQueue) isFull() bool {
	return len(q.items) >= q.capacity
}
