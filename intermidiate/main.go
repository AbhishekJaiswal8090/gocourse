package main

import (
	"fmt"
	"os"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// directories
func main() {
	// err :=os.Mkdir("subdir",0755)
	// checkError(err)

	//creating subdir
	// checkError(os.Mkdir("subdir", 0755))


	//removing subdir
	// os.RemoveAll("subdir")

	// os.WriteFile("subdir",[]byte("Hello brotha"),0755)

	// checkError(os.MkdirAll("subdir1/parent/child1",0755))
	// checkError(os.MkdirAll("subdir1/parent/child2",0755))
	// checkError(os.MkdirAll("subdir1/parent/child3",0755))
    
	os.WriteFile("subdir/parent/file",[]byte(""),0755)

	res ,err:=os.ReadDir("subdir/parent")
    checkError(err)

	for _,val:=range res{
		fmt.Println(val)
	}
}
