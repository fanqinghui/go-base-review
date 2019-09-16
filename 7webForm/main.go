package main

import (
	"net/http"
	"fmt"
	"html/template"
	"go-base-review/7webForm/data"
	"strconv"
	"time"
)

const s string = "str"

/**
构建端口为9000的web应用
 */
func main() {
	//初始化todoList值
	todoList := data.Todo{
		Title: "我的2019年待办事项",
		JobList: []data.Job{
			{Name: "学习go", Status: true},
			{Name: "学习机器学习", Status: true},
			{Name: "看书", Status: false},
		},
	}

	http.HandleFunc("/", logging(func(write http.ResponseWriter, r *http.Request) {
		fmt.Fprint(write, "欢迎来到TODOList", s)
	}))

	http.HandleFunc("/todoList", logging(func(writer http.ResponseWriter, request *http.Request) {
		tpl := template.Must(template.ParseFiles("7webForm/views/list.html"))
		if request.Method == http.MethodGet {
			tpl.Execute(writer, todoList)
			return
		}
		request.ParseForm()

		st := request.Form.Get("status")
		//fmt.Println(st)
		status, _ := strconv.ParseBool(st)

		job := data.Job{
			Name:   request.Form.Get("name"),
			Status: status,
		}
		fmt.Println(job)
		todoList.JobList=append(todoList.JobList,job)
		tpl.Execute(writer, todoList)

		return
	}))

	//设置静态文件路径
	fs := http.FileServer(http.Dir("7webForm/static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	http.ListenAndServe(":9000", nil)
}

/*
日志中间件
 */
func logging(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		beginTime := time.Now()
		defer fmt.Println(request.URL, request.Method, time.Since(beginTime))
		handlerFunc(writer, request)
	}
}
