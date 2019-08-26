# 函数
 函数可以没有参数，或者接受多个参数，同时函数可以没有返回值，也可以有多个返回值
 
 ```
 func add(x,y int) int{
	return x+y
}

//演示了返回值定义的形式
func add2(x,y float64) (result float64){
	result=x+y
	return
}

//多返回值演示
func swap(x,y int) (int,int){
	return y,x
}
//多返回值演示2 --推荐上面的形式
func swap2(x,y int) (yy,xx int){
	return y,x
}
 
 ```