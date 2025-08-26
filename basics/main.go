package main

import (
	"errors"
	"fmt"
	"math"
)

func firstPart(){
	var abhi string
	abhi = "Hello i am abhi"
	fmt.Println(abhi)

	num := 450
	fmt.Println(num)

	const PI = 3.14

	fmt.Printf("the type of value that abhi holds is  %T ",abhi)
	fmt.Printf("The type of value that num holds is %T",num)
	fmt.Println(PI)
	}
	func secondPart(){
		var num int
		fmt.Scanln("Enter the num",&num)

		if num >0{
			fmt.Println("Num is positive",num)
		}else if num <0{
			fmt.Println("num is negative",num)
		}else{
			fmt.Println("num is zero",num)
		}

		if num % 2 != 0{
			fmt.Println("num is odd")
		}else{
			fmt.Println("num is even")
		}
	}



func thirdPart() {
    for i := 2; i <= 50; i++ {
        isPrime := true
        for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
            if i%j == 0 {
                isPrime = false
                break
            }
        }
        if isPrime {
            fmt.Println(i)
        }
    }
}
 func Sum(a , b int)int{
       return a+ b
 }
 func factorial(n int)int{
	if n == 0{
		return 1
	}
	return n*factorial(n-1)
 }
 func isPalindrome(n int)bool{
	original := n
	reversed := 0
	if n > 0{
		digit := n %10
		reversed = reversed*10 +digit
		n =n /10
	}
	return reversed == original
 }

func arrayPart(arr[5] int){
	//getting max value
	maxVal := arr[0]
	for i:=1; i<len(arr); i++{
		if arr[i] > maxVal{
			maxVal = arr[i]
		}
		
	}
	fmt.Println(maxVal)

	//getting minimum value
	minimumVal := arr[0]
	for i:=1; i<len(arr); i++{
		if arr[i] < minimumVal{
			minimumVal = arr[i]
		}
	}
	fmt.Println(minimumVal)

	//getting average value 
	total_sum := 0
	for i:=0; i<len(arr); i++{
		total_sum = total_sum + arr[i]
	}
	avg := total_sum / len(arr)
	fmt.Println(avg)
}
func slicePart(){
	slice1 := []int{1,2,3,4,5}
	slice1 = append(slice1, 10,20,30)
	fmt.Println(len(slice1),cap(slice1))
}

func mapPart(){
	map1 := map[string]int{
		"abhi": 60,
		"akash":70,
		"aman":80,
		"pawan":100,
		"priyank":300,
	}
//updating  one score
	map1["priyank"] =30
	//deleting one student
	delete(map1,"pawan")
	fmt.Println(map1)
}

func countWordFrequency(word string){
    freqMap := make(map[rune]int)

    for _, char := range word {
        freqMap[char]++
    }

    fmt.Println("Character Frequencies:")
    for char, count := range freqMap {
        fmt.Printf("%q: %d\n", char, count)
    }

}

//structPart 
type Shape interface{
	Area() float32
	Perimeter() float32
}
type Rectangle struct{
	width float32
	height float32
}

func(r Rectangle) Area() float32{
	return r.height*r.width
}

func(r Rectangle) Perimeter()float32{
	return 2 *(r.height + r.width)
}

type Circle struct{
	radius float32
}
func(c Circle) Area()float32{
	return c.radius*c.radius*math.Pi
}
func(c Circle) Perimeter()float32{
	return 2*math.Pi*c.radius
}
func calculateArea(s Shape) float32 {
	return s.Area()
}
func calculatePerimeter(s Shape)float32{
	return s.Perimeter()
}
//pointers 

func increament(x *int){
    *x =1
}

//creating a function which takes a function as a argument


func ApplyOperation(x, y float32, operation func(float32, float32) float32) float32 {
	return operation(x, y)
}
func Add(a,b float32)float32{
	return a+b
}
func Subtarct(a,b float32)float32{
	return a -b 
}

//craeting a function which return a function
func ReturnFunc(multiplier float32)func (float32)float32{
	return func (x float32)float32{
		return multiplier * x
	}

}

func Compare(x ,y int)(string ,error){
	if x > y{
		return "x is greater than y",nil
	}else if x < y{
		return "y is graeter than x",nil
	}else{
		return "",errors.New("something went wrong ")
	}
}
func main(){
	//firstPart()
   // secondPart()
   // thirdPart()
  // arr := [5]int{50,5,9,10,45}
  // arrayPart(arr)
  // slicePart()
//  mapPart()
//  countWordFrequency("Abhishek")
/*  rect1 :=Rectangle{
	40.5,
	50,
  }
  cir1 :=Circle{
	40,
 }*/ 
//  ans1 :=  calculateArea(rect1)
//  ans2 :=calculateArea(cir1)
//  ans3 :=calculatePerimeter(rect1)
//  fmt.Println(ans1,ans2,ans3)
  

//  x :=0
//  increament(&x)
//  fmt.Println(x)

//  addOperation :=ApplyOperation(50.78 ,56.8 ,Add)
//  fmt.Println(addOperation)

//  returnFunc :=ReturnFunc(50.56)
  //fmt.Println(returnFunc(5))

  res , err :=Compare(45 ,45)
  if err != nil{
	fmt.Println("Error",err)
  }else{
	fmt.Println(res)
  }

}