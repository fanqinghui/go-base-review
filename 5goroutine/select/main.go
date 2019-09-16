package main

import (
	"fmt"
	"time"
)

func foo(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "success"
}

func main() {
	ch := make(chan string)
	go foo(ch)
	select {
	case val := <-ch:
		fmt.Println(val)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout!!!")
	}

}
