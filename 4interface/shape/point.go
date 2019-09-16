package shape

import (
	"math"
	"fmt"
	"reflect"
)

type Geometry interface {
	Area() float64
	Perim() float64
}

/**
方形
 */
type Rect struct{
	Width,Height float64
}

//圆
type Circle struct {
	Radius float64
}

func (r Rect) Area() float64{
	return r.Width*r.Height
}

func (r Rect) Perim() float64{
	return 2*(r.Width+r.Height)
}

func (c Circle) Area() float64{
	return math.Pi*c.Radius*c.Radius
}

func (c Circle) Perim() float64  {
	return 2*math.Pi*c.Radius
}


func Measure(g Geometry){
	fmt.Println(reflect.TypeOf(g))
	fmt.Println(g)
	fmt.Println(g.Area())
	fmt.Println(g.Perim())
}