package main

import "fmt"


func add(x,y int) int{
	return x+y
}

func swap(x,y int) (int,int){
	return y,x
}

func swap2(x,y int) (yy,xx int){
	return y,x
}

func main(){
	fmt.Println(add(1,2))//3

	fmt.Println(swap(3,5))//5,3
}
