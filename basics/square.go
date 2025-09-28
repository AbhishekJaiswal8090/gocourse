package main

import (
	"fmt"
	"time"
)

func main() {
	tasks := []int{2, 4, 6, 8, 10}
	results := make(chan int)

	for _, val := range tasks {
		go func(n int) {
			time.Sleep(time.Second)
			results <- n * n
		}(val)
	}

	for i := 0; i < len(tasks); i++ {
		fmt.Println("Results ", <-results)
	}
}

// func main(){
// 	task :=[]int{2,4,6,8,10}

// 	for i,val:=range task{
// 		fmt.Printf("Square of %d is %d \n",val,val*val)
// 	}
// }
