package main

import (
	"go-base-review/3class/one"
	"math"
	"fmt"
)

func main(){
	a:=one.Point{X:1,Y:3}
	b:=one.Point{X:11,Y:-11.1}
	fmt.Println(Distance(a,b))
	fmt.Println(a.Distance(b))
}

func Distance(a,b one.Point) float64{
	return math.Hypot(a.X-b.X,a.Y-b.Y)
}