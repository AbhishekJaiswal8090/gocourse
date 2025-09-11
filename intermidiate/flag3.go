package main

import (
	"flag"
	"fmt"
	"os"
)


func flag3(){
	subcommand1 :=flag.NewFlagSet("firstSub",flag.ExitOnError)
	subcommand2:=flag.NewFlagSet("secondSUb",flag.ExitOnError)

	firstflag :=subcommand1.Bool("processing",false,"command processing status")
	secondflag :=subcommand1.Int("bytes",1024,"byte length of result")

	flagsc2 :=subcommand2.String("language","Go","enter your lang")

	if len(os.Args)<2{
       fmt.Println("This process require additional commands")
	   os.Exit(1)
	}

	switch os.Args[1]{
	case "firstSub":
		subcommand1.Parse(os.Args[2:])
		fmt.Println("Subcommand:")
		fmt.Println("processing",*firstflag)
     	fmt.Println(" bytes",*secondflag)
    case "secondSub":
		subcommand2.Parse(os.Args[2:])
		fmt.Println("SUbcommand2",)
		fmt.Println("language: ",*flagsc2)

	default:
		fmt.Println("no subcommand entered")
		os.Exit(1)	
	}

	
    
}