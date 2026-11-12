package scheduler

import (
	"errors"
	"fmt"
	"github.com/Ajay-Ganapathy/distributed-job-scheduler/internal/job"
	"strings"
)

// JobStore defines the storage operations required by the scheduler

type JobStore interface {
	Add(j job.Job)
	FindByID(id string) (job.Job, bool)
	All() []job.Job
}

// error sentinals
var (
	ErrInvalidJobID   = errors.New("Job ID is required")
	ErrDuplicateJobID = errors.New("Job already exists")
	ErrJobNotFound    = errors.New("Job not found")
)

// Service contains the application level job operations

type Service struct {
	store JobStore
}

// New Service creates a scheduler service

func NewService(store JobStore) *Service {
	return &Service{
		store: store,
	}
}

// Submit validates and stores a new Job

func (s *Service) Submit(j job.Job) error {

	if strings.TrimSpace(j.ID) == "" {
		return ErrInvalidJobID
	}

	_, exists := s.store.FindByID(j.ID)
	if exists {
		return fmt.Errorf(
			"submit %q: %w",
			j.ID,
			ErrDuplicateJobID,
		)
	}

	s.store.Add(j)

	return nil
}

// GetByID retrieves the job by its ID

func (s *Service) GetByID(id string) (job.Job, error) {
	if strings.TrimSpace(id) == "" {
		return job.Job{}, ErrInvalidJobID
	}

	j, exists := s.store.FindByID(id)

	if !exists {
		return job.Job{}, fmt.Errorf(
			"%w: %s",
			ErrJobNotFound,
			id,
		)
	}

	return j, nil
}

// List returns all stored jobs

func (s *Service) List() []job.Job {
	return s.store.All()
}
