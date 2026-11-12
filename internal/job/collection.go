package job

type Collection struct {
	jobs []Job
}

// Add inserts a job into the collection.

func (c *Collection) Add(j Job) {
	c.jobs = append(c.jobs, j)
}

// Len returns the number of jobs in the collection.

func (c *Collection) Len() int {
	return len(c.jobs)
}

// All returns a copy of the stored jobs in the collection.

func (c *Collection) All() []Job {

	// Method 1: Using copy to create a new slice and copy the jobs into it
	// var dest = make([]Job , len(c.jobs))
	// copy(dest , c.jobs)
	// return dest

	// Method 2: Using append to create a new slice and append the jobs into it

	return append([]Job(nil), c.jobs...)
}

// GetByID retrieves a job by its ID. It returns the job and a boolean indicating whether the job was found.

func (c *Collection) FindByID(id string) (Job, bool) {
	for _, j := range c.jobs {
		if j.ID == id {
			return j, true
		}
	}

	return Job{}, false
}
