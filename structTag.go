package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Firstname string `json:first_name`
	Lastname  string `json:last_name,omitempty`
	Age       int    `json:age,omitempty`
}

func main() {

	person := Person{
		Firstname: "Jane",
		Lastname:  "Doe",
		Age: 45,
	}
	jsonData, err := json.Marshal(person)
	if err!=nil{
		log.Fatal("Error Mrashaling struct",err)
		return
	}
	fmt.Println(string(jsonData))
}