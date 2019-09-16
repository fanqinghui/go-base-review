package data

type Job struct {
	Name string
	Status bool
}

type Todo struct{
	Title string
	Job []Job
}


func getData() Todo{
	return Todo{
		Title:"2019目标任务",
		Job:[]Job{
			{Name:"上学",Status:false},
			{Name:"读书",Status:true},
		},
	}
}