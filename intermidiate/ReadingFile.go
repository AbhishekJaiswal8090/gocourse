package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadingFile() {

	file, err := os.Open("Output.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		fmt.Println("closing open file")
		defer file.Close()

	}()

	fmt.Println("file opened successfully")
	// /// reading content of the file
	// data := make([]byte, 1024)
	// _, err = file.Read(data)

	// if err != nil {
	// 	fmt.Println("error reading file", err)
	// 	return
	// }

	// fmt.Println("file content", string(data))
   
	scanner :=bufio.NewScanner(file)

	//read file line by line
	for scanner.Scan(){
		line :=scanner.Text()
		fmt.Println(line)
	}

	err =scanner.Err()
	if err!=nil{
		fmt.Println("Error reading file",err)
		return 
	}
	

    
	
}
