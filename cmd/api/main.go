package main

import (
	"fmt"
	"github.com/Ajay-Ganapathy/distributed-job-scheduler/internal/job"
)

func main() {
	job1 := job.NewJob(
		"job-001",
		"embedding-generation",
		"Generate embeddings for documents",
		"document-123",
		5,
	)

	job2 := job.NewJob(
		"job-002",
		"document-summarization",
		"Generate a document summary",
		"document-456",
		3,
	)

	fmt.Printf("Job 1 : %+v\n", job1)
	fmt.Printf("Job 2 : %+v\n", job2)
}
