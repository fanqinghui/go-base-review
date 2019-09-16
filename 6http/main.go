package main

import (
	"net/http"
	"log"
	"html/template"
	"go-base-review/6http/data"
)

//hello world
//template
//struct
//iteration
func main(){
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		//fmt.Fprint(writer,"this is index page")
		index:=template.Must(template.ParseFiles("6http/index.html"))
		index.Execute(writer,initIndexPageData())

	})
	log.Fatal(http.ListenAndServe(":9999",nil))
}


func initIndexPageData() data.Todo{
	return data.Todo{
		Title: "2019年待办事项",
		Job:   []data.Job{
			{"学习", false},
			{Name:"learn",Status:false},
			{Name:"读书",Status:true},
		},
	}
}