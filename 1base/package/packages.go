package main

import (
	"math/rand"
	"fmt"
	"time"
)

func main(){
	fmt.Println("My favorite number is ",rand.Intn(100))
	fmt.Println("Now Time is ",time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("not nuix time is ",time.Unix(time.Now().Unix(),0).Format("2006-01-02 15:04:05"))

}