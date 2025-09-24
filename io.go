package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func readFromReader(r io.Reader) {
     buf:=make([]byte,1024)
	 n,err:=r.Read(buf)
	 if err!=nil{
		log.Fatal("Error reading data ",err)
	 }
	 fmt.Println(string(buf[:n]))
}

func writeTowritter(w io.Writer,data string){
	_,err:=w.Write([]byte(data))
	if err!=nil{
		log.Fatal("Error writting data ",err)
	 }

}

func closeResource(c io.Closer){
	err:= c.Close()
	if err!=nil{
		log.Fatal("Error clossing resorces ",err)
	 }
}

func bufferExample(){
	var buf bytes.Buffer///stack {memory allocation}
	buf.WriteString("Hello buffer")
	fmt.Println(buf.String())
}

func multiReaderExample(){
	r1:=strings.NewReader("Hello")
	r2:=strings.NewReader("world")
	mr:=io.MultiReader(r1,r2)
	buf:=new(bytes.Buffer)//heap {memoery allocation}
    _,err:=buf.ReadFrom(mr)
	if err!=nil{
		log.Fatal("Error :Reading data")
	}
	fmt.Println(buf.String())
}

func pipeExample(){
	
}


func main() {

	fmt.Println("Read from Reader")
	file ,err:=os.Create("reader.txt")
	if err!=nil{
		fmt.Println("Error : Creating file")
		return
	}
	_,err=file.WriteString("Hello motherfucker")
	if err!=nil{
		fmt.Println("Error :Writting into the file")
		return
	}
	defer file.Close()

	readFile ,err:=os.Open("reader.txt")
	if err!=nil{
		fmt.Println("File opened successfully")
		return
	}
	readFromReader(readFile)
    defer readFile.Close()


	fmt.Println("Write to writter")

	var writter bytes.Buffer
	data :="hello motherfucker"
	writeTowritter(&writter,data)
	fmt.Println(writter.String())

	fmt.Println("===BUFFER EXAMPLE===")
	bufferExample()

	fmt.Println("===MultiReader Example===")
	multiReaderExample()



}