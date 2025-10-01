package main

import (
	"fmt"
	"sync"
	"time"
)


func main(){

	var wg sync.WaitGroup

	numberOfWorker:=3
	numberOfJob:=3
	results := make(chan int,numberOfJob)
	wg.Add(numberOfJob)

	for i:= range numberOfWorker{
		go worker(i ,results ,&wg)
	}


	go func(){
		wg.Wait()
		close(results)
	}()

	for result :=range results {
		fmt.Println("Results: ",result)
	}

}


func worker(id int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker id %d starting \n", id)
	time.Sleep(time.Second)
	results <- id * 2
	fmt.Printf("Worker id %d finished", id)
}

// // waitgroup
// => so basically a waitgroup is a kind of synchronization 
//mechanism from the sync package that helps you
//wait for multiple  goroutines to finish before moving 
// on


// func worker(id int, wg *sync.WaitGroup) {
// 	defer wg.Done()//is basically reduces each time a counter 
// 	fmt.Printf("Worker %d starting\n", id)
// 	time.Sleep(time.Second)
// 	//simulate some task for some time
// 	fmt.Printf("Worker %d finished ", id)

// }

// func main() {
// 	var wg sync.WaitGroup
// 	numOfworkerr := 3

// 	wg.Add(numOfworkerr) //it works as the adding counters for that particular number of goroutines 


// 	//launch worker
// 	for i := range numOfworkerr {
// 		go worker(i, &wg)
// 	}

// 	wg.Wait()//blocking mechanism
// 	fmt.Println("All workers finished")
// }
