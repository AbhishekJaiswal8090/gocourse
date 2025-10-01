package main

import (
	"fmt"
	"sync"
)

//=> mutexes in go

// why use Mutexes
// 1.Data INtegrity
// 2. Consistency
// 3.Safety

type Counter struct{
	mu sync.Mutex
	count int
}

func (c *Counter) Increament(){

	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}
func( c *Counter ) getValue()int{
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}


func main(){
  var wg sync.WaitGroup
  counter :=&Counter{}

  numOfgoroutines :=10

  wg.Add(numOfgoroutines)

  for range numOfgoroutines{
	go func(){
      defer wg.Done()
	  for range 1000 {
		counter.Increament()
	  }
	}()
  }

  wg.Wait()
  fmt.Printf("Thr final counter calue is %d \n",counter.getValue())
}



// var counter int

// func increment() {
// 	for i := 0; i < 1000; i++ {
// 		counter++ // race condition: multiple goroutines writing here
// 	}
// }

// func main() {
// 	for i := 0; i < 3; i++ {
// 		go increment()
// 	}

// 	time.Sleep(1 * time.Second) // wait for goroutines
// 	fmt.Println("Final Counter:", counter)
// }
