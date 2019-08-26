package main

import "fmt"
//常用的复合型数据结构
//array   数组
// slice  列表
// map    键值对
func main(){
	name2Age:=map[string]int{}
	//name2Age:=make(map[string]int)
	name2Age["张三"]=12
	name2Age["历史"]=14
	fmt.Println(name2Age)

	//initMap
	name2Age2:=map[string]int{
		"张三":14,
		"历史":15,
	}
	fmt.Println(name2Age2)

}


func slice(){
	list:=[]int{1,2}
	fmt.Println(list,len(list))
	list=append(list, 3,4,5)
	println(len(list),cap(list))
	list=append(list,6,7)
	println(cap(list))
	list=append(list,7,8,9,10,11)
	println(cap(list))
}

func arr(){
	arr:=[3]int{1,2,3}
	fmt.Println(arr)
	arr2:=[3]int{1,2}
	fmt.Println(arr2)
	arr3:=[3]int{1:1,2:2}
	fmt.Println(arr3)
	fmt.Println(len(arr))
}
