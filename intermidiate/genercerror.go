package main

import (
	"errors"
	"fmt"
	"math"
)

func genercError() {
	/*fmt.Println("Hello!")
		//starting generics in go
		hii,bro :=genercs("abhi","fuck")
		fmt.Println(hii,bro)
		hello("abhi","brotha","hello","you","hell")

		b,a := swap("abhi","akash")
		fmt.Println(b,a)

		//implementing stack
		s := Stack[int]{
			elements: []int{},
		}

		s.push(45)
	    s.push(50)
		s.printElement()
		s.pop()
		s.pop()
		s.printElement()
	*/

	//error in go

	//res, err := sqrt(-56)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
//	fmt.Println(res)

//	data := []byte{}
//	if err := process(data); err != nil {
//		fmt.Println("error", err)
//		return
//	}
//	fmt.Println(data)

	//if err1 := eprocess(); err1 != nil {
	//	fmt.Println("error", err1)
	//	return
	//}
 if err3 :=readData();err3!=nil{
	fmt.Println(err3)
	return
 }
 fmt.Println("read data successfully")
}


func genercs[T any](a, b T) (T, T) {
	return a, b
}
func hello[T any](a ...T) bool {
	for _, val := range a {
		fmt.Println(val)
	}
	return true
}
func swap[T any](a, b T) (T, T) {
	return b, a
}

// implementing stack in go
type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) push(element T) bool {
	s.elements = append(s.elements, element)
	fmt.Println("elements is pushed")
	return true
}

func (s *Stack[T]) pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		fmt.Println("underflow condition")
		return zero, false
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	fmt.Println("element is popped")
	return element, true
}

func (s *Stack[T]) printElement() {
	fmt.Println(s.elements)
}

func (s *Stack[T]) isOkay() string {
	if len(s.elements) == 0 {
		return "none elements are pushed "
	}
	fmt.Println(s.elements)
	return "there are all element"
}

// implementing error
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("math : square root of negative number")
	}
	return math.Sqrt(x), nil
}

func process(data []byte) error {
	if len(data) == 0 {
		return errors.New("Error:empty data")
	}
	return nil
}

//custom error

type myError struct{
	message string
} 

func(m *myError) Error()string{
	return fmt.Sprintf("error :%s",m.message)
}

func eprocess ()error{
	return &myError{"custom error message"}
}

func readData()error{
	if err :=readconfig(); err !=nil{
		return fmt.Errorf("readData %w",err)
	}
	return nil
}
 func readconfig()error{
	return errors.New("config error")
 }