package main

import "fmt"

// 与java命名一致，命名：驼峰
// 避免关键字
// var变量 const常量 变量声明了必须进行使用，否则程序编译不成功

func main()  {
	//const,var,type,func
	const c1,c2=1,2

	//var
	//	 var 变量名 类型=表达式
	//   var 变量名 类型
	//   变量名 :=表达式   --自动推断类型
	var a int=3
	var b string
	c:=true
	d:=2.0
	fmt.Println(a,b,c,d)
	foo(c1)
	a=square(a)
	foo(a)
}

func foo(x int){
	fmt.Println(x)
}

func square(x int) int{
	return x*x
}
