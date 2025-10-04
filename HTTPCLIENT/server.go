package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "hello server ")
	})

	fmt.Println("Server is listening on port 3000")
	const port string = ":3000"
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("Error starting server", err)
		return
	}
}
