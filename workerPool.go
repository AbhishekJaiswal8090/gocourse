package main

import (
	"fmt"
	"time"
)

func main() {

	//worker pool -> are the concureency pattern when we have
	//worker -> goroutines to get some work done
	// jobs channel-> jobs channels are more like task that we want to perform using the channel
	//result channel -> result channels that are used to receive the result after completing job/task

	numberOfWorker := 3
	numberOfJob := 10
	tasks := make(chan int, numberOfJob)
	results := make(chan int, numberOfJob)

	//craete worker
	for i := range numberOfWorker {
		go worker(i, tasks, results)
	}

	// sebd values to task channel
	for i := range numberOfJob {
		tasks <- i
	}

	close(tasks)

	for range numberOfJob {
		result := <-results
		fmt.Println(result)
	}
}

func worker(id int, tasks <-chan int, results chan<- int) {
	for task := range tasks {
		fmt.Printf("Worker %d processing task %d \n", id, task)
		//simulate task
		time.Sleep(time.Second)
		results <- task * 2
	}

}
