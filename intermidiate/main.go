package main

import "fmt"

// demosntrating closures

func main() {

	//sequence :=adder()
	//fmt.Println(sequence())
	//fmt.Println(sequence())
	//fmt.Println(sequence())

	//another closure demonsatrating anonymous function
/*	subTractor := func() func(int) int {
		countDown := 99
		return func(x int) int {
			countDown -= x
			return countDown
		}
	}()*/

	//	fmt.Println(subTractor(1))
	//   fmt.Println(subTractor(2))
	//  fmt.Println(subTractor(3))
	//  fmt.Println(subTractor(4))
	//  fmt.Println(subTractor(5))

	//recursion
	//fmt.Println(factorial(5))
	//fmt.Println(sumOfDigit(10))

	//pointers in go
  /*  var ptr *int
	a := 10
	ptr =&a
	fmt.Println(a,ptr)

	//modifying values of a 
	*ptr =20 //this is called derefrencing 
	fmt.Println(a)*/

	message :="hELLO GO"
	randMessage :="hrllo \ngo"
    
	fmt.Println(message)
	fmt.Println(randMessage)

}

//a function that demonstartes tha closures
func adder() func() int {
	i := 0
	fmt.Println("previous value of i", i)
	return func() int {
		i++
		fmt.Println("add 1 to i ", i)
		return i
	}

}

//recursion in go
func factorial(x int) int {
	//base case factorial of 0 is 1
	if x == 0 {
		return 1
	}
	//recursive case
	return x * factorial(x-1)

}

func sumOfDigit(x int)int{
	sum :=x
	if x==0{
		return sum
	}

	return sum + sumOfDigit(x -1)
}
