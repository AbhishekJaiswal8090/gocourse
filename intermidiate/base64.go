package main

//base 64 coding
import (
	
	"encoding/base64"
	"fmt"
)
 func base64Encoding(){
	//encoding  and decoding data using base64

	//encoding
	data :=[]byte("Here is the data")
    encodedData:=base64.StdEncoding.EncodeToString(data)

	fmt.Println(encodedData)

	//decoding
	data2 :=[]byte("123 IS THE DATA THAT WE WANT TO ENCODE AND THEN DECODE")
	encodedData2:=base64.StdEncoding.EncodeToString(data2)

	decodeData2,err:=base64.StdEncoding.DecodeString(encodedData2)
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println(string(decodeData2))
    fmt.Println(len(decodeData2))
//URL safe ,avoid '/' and '+'

 }