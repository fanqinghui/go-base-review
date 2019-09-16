package main

import (
	"fmt"
	"go-base-review/4interface/shape"
)

func PrintStr(str string){
	println(str)
}

func PrintInt(num int)  {
	println(num)
}

func PrintAny(v interface{}){
	fmt.Println("any:",v)
}

func main(){

	str:="hello"
	num:=1
	PrintStr(str)
	PrintInt(num)

	PrintAny(str)
	PrintAny(num)

	c:=shape.Circle{Radius:10}
	r:=shape.Rect{Height:2,Width:5}
	shape.Measure(c)
	shape.Measure(r)

}

