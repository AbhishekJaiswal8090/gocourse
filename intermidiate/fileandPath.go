package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func FilePath() {
	relativePath := "./data/file.txt"
	absolutePaath := "C:/Users/Abhi/Documents/notes.txt" // for Windows, use forward slashes or escape backslashes

	//joined path using filepath.join
	joinedpath := filepath.Join("downloads","file.zip")
	fmt.Println(joinedpath)

	normalizePath :=filepath.Clean("./data/../data/file.txt")

// normalized path
	fmt.Println(normalizePath)

	dir ,file :=filepath.Split("/home/user/docs/file.txt")
	fmt.Println(dir,file)

	//checking for whether path is absolute or not 

	fmt.Println(filepath.IsAbs(relativePath))
	fmt.Println(filepath.IsAbs(absolutePaath))

	fmt.Println(filepath.Ext(file))
	fmt.Println(strings.TrimSuffix(file,filepath.Ext(file)))


}