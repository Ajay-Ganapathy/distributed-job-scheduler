package scheduler

import (
	"errors"
	"github.com/Ajay-Ganapathy/distributed-job-scheduler/internal/job"
	"testing"
)

// Compile-time verification that Collection satisfies JobStore.
var _ JobStore = (*job.Collection)(nil)

func newTestService() *Service {
	store := &job.Collection{}
	return NewService(store)
}

func TestSubmitAndGetByID(t *testing.T) {
	service := newTestService()

	j := job.NewJob(
		"job-001",
		"embedding",
		"Generate embeddings",
		"doc-001",
		5,
	)

	if err := service.Submit(j); err != nil {
		t.Fatalf("unexpected submit error: %v", err)
	}

	result, err := service.GetByID("job-001")

	if err != nil {
		t.Fatalf("unexpected get error: %v", err)
	}

	if result.ID != "job-001" {
		t.Errorf("expected job-001, got %s", result.ID)
	}
}

func TestSubmitRejectsDuplicate(t *testing.T) {
	service := newTestService()

	j := job.NewJob(
		"job-001",
		"embedding",
		"Generate embeddings",
		"doc-001",
		5,
	)

	if err := service.Submit(j); err != nil {
		t.Fatal(err)
	}

	err := service.Submit(j)

	if !errors.Is(err, ErrDuplicateJobID) {
		t.Fatalf("expected duplicate error, got %v", err)
	}
}

func TestSubmitRejectsEmptyID(t *testing.T) {
	service := newTestService()

	j := job.NewJob(
		"",
		"embedding",
		"Generate embeddings",
		"doc-001",
		5,
	)

	err := service.Submit(j)

	if !errors.Is(err, ErrInvalidJobID) {
		t.Fatalf("expected invalid ID error, got %v", err)
	}
}

func TestGetByIDMissing(t *testing.T) {
	service := newTestService()

	_, err := service.GetByID("missing")

	if !errors.Is(err, ErrJobNotFound) {
		t.Fatalf("expected not found error, got %v", err)
	}
}
