package job

type WorkerStatus string

const (
	WorkerIdle    WorkerStatus = "IDLE"
	WorkerBusy    WorkerStatus = "BUSY"
	WorkerStopped WorkerStatus = "STOPPED"
)

// Worker Capabilities describe the resources available to a Worker

type WorkerCapabilities struct {
	CPUCores int `json:"cpu_cores"`
	MemoryMB int `json:"memory_mb"`
	GPUCount int `json:"gpu_count,omitempty"`
}

type Worker struct {
	ID           string             `json:"id"`
	Status       WorkerStatus       `json:"status"`
	CurrentJobID string             `json:"current_job_id,omitempty"`
	Capabilities WorkerCapabilities `json:"capabilities"`
}

// NewWorker creates an idle worker.

func NewWorker(
	id string,
	capabilities WorkerCapabilities,
) Worker {
	return Worker{
		ID:           id,
		Status:       WorkerIdle,
		CurrentJobID: "",
		Capabilities: capabilities,
	}
}
