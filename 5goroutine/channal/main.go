package main

import "fmt"

//chan   <- 左侧号
//buffer
//range
//defer
//wait group


func main()  {
	ch:=make(chan int)
	defer close(ch)
	go func() {
		ch<- 3+4
	}()

	val:=<-ch
	fmt.Println(val)

}
