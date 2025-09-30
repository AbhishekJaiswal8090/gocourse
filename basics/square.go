package 

import (
	"fmt"
	"time"
 )



 

//  func Worker(id int , jobs <-chan int ,results chan int){
// 	for j:= range jobs{
// 		fmt.Printf("Worker %d started job %d\n",id,j)
// 		time.Sleep(time.Second)
// 		//simulate work
// 		results  <- j * j
// 		fmt.Printf("worker %d finished job %d \n",id,j)
// 	}


//  }

//  func main(){
// 	tasks := []int{2,4,6,8,15}
// 	jobs := make(chan int,len(tasks))
// 	results :=make(chan int,len(tasks))

// 	//start three worker only

// 	for i:=0;  i<3; i++{
// 		go Worker(i,jobs,results)
// 	}

// 	//send jobs
// 	for _,t:=range tasks{
// 		jobs <- t
// 	}

// 	close(jobs)

// 	//collect results 
// 	for i:=0; i<len(tasks); i++{
// 		fmt.Println("Result ",<-results)
// 	}
//  }







// func main() {
// 	tasks := []int{2, 4, 6, 8, 10}
// 	results := make(chan int)

// 	for _, val := range tasks {
// 		go func(n int) {
// 			time.Sleep(time.Second)
// 			results <- n * n
// 		}(val)
// 	}

// 	for i := 0; i < len(tasks); i++ {
// 		fmt.Println("Results ", <-results)
// 	}
// }

// func main(){
// 	task :=[]int{2,4,6,8,10}

// 	for i,val:=range task{
// 		fmt.Printf("Square of %d is %d \n",val,val*val)
// 	}
// }
