package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"golang.org/x/net/http2"
)

func main(){

	http.HandleFunc("/orders",func( w http.ResponseWriter,r *http.Request){
		fmt.Fprintf(w,"Handling incoming orders")
	})

	http.HandleFunc("/users",func(w http.ResponseWriter,r *http.Request){
		fmt.Fprintf(w,"Handling  users")
	})



	port:=3000;

	// load the TLS cert and key
	cert :="cert.pem"
	key:="key.pem"

	// configure tls
	tlsconfig :=&tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	// now create a custom server
	server:=&http.Server{
		Addr: fmt.Sprintf(":%d", port),
		Handler: nil,
		TLSConfig: tlsconfig,

	}

	// enable http2
	http2.ConfigureServer(server,&http2.Server{})


	fmt.Println("Our server is working on: ",port)

	err:=server.ListenAndServeTLS(cert,key)
	if(err!=nil){
		log.Fatal("Error making connection",err)
	}
//    err :=	http.ListenAndServe(fmt.Sprintf(": %d",port),nil)
//     if(err!=nil){
// 		log.Fatal("Error listening port : ",err)
// 	}
}
//  openssl req -x509 -nodes -newkey rsa:2048 -keyout key.pem -out cert.pem -days 365
// => is used for creating STL/TLS Certificates