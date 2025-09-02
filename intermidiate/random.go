package main

import (
	"fmt"
	"math/rand"
)

func random(){
	
for {
	fmt.Println("welcome to the dice game")
	fmt.Println("1. roll the dice")
	fmt.Println("2. Exit")
	fmt.Println("Enter your choice")
    
	var choice int
	_,err :=fmt.Scan(&choice)

	if err !=nil ||(choice != 1 && choice !=2){
		continue
	}
	if choice == 2{
		fmt.Println("quitting the game ")
		break
	}
	die1 :=rand.Intn(6) +1
	die2 :=rand.Intn(6)+1

	fmt.Printf("you rolled a %d and a %d \n",die1,die2)
	fmt.Println("Total is ",die1+die2)

	fmt.Println("do you want to roll again (y/n)")

	var rollAgain string
	_ ,err =fmt.Scan(&rollAgain)
	if err !=nil || (rollAgain !="y" && rollAgain !="n"){
       fmt.Println("Invalid input ,assuming no.")
	   break
	}
	if rollAgain =="n"{
		fmt.Println("thanks for playing good bye")
		break
	}

}
}