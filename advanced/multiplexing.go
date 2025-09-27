package main

import (
	"fmt"
	"time"
)

//multiplexing
//multiplexing is handling multiple  channel to send and receive the data


func main(){
	ch := make(chan int )

	go func(){
		ch <-1
		close(ch)
	}()

	for{
		select {
		case msg ,ok :=<-ch:
			if !ok{
				fmt.Println("Channel closed")
				//cleanup activities
				return
			}
			fmt.Println(msg)
		}
	}
}



// func main(){
// 	ch := make(chan int)
// 	go func(){
// 		time.Sleep(1*time.Second)
// 		ch <-1
// 		close(ch)
// 	}()

// 	select {
// 	case msg := <-ch:
// 		fmt.Println(msg)
// 	case <- time.After(1* time.Second):
// 		fmt.Println("Timeout")
// 	}
// }




// func main() {
// 	ch1 := make(chan int)
// 	ch2 := make(chan int)

// 	go func() {
// 		ch1 <- 1
// 	}()
// 	go func() {
// 		time.Sleep(time.Second)
// 		ch2 <- 1
// 	}()
// 	for i := 0; i < 2; i++ {
// 		select {
// 		case msg1 := <-ch1:
// 			fmt.Println("Received from ch1", msg1)

// 		case msg2 := <-ch2:
// 			fmt.Println("received from ch2", msg2)

// 			// default:
// 			// 	fmt.Println("No channel is ready")

// 		}

// 	}
// }
