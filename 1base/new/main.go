package main

import "fmt"

//point
//new

func main(){
	//& *   *类型  代表形式参数
	//i:=1 等于
	ii:=new(int)
	ii_val:=*ii
	fmt.Println(ii,ii_val)

	i:=1
	p:=&i
	fmt.Println(i,p,*p)
	*p++
	fmt.Println(i)

	incr_1(i)
	fmt.Println(i)
	incr_2(&i)
	fmt.Println(i)
}


func incr_1(val int) int{
	val++
	return val
}

func incr_2(val *int){
	*val++
}
