package main

import (
	"fmt"
	"time"
)

//channels in go



func main(){


	//creation of channels 
	greeting := make(chan string)
	greetString :="Hello string"


	//go routines are the very necessary key
	//to use channels without using goroutines the program get into the deadlock condn
	//
	go func(){
		greeting <- greetString
		for _,val :=range "abcde"{
			greeting <- string(val)
		}

	}()
	go func(){
	receiver := <- greeting
	for i:=0; i<5; i++{
		recvr:=<-greeting
		fmt.Println(recvr) 
	}
	fmt.Println(receiver)
    }()
	
    fmt.Println("Here is the end of the program")
	time.Sleep(5*time.Second)


}