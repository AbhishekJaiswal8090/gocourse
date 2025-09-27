package main           

import "fmt"


//range over closed channel
func main(){
	ch := make(chan int)
	go func(){
		for i:= range 5{
          ch <-i
		}
		close(ch)
	}()

	for val :=range ch{
		fmt.Println("Received",val)
	}
}






// func main(){
// 	ch := make(chan int)
// 	close(ch)

// 	val ,ok:=<-ch
// 	if !ok{
// 		fmt.Println("Channel is closed")
// 	}
// 	fmt.Println(val)
// }





//simple closing channel example
// func main(){
// 	ch :=make(chan int)

// 	go func(){
// 		for i:= range 5{
// 			ch <-i
// 		}
// 		close(ch)
// 	}()
// 	for val :=range ch{
// 		fmt.Println(val)
// 	}
// }