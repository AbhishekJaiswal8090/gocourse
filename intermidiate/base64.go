package main

//base 64 coding
import (
	
	"encoding/base64"
	"fmt"
)
 func base64Encoding(){
data :=[]byte("He~lo , base64 coding") 
//encode base64

encoded :=base64.StdEncoding.EncodeToString(data)
fmt.Println(encoded)

//decode frome base64
decoded,err :=base64.StdEncoding.DecodeString(encoded)
if err !=nil{
	fmt.Println(err)
}
fmt.Println(string(decoded))
fmt.Println(len(decoded),len(data))

//URL safe ,avoid '/' and '+'

 }