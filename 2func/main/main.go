package main

//函数定义
//错误处理
//返回值
import "fmt"


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
//返回函数
func returnFunc(a string,b int) (func() (string,int,error)){
	return func() (string, int, error) {
		if(b!=0){
			return a, b, nil
		}else {
			return "",0,fmt.Errorf("b must be no-zero num!")
		}
	}
}

func main(){

	f := returnFunc("1", 3)
	//fmt.Println(f)
	a,b,c:=f()
	if(c!=nil){
		fmt.Println(c)
	}else{
		fmt.Println(a,b)
	}



	//fmt.Println(add(1,2))//3

	//fmt.Println(swap(3,5))//5,3

	//fmt.Println(swap2(4,5))//5,4
}
