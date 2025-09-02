package main

import (
	"html/template"
	"os"
)

//text template

func templateFunc() {
	//temp1 := template.New("example")

	temp1, err := template.New("example").Parse("welcome ,{{.name}}! how are you")
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"name": "john",
	}
	err = temp1.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}



} 
