package main

import (
	"flag"
	"fmt"
	
)

func flag2() {

	// fmt.Println("Command", os.Args[0])
	
	
	// for i, arg := range os.Args {
	// 	fmt.Println("Argument", i, arg)
	// }

	//define flags
	var name string
	var age int
	var male bool

	flag.StringVar(&name, "name", "Abhi", "Name of the user")
	flag.IntVar(&age, "age", 18, "Age of the user")
	flag.BoolVar(&male, "male", true, "gender of the user")

	flag.Parse()
    

	fmt.Println("NAME:", name)
	fmt.Println("AGE:", age)
	fmt.Println("MALE", male)

}
