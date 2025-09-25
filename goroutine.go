package main 

import (
	"fmt"
	"time"
)

//goroutine are the tiny threads connected to  main thread of the program
//once a gouroutine are assigned they leave the main thread 
//perform their task and then catch the main thread 
//if the main thread remains into the program execution


//goroutines are defined using go keyword before the any function

func main(){

fmt.Println("Calling goroutine function")
 go saYhello()
 time.Sleep(2*time.Second)
 fmt.Println("after sayhello function")

 go PrintNumber()
 go PrintString()

 time.Sleep(5*time.Second)
}
func saYhello(){

	time.Sleep(1 *time.Second)
	fmt.Println("Hello from goroutine")
}

func PrintNumber(){
	for i:=0; i<=5; i++{
		fmt.Println(i)
		time.Sleep(100*time.Microsecond)
	}
}

func PrintString(){
	for _,letter :=range "abcde"{
		fmt.Println(string(letter))
		time.Sleep(200*time.Millisecond)
	}
}