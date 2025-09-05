package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func bufIo(){

// 	  reader :=bufio.NewReader(strings.NewReader("hello ? bufio package fvk u\n"))
//   data := make([]byte,40)
  
//   n,err:=  reader.Read(data)
//   if err !=nil{
// 	fmt.Println(err)
// 	return 
//   }
//   fmt.Printf("read %d bytes %s",n,data[:n])

//   line ,err:=reader.ReadString('\n')
//   if  err != nil{
// 	fmt.Println(err)
// 	return
//   }
//   fmt.Println("read string",line)
writer :=bufio.NewWriter(os.Stdout)
// writting bytes slice first
data :=[]byte("Hello ,bufio package \n")
 n ,err :=writer.Write(data)
 if err !=nil{
	fmt.Println(err)
	return
 }
 fmt.Printf("we wrote %d bytes \n",n)
 //flush the buffer to ensure all the data is written to os.stdout
 err = writer.Flush()
 if err !=nil{
	fmt.Println(err)
 }

 Str :="this is a new string"
  n ,err =writer.WriteString(Str)
  if err!=nil{
	fmt.Println(err)
	return
  }
  fmt.Printf("we wrote %d bytes",n)
  err = writer.Flush()
  fmt.Println(err)

  

//Reading
	text :=strings.NewReader("Hii brotha i am here \n to study")
	reader :=bufio.NewReader(text)
	for {
		line ,err :=reader.ReadString('\n')
		if err != nil{
			fmt.Println(err)
			break
		}
		fmt.Printf("we read  %s",line)
	}
	
	// Writing
	writerr := bufio.NewWriter(os.Stdout)
	

	daataa :=[]byte("hello from buffer iio \n")


	m,err:=writerr.Write(daataa)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(m)
    
}