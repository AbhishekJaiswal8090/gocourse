package main 

import "fmt"



func main(){
	/*fmt.Println("Hello!")
	//starting generics in go
	hii,bro :=genercs("abhi","fuck")
	fmt.Println(hii,bro)
	hello("abhi","brotha","hello","you","hell")
	
	b,a := swap("abhi","akash")
	fmt.Println(b,a)
*/
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

}



func genercs[T any](a,b T)(T,T){
	return a,b
}
func hello[T any](a ...T)bool{
	for _,val :=range a{
        fmt.Println(val)
	}
	return true
}
func swap[T any](a,b T)(T,T){
	return b,a
}

//implementing stack in go
type Stack[T any] struct{
	elements []T
}

func(s *Stack[T]) push(element T)bool{
	s.elements=append(s.elements,element)
	fmt.Println("elements is pushed")
    return true
}

func (s *Stack [T]) pop()(T, bool){
	if len(s.elements)==0{
		var zero T
		fmt.Println("underflow condition")
		return zero,false
	}
	element := s.elements[len(s.elements)-1]
	s.elements=s.elements[:len(s.elements)-1]
	fmt.Println("element is popped")
    return element,true
}

func (s *Stack[T]) printElement(){
	fmt.Println(s.elements)
}

func(s *Stack[T]) isOkay()string{
    if len(s.elements)==0{
		return "none elements are pushed "
	}
	fmt.Println(s.elements)
	return "there are all element"
}