package main

import (
	"fmt"
	"sync"
	"time"
)

// waitgroup
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	//simulate some task for some time
	fmt.Printf("Worker %d finished ", id)

}

func main() {
	var wg sync.WaitGroup
	numOfworkerr := 3

	wg.Add(numOfworkerr)

	//launch worker
	for i := range numOfworkerr {
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("All workers finished")
}
