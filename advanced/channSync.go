package main

import (
	"fmt"
	"time"
)

//channel synchronization

// func main(){
// 	done :=make(chan struct{})

// 	go func(){
// 		fmt.Println("Working...")
// 		time.Sleep(2*time.Second)
// 		done <-struct{}{}
// 	}()

// 	<-done //this is called channel synchronization
// 	//here the main thread waits untill it receives the value we sent 
// 	fmt.Println("finished")


//  }


func main(){
	ch :=make(chan int)

	go func(){
		fmt.Println("sending")
		ch <-9//blocking until the value is received
		fmt.Println("Sent value")
		time.Sleep(2*time.Second)
		
	}()
	value := <-ch
	fmt.Println("VALUE RECEIVES",value)

	time.Sleep(3*time.Second)

}