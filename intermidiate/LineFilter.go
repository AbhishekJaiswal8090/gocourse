package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LineFilter(){

	fmt.Println("Line Filtering")
    file,err :=os.Open("example.txt")
	if err!=nil{
		fmt.Println("Error : opening file")
		return
	}
	defer file.Close()

	scanner :=bufio.NewScanner(file)
	
	// keyword to filter line
	keyword:="important"

	//read and filter line
	lineNumber := 0

	for scanner.Scan(){
		line:=scanner.Text()
		if strings.Contains(line ,keyword) {
 
			// replacing that particular word from that line to some else word
			udpatedLine:=strings.ReplaceAll(line,keyword,"necessary")
 
			//filter the line
			  fmt.Printf("%d filtered line %v\n",lineNumber,line)
			  lineNumber++
			  fmt.Printf("%d updated line: %v\n",lineNumber, udpatedLine)

			 lineNumber++
		}
	}
	err =scanner.Err()
	if err!=nil{
		fmt.Println(err)
		return
	}



}