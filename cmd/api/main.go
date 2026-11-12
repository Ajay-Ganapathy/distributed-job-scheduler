package main

import (
	"fmt"
	"github.com/Ajay-Ganapathy/distributed-job-scheduler/internal/job"
	"github.com/Ajay-Ganapathy/distributed-job-scheduler/internal/scheduler"
	"log"
)

func main() {

	store := &job.Collection{}

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

	service := scheduler.NewService(store)

	if err := service.Submit(job1); err != nil {
		log.Fatal(err)
	}

	if err := service.Submit(job2); err != nil {
		log.Fatal(err)
	}

	result, err := service.GetByID("job-001")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Retrieved job: %+v\n", result)
	fmt.Println("All jobs in the store:", service.List())

	// collection.Add(job1)
	// collection.Add(job2)

	// fmt.Printf("total Jobs in Collection = %d\n", collection.Len())
	// fmt.Println("All Jobs in Collection = ", collection.All())

	// for index, job := range collection.All() {
	// 	fmt.Printf(" Job Index :  %d , Job = %+v\n", index, job)
	// }

	// var job, found = collection.FindByID("job-001")

	// if !found {
	// 	fmt.Println("Job Not Found")
	// } else {
	// 	fmt.Printf("Job Found : %+v\n", job)
	// }

	// fmt.Printf("Job 1 : %+v\n", job1)
	// fmt.Printf("Job 2 : %+v\n", job2)
}
