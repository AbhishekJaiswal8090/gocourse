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
		elements: []int{45,78,500,34},
	}

	val :=s.push(45)
	fmt.Println(val)
	s.printElement()
	ele ,_:=s.pop()
	fmt.Println(ele)
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

func (s *Stack[T]) push(element T) string {
	 s.elements = append(s.elements, element)
	 return "element is push into the stack"
}

func (s *Stack[T]) pop() (T ,bool) {
	if len(s.elements) == 0{
		var zero T
		return zero ,false
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element,true
}

func (s *Stack[T]) isEmpty()bool{
	return len(s.elements)==0
}
func (s *Stack[T]) printElement(){
	fmt.Println(s.elements)
}