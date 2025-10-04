package main

import (
	"fmt"
	"io"
	"net/http"
)

func client(){

	// create a new client
	client :=&http.Client{}

	// res ,err:=client.Get("https://jsonplaceholder.typicode.com/posts/1")

	res,err:=client.Get("https://swapi.dev/api/people/1")
	if(err!=nil){
		fmt.Println("Error making request: ",err);
		return
	}
	defer res.Body.Close()

	// Read and print the response body

	body,err:=io.ReadAll(res.Body)
	if(err!=nil){
		fmt.Println(err);
		return
	}
	fmt.Println(body)
	fmt.Println(string(body))


}