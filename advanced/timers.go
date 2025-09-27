package main

import (
	"fmt"
	"time"
)

func main() {
	//timers in go
	fmt.Println("Starting the app ...")
	timer := time.NewTimer(2 * time.Second)
	
	fmt.Println("waiting of timer.C")
	<-timer.C//this is blocking in nature

	fmt.Println("Timer expired")

}
