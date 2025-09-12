package main

import (
	"log"
	"os"
	
)

func logging(){
	log.Println("This is a log message")
	

	log.SetPrefix("INFO: ")
	log.Println("This is a info meassage")


	//log Flags
	log.SetFlags(log.Ldate | log.Ltime |log.Lshortfile)
	log.Println("This is a log message with only date") 

	infoLogger.Println("This is an info message :")
	warnLogger.Println("This is an warning message")
	errorLogger.Println("This is an error messaege")
    
	file,err :=os.OpenFile("logger.txt",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	if err !=nil{
		log.Fatalf("Failed to opened a file")
	}

	infoLogger1 :=log.New(os.Stdout,"info: ",log.Ldate | log.Ltime |log.Llongfile)
	warnLogger1 :=log.New(os.Stdout,"WARN :",log.Ldate | log.Ltime |log.Lshortfile)
	errorLogger1 :=log.New(os.Stdout,"ERROR :",log.Ldate |log.Ltime |log.Lmicroseconds)

	infoLogger1.Println("This is a info message")
	warnLogger1.Println("This is a warning message")
	errorLogger1.Println("This is error message")

	defer file.Close()

	debugLogger :=log.New(file,"DEBUG",log.Ldate|log.Ldate|log.Llongfile)
	debugLogger.Println("This is a debug message")
	

}
//custom logging function

var (
	infoLogger =log.New(os.Stdout,"info: ",log.Ldate | log.Ltime |log.Llongfile)
	warnLogger =log.New(os.Stdout,"WARN :",log.Ldate | log.Ltime |log.Lshortfile)
	errorLogger =log.New(os.Stdout,"ERROR :",log.Ldate |log.Ltime |log.Lmicroseconds)
)