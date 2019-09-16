package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func square(ch chan int,x int)  {
	defer wg.Done()
	ch<- x*x
	//close(ch)
}


func main()  {
	ch:=make(chan int,2)
	//defer close(ch)

	/*go func(){
		for i:=0;i<10;i++{
			ch<-i*i
		}
		//close(ch)//必须关闭ch
	}()*/
	arr:=[2]int{1,2}
	for _,num:=range arr {
		wg.Add(1)
		go square(ch,num)
		//close(ch)
	}

	wg.Wait()
	close(ch)

	for val:=range ch {
		fmt.Println(val)
	}

}
