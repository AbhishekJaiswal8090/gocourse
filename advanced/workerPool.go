package workerpool

import (
	"fmt"
	"time"
)

type ticketRequest struct {
	personID     int
	numofTickets int
	cost         int
}

func ticketProcessor(request <-chan ticketRequest, results chan<- int) {
	for req := range request {
		fmt.Printf("Processing %d tickets(s) of personID %d with total cost %d \n", req.numofTickets, req.personID, req.cost)

		//simulate processing time
		time.Sleep(time.Second)
		results <- req.personID
	}

}
func workerPool() {
	numRequest := 5
	price := 5
	ticketRequests := make(chan ticketRequest, numRequest)
	ticketResult := make(chan int)

	// start ticket processor
	for i := 0; i < 3; i++ {
		go ticketProcessor(ticketRequests, ticketResult)
	}

	//send ticket request
	for i := 0; i < numRequest; i++ {
		ticketRequests <- ticketRequest{personID: i + 1, numofTickets: (i + 1) * 2, cost: (i + 1) * price}
	}
	close(ticketRequests)

	for range numRequest {
		fmt.Printf("Tickets for personID %d processed successfuly\n", <-ticketResult)
	}
}

// func main() {

// 	//worker pool -> are the concureency pattern when we have
// 	//worker -> goroutines to get some work done
// 	// jobs channel-> jobs channels are more like task that we want to perform using the channel
// 	//result channel -> result channels that are used to receive the result after completing job/task

// 	numberOfWorker := 3
// 	numberOfJob := 10
// 	tasks := make(chan int, numberOfJob)
// 	results := make(chan int, numberOfJob)

// 	//craete worker
// 	for i := range numberOfWorker {
// 		go worker(i, tasks, results)
// 	}

// 	// sebd values to task channel
// 	for i := range numberOfJob {
// 		tasks <- i
// 	}

// 	close(tasks)

// 	for range numberOfJob {
// 		result := <-results
// 		fmt.Println(result)
// 	}

// 	time.Sleep(3*time.Second)
// }

// func worker(id int, tasks <-chan int, results chan<- int) {
// 	for task := range tasks {
// 		fmt.Printf("Worker %d processing task %d \n", id, task)
// 		//simulate task
// 		time.Sleep(time.Second)
// 		results <- task * 2
// 	}

// }
