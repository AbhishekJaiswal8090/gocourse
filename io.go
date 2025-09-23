package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
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
	var buf bytes.Buffer
	buf.WriteString("Hello buffer")
	fmt.Println(buf.String())
}

func multiReaderExample(){
	r1:=strings.NewReader("Hello")
	r2:=strings.NewReader("world")
	mr:=io.MultiReader(r1,r2)
	buf:=new(bytes.Buffer)
}

func main() {

}