package main

import "fmt"

func main() {

	//int into int32 and float
	var a int =32
	b :=int32(a)
	c:=float64(b)
	// d:=bool(true)

	e:=3.14
	f:=int(e)
	fmt.Println(f,c)


	//string into byte
	g:="hello 0 Arigato ghojaimus"
	h:=[]byte(g)
	fmt.Println(h)


	
}
