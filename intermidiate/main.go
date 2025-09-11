package main

import (
	"fmt"
	"os"
	"strings"
)
//env varibale are usulaly used for configuration
//special keys have their own constant values 
//so to avoid changing everywhere there value 
//we declare it at a single place accese it via theri key

func main(){
    user :=os.Getenv("USER")
	home :=os.Getenv("HOME")

	fmt.Println(user)
	fmt.Println(home)

	err:=os.Setenv("KEY","Value")
	if err !=nil{
		fmt.Println(err)
		return
	}

	fmt.Println("FRUIT env var is:",os.Getenv("KEY"))

	for _,e :=range os.Environ(){
		kvpair :=strings.SplitN(e,"=",2)
		fmt.Println(kvpair[0])

	}

	err=os.Unsetenv("KEY")
	if err !=nil{
		fmt.Println(err)
		return
	}


	str :="a=b=c=d=e"

	fmt.Println(strings.SplitN(str,"=",-1))
	fmt.Println(strings.SplitN(str,"=",0))
	fmt.Println(strings.SplitN(str,"=",1))
	fmt.Println(strings.SplitN(str,"=",2))
	fmt.Println(strings.SplitN(str,"=",3))
	fmt.Println(strings.SplitN(str,"=",4))
}