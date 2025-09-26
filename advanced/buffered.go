package main

import (
	"fmt"
	"time"
)

func main(){
	//buffered channels 
	//since the name itself is bufferd it means storgae 
	//buffered channels have storage 
	//widely used for async communication

    // 	Send blocks only when buffer is full.
    // Receive blocks only when buffer is empty.
    fmt.Println("Bufferd channel")
	ch :=make(chan int ,2)
	ch <-1
	ch <-2
	fmt.Println("sent 1 and 2")

	// go func(){
	// 	time.Sleep(2*time.Second)
	// }()
    
	//if we send value more than its capacity the prgram gets into deadlock

	fmt.Println("receive ",<-ch)
	fmt.Println("receive ",<-ch)


	//and if we do want to use another value to be sent from buffered channel here's how it be done

	//since the chan have capcity 2 and the value are send and receives by the goroutines 
	//now we cn send another two value to this same channel

	ch <- 70
	ch <- 80
	fmt.Println("send 70 and 80")

	fmt.Println("receive",<-ch)
	fmt.Println("receive",<-ch)

	time.Sleep(2*time.Second)

}