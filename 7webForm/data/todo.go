package data

type Job struct {
	Name string
	Status bool
}

type Todo struct {
	Title string
	JobList []Job
}
