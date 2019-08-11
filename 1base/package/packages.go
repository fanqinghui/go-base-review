package main

import (
	"math/rand"
	"fmt"
	"time"
)

func main(){
	fmt.Println("My favorite number is ",rand.Intn(10))
	fmt.Printf("Now Time is ",time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("not nuix time is ",time.Now().Unix())
}