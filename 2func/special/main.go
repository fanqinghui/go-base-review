package main

import (
	"fmt"
	"sync"
	"strconv"
)

//三个特殊的函数
//defer  延迟执行、、程序在return之前，报错之前进行执行，文件操作关闭刘，数据库关闭连接等操作场景
//panic 错误处理，类似java里的异常throw
//recover
func main(){
	fmt.Println("main")
	//deferTest()
	fmt.Println(lookUp(1))
	fmt.Println(lookUp(4))
	fmt.Println(lookUp(6))
}

func deferTest(){
	fmt.Println("a")
	defer fmt.Println("b")
	fmt.Println("c")
}

var m=make(map[int]string)
var mu=sync.Mutex{}

func lookUp(key int) string{
	mu.Lock()
	defer mu.Unlock()
	return m[key]
}

func init(){
	fmt.Println("init")
	for i:=0;i<10;i++ {
		m[i]=strconv.Itoa(i)
	}
}