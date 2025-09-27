package main

import (
	"fmt"
	"time"
)

func main(){
	ticker :=time.NewTicker(time.Second)
	stop :=time.After(5*time.Second)
	defer ticker.Stop()

	for{
		select{
		case tick := <-ticker.C:
			fmt.Println("Tick at", tick)
		case <-stop:
			fmt.Println("Stopping ticker")
			return
		}
	}
}




// func main() {
// 	ticker := time.NewTicker(time.Second)
// 	defer ticker.Stop()
// 	// for tick := range ticker.C {
// 	// 	fmt.Println("Tick at", tick)
// 	// }
// 	i := 0
// 	for range 5 {
// 		i++
// 		fmt.Println(i)
// 	}

// 	//handling tickerFunc
// 	TickerFunc()
// }

// func PeriodicTaks() {
// 	fmt.Println("Performing Task at:", time.Now())
// }

// func TickerFunc() {
// 	ticker := time.NewTicker(time.Second)
// 	defer ticker.Stop()

// 	for range ticker.C {
// 		PeriodicTaks()
// 	}
// }
