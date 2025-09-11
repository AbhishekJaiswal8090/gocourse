package main

import (
	"log"
	"os"
	"fmt"
)

func main(){
	log.Println("This is a log message")
	

	log.SetPrefix("INFO: ")
	log.Println("This is a info meassage")


	//log Flags
	log.SetFlags(log.Ldate | log.Ltime |log.Lshortfile)
	log.Println("This is a log message with only date") 

	infoLogger.Println("This is an info message :")
	warnLogger.Println("This is an warning message")
	errorLogger.Println("This is an error messaege")
    
	file,err :=os.OpenFile("logger.txt",os.O_CREATE|os.O_WRONLY,0755)
	if err !=nil{
		fmt.Println(err)
		return
	}

}
//custom logging function

var (
	infoLogger =log.New(os.Stdout,"info: ",log.Ldate | log.Ltime |log.Llongfile)
	warnLogger =log.New(os.Stdout,"WARN :",log.Ldate | log.Ltime |log.Lshortfile)
	errorLogger =log.New(os.Stdout,"ERROR :",log.Ldate |log.Ltime |log.Lmicroseconds)
)