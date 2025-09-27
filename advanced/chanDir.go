package main 

import ("fmt"
        "time")  
//channel direction
//we can both send and receive the data but sometimes we wants to only send or recive the data
//thats where channel direction comes in

func main(){

	//send only channel chan <- int

	//receive only channel <- chan int

	ch := make(chan int) //it is bidirectional channel
	go sendData(ch)
	receiveOnly(ch)
}

//example of sending only channel direction
func sendData(ch chan<- int) {
	for i := 0; i < 3; i++ {
		ch <- i //here ch <- is sendonly channel if we try to <-ch we'll get an error
		fmt.Println("Sending... ")
	}
	close(ch)
}


//example of receiving only channel direction
func receiveOnly(ch <-chan int) {
	for val := range ch {
		fmt.Println("Recived ", val)
	}
}