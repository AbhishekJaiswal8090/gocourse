package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var counter int32 = 0

	// CAS: only update if current value == expected
	success := atomic.CompareAndSwapInt32(&counter, 0, 10)

	fmt.Println("Swap successful?", success) // true
	fmt.Println("Counter:", counter)         // 10
}








// func main() {
// 	var counter int32 = 0
// 	var wg sync.WaitGroup

// 	// Spawn 100 goroutines, each incrementing the counter
// 	for i := 0; i < 100; i++ {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			// Atomically add 1
// 			atomic.AddInt32(&counter, 1)
// 		}()
// 	}

// 	wg.Wait()
// 	fmt.Println("Final Counter:", counter) // Should be 100
// }
