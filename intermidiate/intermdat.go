package main

import (
	"fmt"
	"math"
)

// demosntrating closures

func intermdat() {

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

	/*message :="hELLO GO"
	randMessage :="hrllo \ngo"

	fmt.Println(message)
	fmt.Println(randMessage)
	*/

	//struct ing go

	//there are two different way of declaring struct variables
	person := Person{
		firstname: "Abhi",
		lastname:  "Jaiswal",
		age:       32,
		gender:    "male",
		address: Address{
			city:     "hetimpur",
			district: "deoria",
		},
		homecontact: HomePhone{
			home: "465738",
			cell: "563903",
		},
	}

	fmt.Println(person)
	//craeting instance of struct
	var person1 Person
	person1.firstname = "akash"
	person1.lastname = "jaiswal"
	person1.age = 32
	person1.gender = "male"
	person1.address.city = "kushinagr"
	person1.address.district = "rampur"
	person1.homecontact.cell = "23451"
	person1.homecontact.home = "56730"

	//fmt.Println(person1)

	//struct method call
	// person.structMethod()

	//method
	r := Rectangle{
		length: 4.4,
		width:  60,
	}
	r.Scale(10)
	//fmt.Println(r.Area())

	//m := Myint(-5)
	//fmt.Println(m.Ispositive())

	//using our interface concepts here
	c := Cat{}
	ans := makeSound(c)
	fmt.Println(ans)

	//implementing  another interafce
	rect := Rectangla{
		45.8,
		56.0,
	}
	cir := Circle{
		34.90,
	}
	sqr := Square{
		45.46,
	}

	rectArea := measureArea(rect)
	fmt.Println(rectArea)

	rectPeri := measurePerimeter(rect)
	fmt.Println(rectPeri)

	cirArea := measureArea(cir)
	fmt.Println(cirArea)

	cirPeri := measurePerimeter(cir)
	fmt.Println(cirPeri)

	sqrArea := measureArea(sqr)
	fmt.Println(sqrArea)

	sqrPeri := measurePerimeter(sqr)
	fmt.Println(sqrPeri)
    
	myPrinter(1,"abhi",56.90)
	value("abhi")
}

// a function that demonstartes tha closures
func adder() func() int {
	i := 0
	fmt.Println("previous value of i", i)
	return func() int {
		i++
		fmt.Println("add 1 to i ", i)
		return i
	}

}

// recursion in go
func factorial(x int) int {
	//base case factorial of 0 is 1
	if x == 0 {
		return 1
	}
	//recursive case
	return x * factorial(x-1)

}

func sumOfDigit(x int) int {
	sum := x
	if x == 0 {
		return sum
	}

	return sum + sumOfDigit(x-1)
}

// struct in go
type Person struct {
	firstname   string
	lastname    string
	age         int16
	gender      string
	address     Address
	homecontact HomePhone
}
type Address struct {
	city     string
	district string
}

type HomePhone struct {
	home string
	cell string
}

// defining strutc method
func (p Person) structMethod() {
	fmt.Println(p.firstname)
	fmt.Println(p.age)
}

func (p Person) changeAge() {
	p.age++
}

// methods in detail
type Rectangle struct {
	length float32
	width  float32
}

// method with value receiver
func (r Rectangle) Area() float32 {
	return r.length * r.width
}

// method with pointer receiver
func (r *Rectangle) Scale(factor float32) {
	r.length *= factor
	r.width *= factor
}

type Myint int

func (m Myint) Ispositive() bool {
	return m > 0
}

// implementing interface in go
// interfaces in java is a concept which is widely used for method receviver for diffrent types
type AnimalSound interface {
	Sound() string
}

type Cat struct{}

func (c Cat) Sound() string {
	return "meow - meow"
}

type Dog struct{}

func (d Dog) Sound() string {
	return "bhau-bhau"
}

type Crow struct{}

func (c Crow) Sound() string {
	return "kaww-kaww"
}
func makeSound(a AnimalSound) string {
	return a.Sound()
}

// implemeting another interface
type ShapeMeasure interface {
	Area() float32
	Perimeter() float32
}

type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) Perimeter() float32 {
	return 2 * math.Pi * c.radius
}

type Rectangla struct {
	length float32
	width  float32
}

func (r Rectangla) Area() float32 {
	return r.length * r.width
}

func (r Rectangla) Perimeter() float32 {
	return 2 * (r.length + r.width)
}

type Square struct {
	side float32
}

func (s Square) Area() float32 {
	return 4 * s.side * s.side
}

func (s Square) Perimeter() float32 {
	return 6 * s.side
}

func measureArea(s ShapeMeasure) float32 {
	return s.Area()
}

func measurePerimeter(s ShapeMeasure) float32 {
	return s.Perimeter()
}
func myPrinter(i ...interface{}){
	for _,val:= range i{
		fmt.Println(val)
	}
}

func value(i interface{}){
     fmt.Println(i)
	 switch i:=i.(type) {
	 case int:
		fmt.Println("the value is int")
	case string:
		fmt.Println("type is string")
	default:
		fmt.Println("unknown",i)	
	 }

}

