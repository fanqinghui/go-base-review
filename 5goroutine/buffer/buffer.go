package main

import "fmt"

func main()  {
	testDeadlock()

}

/**
阻塞 chan有buffer是缓存区
 */
func testDeadlock(){
	ch:=make(chan int,2)//第二个参数是buffer的大小
	defer close(ch)
	ch<-4+4
	ch<-4+1
	val:=<-ch
	fmt.Println(val)
}
