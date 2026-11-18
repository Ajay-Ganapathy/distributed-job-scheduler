package job

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestNewWorker(t *testing.T) {
	worker := NewWorker(
		"worker-001",
		WorkerCapabilities{
			CPUCores: 8,
			MemoryMB: 16384,
			GPUCount: 1,
		},
	)

	if worker.ID != "worker-001" {
		t.Fatalf("expected worker-001, got %s", worker.ID)
	}

	if worker.Status != WorkerIdle {
		t.Fatalf("expected IDLE, got %s", worker.Status)
	}

	if worker.Capabilities.GPUCount != 1 {
		t.Fatalf("expected one GPU, got %d", worker.Capabilities.GPUCount)
	}
}

func TestWorkerJSON(t *testing.T) {
	worker := NewWorker(
		"worker-001",
		WorkerCapabilities{
			CPUCores: 8,
			MemoryMB: 16384,
		},
	)

	data, err := json.Marshal(worker)

	if err != nil {
		t.Fatal(err)
	}

	result := string(data)

	if !strings.Contains(result, `"id":"worker-001`) {
		t.Fatalf("expected JSON id field, got %s", result)
	}

	if strings.Contains(result, "current_job_id") {
		t.Fatalf("expected empty current_job_id to be omitted: %s", result)
	}
}
