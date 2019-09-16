package main

import (
	"net"
	"log"
	"io"
	"time"
	"fmt"
)

func handleConn(c net.Conn){
	defer c.Close()
	for{
		_,err:=io.WriteString(c,time.Now().String()+"\n")
		if err!=nil{
			fmt.Println(err)
		}
		time.Sleep(1*time.Second)
	}
}

func main(){
	listener, err:=net.Listen("tcp","localhost:7771")
	if err!=nil{
		log.Fatal(err)
	}
	for {
		conn,err:=listener.Accept()
		if err!=nil{
			log.Fatal(err)
		}
		go handleConn(conn)
	}
}
