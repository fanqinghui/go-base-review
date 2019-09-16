package main

import (
	"go-base-review/1base/dataType2/shape"
	"fmt"
)

func main()  {
	p:=shape.Point{1,2}
	//p:=shape.Point{X:1,Y:2}
	fmt.Println(p)
	p.X=3
	fmt.Println(p)
	p_add:=&p
	fmt.Println(p_add)//结构体p的地址
	fmt.Println(*p_add,(*p_add).X)//结构体地址的取值操作
	fmt.Println(p_add.X)//其实结构体的地址直接取属性值也可以取到对应的属性值
	//测试组合
	student:=shape.Student{ID:1,Other:shape.People{Name:"zj",Age:30}}
	fmt.Println(student)
}