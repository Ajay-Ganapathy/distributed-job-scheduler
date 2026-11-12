package job

import "testing"

func TestCollectionAddAndLen(t *testing.T) {
	collection := Collection{}
	var job1 = NewJob("job-001", "embedding-generation", "Generate embeddings for documents", "document-123", 5)
	collection.Add(job1)

	if collection.Len() != 1 {
		t.Fatalf("expected 1 job, got %d", collection.Len())
	}

}

func TestCollectionFindByID(t *testing.T) {
	collection := Collection{}
	var job1 = NewJob("job-001", "embedding-generation", "Generate embeddings for documents", "document-123", 5)
	collection.Add(job1)

	job, found := collection.FindByID("job-001")

	if !found {
		t.Fatal("expected to find job with ID 'job-001', but it was not found")
	}

	if job.ID != "job-001" {
		t.Fatalf("expected job ID to be 'job-001', but got '%s'", job.ID)
	}
}

func TestFindMissingJobByID(t *testing.T) {
	collection := Collection{}

	var job1 = NewJob("job-001", "embedding-generation", "Generate embeddings for documents", "document-123", 5)
	collection.Add(job1)
	_, found := collection.FindByID("missing")

	if found {
		t.Fatal("expected not to find job with ID 'missing', but it was found")
	}
}

func TestCollectionAllReturnsCopy(t *testing.T) {
	collection := Collection{}

	collection.Add(NewJob(
		"job-001",
		"embedding",
		"Generate embeddings",
		"doc-001",
		5,
	))

	result := collection.All()
	result[0].Name = "modified"

	original, ok := collection.FindByID("job-001")

	if !ok {
		t.Fatal("expected original job to exist")
	}

	if original.Name == "modified" {
		t.Fatal("modifying returned slice changed internal collection")
	}
}
