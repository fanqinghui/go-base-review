package main

import "fmt"

//数据类型
// string  []切片
// bool
// int float 8-64 uint
// complex 复数


func main()  {
	s:="hello world"
	double:=3.0
	b:=false

	fmt.Println(s,s[0:5],double,b)//s[]切片

	x:=2+1i
	y:=3-2i
	fmt.Println(x*y)
	z:=x*y
	//输出 实数 与虚数部分
	fmt.Println(real(z),imag(z))
}

