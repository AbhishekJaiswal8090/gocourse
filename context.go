package main 

import ("fmt"
"context")

func main(){

	todoContext :=context.TODO()
	contextBkg:=context.Background()


	ctx:=context.WithValue(todoContext,"name","Abhi")
	fmt.Println(ctx)
	fmt.Println(ctx.Value("name"))


	ctx1:=context.WithValue(contextBkg,"city","LKO")
	fmt.Println(ctx1)
	fmt.Println(ctx1.Value("city"))


}