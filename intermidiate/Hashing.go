package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
)

func Hashing(){
	fmt.Println("Hashing")
	password :="password123"
	// fmt.Println(len(password))

	// hash :=sha256.Sum256([]byte(password))
    // // hash :=sha512.Sum512([]byte(password))

	// fmt.Println(hash)
	// fmt.Println(len(hash))
	// fmt.Printf("SHA - 512 Hash hex val :%x \n",hash)

	//salting
	
     salt ,err :=generateSalt()
	 if err!=nil{
	  fmt.Println(err)
	 }

	 hashpass :=hashPassword(password,salt)
	 fmt.Println(hashpass)

	 
    res :=signIn("abhi123" ,salt ,hashpass)
	fmt.Println(res)
}
func generateSalt()([]byte ,error){
	salt:=make([]byte,16)
	_ ,err:=io.ReadFull(rand.Reader,salt)
	if err !=nil{
		fmt.Println(err)
		return nil,err
	}
	return salt ,nil
}

//func to hash password
func hashPassword(password string,salt []byte)string{
	saltedPassword :=append(salt,[]byte(password)...)
	hash :=sha256.Sum256(saltedPassword)
	return base64.StdEncoding.EncodeToString(hash[:])
}

func signIn(password string, salt []byte, hashPass string) (interface{}) {
	salt_password := append(salt, []byte(password)...)
	hash := sha256.Sum256(salt_password)
	encodedHash := base64.StdEncoding.EncodeToString(hash[:])
	if  encodedHash == hashPass{
		fmt.Println("successful login")
	}else{
		return errors.New("Error :Password not match")
	}
	return true
	
}

