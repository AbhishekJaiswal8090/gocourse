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


// func main(){
// 	ch :=make(chan int)

// 	go func(){
// 		fmt.Println("sending")
// 		ch <-9//blocking until the value is received
// 		fmt.Println("Sent value")
// 		time.Sleep(2*time.Second)
		
// 	}()
// 	value := <-ch
// 	fmt.Println("VALUE RECEIVES",value)

// 	time.Sleep(3*time.Second)

// }

func main(){
	numOfGoroutine:=3
	done :=make(chan int ,3)

	for i:= range numOfGoroutine{
		go func(id int){
          fmt.Printf("Goroutines %d working \n",id)
		  time.Sleep(time.Second)
		  done <- 1
		}(i)
	}
	for range numOfGoroutine{
		<-done //wait for each goroutines to finish
	}

	fmt.Println("Finished")
}

// Channel synchronization means using channels as a way for
//  goroutines to signal each other 
// and ensure proper ordering or completion.