package main

import (
	"encoding/xml"
	"log"
	"fmt"
)

type Person struct {
	XMLName xml.Name  `xml:"person"`
	Name string       `xml:"name"`
	Age int           `xml:"age"`
	City string       `xml:"city"`
	Email string      `xml:"email"`
}

func main() {

	person :=Person{
		Name: "Abhishek",
		Age:19,
		City:"Kushinagar",
		Email:"abhi84514aaa.in@gmail.com",
	}
	xmldata ,err:=xml.Marshal(person)
	if err!=nil{
		log.Fatal("Error marshalling the data,",err)
		return
	}
	fmt.Println(string(xmldata))


}