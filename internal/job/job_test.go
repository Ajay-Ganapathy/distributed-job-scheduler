package job

import "testing"

func TestNewJob(t *testing.T){
	j := NewJob("job-001", "embedding-generation", "Generate embeddings for documents", "document-123", 5)

	if(j.ID != "job-001"){
		t.Errorf("Expected job ID to be 'job-001', but got '%s'", j.ID)
	}

	if(j.Status != Pending){
		t.Errorf("Expected job status to be 'PENDING', but got '%s'", j.Status)

	}

	if(j.Priority != 5){
		t.Errorf("Expected job priority to be 5, but got '%d'", j.Priority)
	}

	if(j.CreatedAt.IsZero()){
		t.Error("Expected job CreatedAt to be set, but it is zero")
	}


}