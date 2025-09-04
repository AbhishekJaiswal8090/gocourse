package main

import (
	"fmt"
	"strconv"
)

func Numparsing(){
	numstr :="12345"
	num ,err :=strconv.Atoi(numstr)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println("parsed integer",num +1)

	num1 ,err :=strconv.ParseInt(numstr ,10,32)
	if err!= nil{
		fmt.Println(err)
	}
	fmt.Println(num1)

	floatstr :="34.78"
	floatnum ,err :=strconv.ParseFloat(floatstr,64)
	if err!= nil{
		fmt.Println(err)
	}
	fmt.Println(floatnum)

	binarystr :="1010"
	decimal ,err :=strconv.ParseInt(binarystr,2,64)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println("parsed binary to decimal val",decimal)

	hexstr :="ff"
	decnum ,err:=strconv.ParseInt(hexstr,16,64)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(decnum)
}