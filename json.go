package main

import (
	"encoding/json"
	"fmt"
	"log"
)


func main() {
	address:=Address{
		City : "Kushinagar",
		State : "UP",
	}

	person := Person{
		First_Name: "Abhishek",
		Age: 19,
	    Email: "sample@gmail.com",
		Address: address,
	}
	//marshalling struct or encoding struct into json
    data,err:=	json.Marshal(person)
	if err!=nil{
		fmt.Println("Error marshlling the json",err)
		return
	}

	fmt.Println(string(data))


	//unmarsaling json or decoding json into struct
	var user Person

	err = json.Unmarshal([]byte(data), &user)
	if err != nil {
		fmt.Println("Error :decoding json", err)
		return
	}
    fmt.Println("Decoding json and extracting information back into the norml string")
	fmt.Println(user)

	jsonData :=`{"Fullname":"Abhishek","City":"kshinagar","State":"up"}`
    
	//creating struct to receive the data back from the json
	type Employee struct{
		Fullname string `json :"fullname"`
		City string `json :"city"`
		State string `json:"state"`
	}
	var Emp Employee
	err =json.Unmarshal([]byte(jsonData),&Emp)
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println(Emp.Fullname,Emp.City,Emp.State)

//encoding list into json
	listOfcityState :=[]Address{
		{
			City: "newyork",
			State: "London",
		},
		{
			City: "Wahington",
			State: "USA",
		},
		{
			City: "Delhi",
			State: "India",
		},
	}

	fmt.Println(listOfcityState)
	data2,err:=json.Marshal(listOfcityState)
	if err!=nil{
		fmt.Println("Error :Encoding listofCityState",err)
		return
	}
	fmt.Println("Json list: ",string(data2))


	//handling unknown json structure

	jsonData2:=`{"name":"John","age":20,"addres":"NewYork","state":"NY"}`
     
   var data3 map[string]interface{}
   err=json.Unmarshal([]byte(jsonData2),&data3)
   if err!=nil{
	log.Fatal("Error :unmarshalling the data")
	return
   }

   fmt.Println("Decode data: ",data3)
    fmt.Println("Decode Value of name: ",data3["name"])
	
}

type Person struct {
	First_Name string `json:"first_name"`
	Age        int    `json:"age, omitempty"`//omitempty removes the spaces
	Email string      `json:"email"`
	Address Address   `json :"address"`
}
type Address struct{
	City string `json:"city"`
	State string `json: "state"`
}
