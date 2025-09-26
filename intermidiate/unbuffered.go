package main

import (
	"fmt"
	"time"
)

func main(){
	//unbuffered channels can't hold value they don't have storgae
	//and the unbuffered channel needs an immediatedly receiver after while sending the data otherwise it deadlocks
	//it send the data even when block is not full unlike buffered channel

	 ch :=make(chan int)
    // ch <-2
	// recvr := <-ch
	//  fmt.Println(recvr)
	//here we get an error because of their is no immediate receiver 
	//an receiver could be a goroutine as we said the unbuffered channel needs an immediate recvr
	//since there is singlr main thread there are no other thats why it get deadlock

    go func(){
		ch <- 40

	}()
	go func(){
		recvr := <- ch
		fmt.Println(recvr)
	}()

	time.Sleep(2*time.Second)
	fmt.Println("END OF PROGRAM")
}