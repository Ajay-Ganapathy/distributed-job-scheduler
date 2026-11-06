package job

import "time"

type Status string

const (
	Pending   Status = "PENDING"
	Running   Status = "RUNNING"
	Completed Status = "COMPLETED"
	Failed    Status = "FAILED"
)

type Job struct {
	ID          string    `json:id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Payload     string    `json:"payload"`
	Status      Status    `json:"status"`
	Priority    int       `json:"priority"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewJob(id, name, description, payload string, priority int) Job {
	return Job{
		ID:          id,
		Name:        name,
		Description: description,
		Payload:     payload,
		Status:      Pending,
		Priority:    priority,
		CreatedAt:   time.Now(),
	}
}
