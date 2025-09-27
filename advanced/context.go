package main

import (
	"context"
	"fmt"
	"time"
)

func main(){



	// todoContext :=context.TODO()
	// contextBkg:=context.Background()


	// ctx:=context.WithValue(todoContext,"name","Abhi")
	// fmt.Println(ctx)
	// fmt.Println(ctx.Value("name"))


	// ctx1:=context.WithValue(contextBkg,"city","LKO")
	// fmt.Println(ctx1)
	// fmt.Println(ctx1.Value("city"))




	// ctx :=context.TODO()
	// result :=checkEvenOdd(ctx ,5)
	// fmt.Println(result)


	ctx ,cancel:=context.WithTimeout(context.Background(),2* time.Second)

	// defer cancel()


    // result =checkEvenOdd(ctx ,5)
	// fmt.Println("Result from timeoout context",result)


	//handling ticker function
	// ctx ,cancel= context.WithCancel(context.Background())

	// go tickerFunc(ctx)

	// time.Sleep(5 *time.Second)
	// cancel()

	// time.Sleep(1*time.Second)
	// fmt.Println("Main function ended ")


	//handling dowork

	rootctx:=context.Background()
    ctx , cancel =	context.WithTimeout(rootctx,2*time.Second)
	defer cancel()
	ctx =context.WithValue(ctx,"requestID","abi1234")

	 go dowork(ctx)
    time.Sleep(3*time.Second)

	requestID :=ctx.Value("requestID")
	if requestID!=nil{
		fmt.Println("Request ID is",requestID)
	}else{
		fmt.Println("Not found")
	}
}

func checkEvenOdd(ctx context.Context,num int)string{
	select{
	case <-ctx.Done():
		return "Operation Cancel"
	default:
		if num %2 == 0{
			return fmt.Sprintf("%d is even",num)
		}else{
			return fmt.Sprintf("%d is odd",num)
		}
	}

}

func tickerFunc(ctx context.Context){
	ticker :=time.NewTicker(1*time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Ticker stops")
			return
		case t:= <-ticker.C:
			fmt.Println("Ticker tick at",t)
			// Handle ticker tick here if needed
		}
	}
}

func dowork(ctx context.Context){
	for{
		select{
		case <- ctx.Done():
			fmt.Println("works canceled",ctx.Err())
			return
		default:
			fmt.Println("working ...")
		}
		time.Sleep(500*time.Millisecond)
	}
}