package main

import (
	"net/http"
	"fmt"
	"html/template"
	"go-base-review/6http/data"
	"time"
)

func main(){
	http.HandleFunc("/goodLuck", loging(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer,"good luck")
	}))

	http.HandleFunc("/todoList", loging( func(writer http.ResponseWriter, request *http.Request) {
		tp:=template.Must(template.ParseFiles("6http/test/totoList.html"))

		tp.Execute(writer, initData())
	}))
	//静态文件服务器
	fs:=http.FileServer(http.Dir("6http/test/static"))
	http.Handle("/static/",http.StripPrefix("/static/",fs))

	http.ListenAndServe(":7778",nil)
}

func loging(handler http.HandlerFunc) http.HandlerFunc{
	return func(writer http.ResponseWriter, request *http.Request) {
		  t:=time.Now()
		  defer fmt.Println(request.URL,request.Method,time.Since(t))
		  handler(writer,request)
	}
}


func initData() data.Todo {
		return data.Todo{
			Title:"learn",
			Job: []data.Job{
				{Name:"学习go",Status:true},
				{Name:"学习机器学习",Status:false},
			},
		}
}