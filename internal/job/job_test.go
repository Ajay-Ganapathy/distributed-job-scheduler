package job

import "testing"
import "encoding/json"
import "errors"

func TestNewJob(t *testing.T) {
	j := NewJob("job-001", "embedding-generation", "Generate embeddings for documents", "document-123", 5)

	if j.ID != "job-001" {
		t.Errorf("Expected job ID to be 'job-001', but got '%s'", j.ID)
	}

	if j.Status != Pending {
		t.Errorf("Expected job status to be 'PENDING', but got '%s'", j.Status)

	}

	if j.Priority != 5 {
		t.Errorf("Expected job priority to be 5, but got '%d'", j.Priority)
	}

	if j.CreatedAt.IsZero() {
		t.Error("Expected job CreatedAt to be set, but it is zero")
	}

}

func TestJobJSONSerialization(t *testing.T) {
	job := NewJob(
		"job-001",
		"embedding-generation",
		"Generate embeddings for documents",
		"document-123",
		5,
	)

	data, err := json.Marshal(job)

	if err != nil {
		t.Fatalf("Failed to serialize job to JSON: %v", err)
	}

	var result map[string]interface{}

	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("failed to unmarshal JSON: %v", err)
	}

	if _, exists := result["id"]; !exists {
		t.Fatal("expected JSON field 'id' to exist")
	}

	if _, exists := result["ID"]; exists {
		t.Fatal("unexpected JSON field 'ID'; expected lowercase 'id'")
	}

	if result["status"] != "PENDING" {
		t.Fatalf("expected status 'PENDING', got '%v'", result["status"])
	}

	if _, exists := result["created_at"]; !exists {
		t.Fatal("expected JSON field 'created_at' to exist")
	}

	t.Logf("Serialized Job: %s", string(data))
}

func TestJobValidTransitions(t *testing.T) {
	j := NewJob(
		"job-001",
		"embedding",
		"Generate embeddings",
		"doc-001",
		5,
	)

	if err := j.TransitionTo(Running); err != nil {
		t.Fatalf("unexpected transition error: %v", err)
	}

	if j.Status != Running {
		t.Fatalf("expected RUNNING, got %s", j.Status)
	}

	if err := j.TransitionTo(Completed); err != nil {
		t.Fatalf("unexpected transition error: %v", err)
	}

	if j.Status != Completed {
		t.Fatalf("expected COMPLETED, got %s", j.Status)
	}
}

func TestJobInvalidTransition(t *testing.T) {
	j := NewJob(
		"job-001",
		"embedding",
		"Generate embeddings",
		"doc-001",
		5,
	)

	err := j.TransitionTo(Completed)

	if !errors.Is(err, ErrInvalidStatusTransition) {
		t.Fatalf("expected invalid transition error, got %v", err)
	}

	if j.Status != Pending {
		t.Fatalf("expected PENDING after failed transition, got %s", j.Status)
	}
}

func TestTerminalJobCannotTransition(t *testing.T) {
	j := NewJob(
		"job-001",
		"embedding",
		"Generate embeddings",
		"doc-001",
		5,
	)

	if err := j.TransitionTo(Running); err != nil {
		t.Fatal(err)
	}

	if err := j.TransitionTo(Completed); err != nil {
		t.Fatal(err)
	}

	if err := j.TransitionTo(Running); !errors.Is(err, ErrInvalidStatusTransition) {
		t.Fatalf("expected invalid transition, got %v", err)
	}
}
