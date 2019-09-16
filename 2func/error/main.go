package main

import "fmt"

//panic recover
func main(){
	test()
}

func test(){
	defer func() {
		if p:=recover();p!=nil{
			fmt.Println(p)
		}
	}()

	fmt.Println("call")
	call()
	fmt.Println("end")
}

func call(){
	panic("call bad!!!")
}


