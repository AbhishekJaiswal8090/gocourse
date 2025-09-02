//string function in go
package main 

import ("fmt"
       "strings"
	    "strconv")

func main(){
  str := "Golang"
fmt.Println(str[0:4]) // "Gola"




fmt.Println(strings.Contains("golang", "go"))   // true
fmt.Println(strings.HasPrefix("golang", "go"))  // true
fmt.Println(strings.HasSuffix("golang", "ang")) // true


fmt.Println(strings.Index("golang", "lan"))   // 2
fmt.Println(strings.Count("golang", "g"))     // 2




fmt.Println(strings.Contains("golang", "go"))   // true
fmt.Println(strings.HasPrefix("golang", "go"))  // true
fmt.Println(strings.HasSuffix("golang", "ang")) // true


fmt.Println(strings.Index("golang", "lan"))   // 2
fmt.Println(strings.Count("golang", "g"))     // 2


str1 := "a,b,c"
parts := strings.Split(str1, ",")
fmt.Println(parts) // [a b c]

joined := strings.Join(parts, "-")
fmt.Println(joined) // a-b-c


str2 := "I love Java"
newStr := strings.Replace(str2, "Java", "Go", 1)
fmt.Println(newStr) // I love Go



fmt.Println(strings.ToUpper("golang")) // GOLANG
fmt.Println(strings.ToLower("GOLANG")) // golang


fmt.Println(strings.TrimSpace("   abhi   ")) // "abhi"
fmt.Println(strings.Trim("!!!hello!!!", "!")) // "hello"




// String → Int
i, _ := strconv.Atoi("123")
fmt.Println(i) // 123

// Int → String
s := strconv.Itoa(456)
fmt.Println(s) // "456"


//string formatting




}