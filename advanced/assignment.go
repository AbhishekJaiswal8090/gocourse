package main

import (
	"fmt"
	"time"
)


func Worker (id int,jobs <-chan string ,results <- chan string){
	for job :=range jobs{
		fmt.Printf("Worker %d started the job %d",id,job)
		time.Sleep(time.Second)
		//simulate the work 
		//process the url and get the job done 
		//downloading the webpages  and store into any particular format
		//while i am here considering string 
		results <- "Hello brotha how's my code "
	}
}

func main(){
	listOfurl :=[]string {"url1","url2","url3","url4","url5"}
	jobs :=make(chan string,len(listOfurl))
	results :=make(chan string)

	//set the worker to get the job done

	for w:=0; w<3; w++{
		go Worker(w, jobs ,results)
	}

	//now sent the each jobs to the channel
	for _,j :=range listOfurl{
		jobs <- j
	}

	for a:=0; a<len(listOfurl); a++{
	fmt.Println("Here is the result",<-results)
	}
}