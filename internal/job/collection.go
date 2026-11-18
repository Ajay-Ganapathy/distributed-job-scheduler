package job

// Collection stores jobs and maintains an index by ID.
type Collection struct {
	jobs  []Job
	index map[string]int
}

// Add inserts a job into the collection.

func (c *Collection) Add(j Job) {

	if c.index == nil {
		c.index = make(map[string]int)
	}

	// Upsert: If the job already exists, update it; otherwise, add it to the collection.
	if position, exists := c.index[j.ID]; exists {
		c.jobs[position] = j
		return
	}

	c.index[j.ID] = len(c.jobs)
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

// FindByID retrieves a job using the index.
func (c *Collection) FindByID(id string) (Job, bool) {

	if position, exists := c.index[id]; exists {
		return c.jobs[position], true
	}
	// for _, j := range c.jobs {
	// 	if j.ID == id {
	// 		return j, true
	// 	}
	// }

	return Job{}, false
}
