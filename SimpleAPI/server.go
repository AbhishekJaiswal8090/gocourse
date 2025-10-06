package main

import (
	"fmt"
	"net/http"
)

func main(){

	http.HandleFunc("/orders",func( w http.ResponseWriter,r *http.Request){
		fmt.Fprintf(w,"Handling incoming orders")
	})

	http.HandleFunc("/users",func(w http.ResponseWriter,r *http.Request){
		fmt.Fprintf(w,"Handling  users")
	})



	port:=3000;

	fmt.Println("Our server is working on: ",port)
	http.ListenAndServe(fmt.Sprintf(": %d",port),nil)
}