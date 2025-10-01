package main

import (
	"fmt"
	"time"
)
//=> mutexes in go

// why use Mutexes 
// 1.Data INtegrity
// 2. Consistency
// 3.Safety


var counter int

func increment() {
	for i := 0; i < 1000; i++ {
		counter++ // race condition: multiple goroutines writing here
	}
}

func main() {
	for i := 0; i < 3; i++ {
		go increment()
	}

	time.Sleep(1 * time.Second) // wait for goroutines
	fmt.Println("Final Counter:", counter)
}
