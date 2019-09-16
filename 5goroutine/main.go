package main

import (
	"fmt"
	"time"
)

func f(){
	for i:=0;;i++{
		fmt.Println(i)
		time.Sleep(1*time.Second)
	}
}

func fib(n int) int{
	if n<2{
		return n
	}
	return fib(n-1)+fib(n-2)
}

func main()  {
	go f()
	fib(45)
}
