package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/http2"
)

func main() {

	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		LogRequestDetails(r)
		fmt.Fprintf(w, "Handling incoming orders")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Handling  users")
	})

	port := 3000

	// load the TLS cert and key
	cert := "cert.pem"
	key := "key.pem"

	// configure tls
	tlsconfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	// now create a custom server
	server := &http.Server{
		Addr:      fmt.Sprintf(":%d", port),
		Handler:   nil,
		TLSConfig: tlsconfig,
	}

	// enable http2
	http2.ConfigureServer(server, &http2.Server{})

	fmt.Println("Our server is working on: ", port)

	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatal("Error making connection", err)
	}
	//    err :=	http.ListenAndServe(fmt.Sprintf(": %d",port),nil)
	//     if(err!=nil){
	// 		log.Fatal("Error listening port : ",err)
	// 	}
}

//  openssl req -x509 -nodes -newkey rsa:2048 -keyout key.pem -out cert.pem -days 365
// => is used for creating STL/TLS Certificates

func LogRequestDetails(r *http.Request) {
	httpversion := r.Proto
	fmt.Println("Received request with HTTP version", httpversion)

	if r.TLS !=nil{
		tlsversion :=getTLSVersionName(r.TLS.Version)
		fmt.Println("Receiverd Request with TLS vesion", tlsversion)
	}else{
		fmt.Println("Receuved request without tls")
	}
}

func getTLSVersionName(version uint16) string {
	switch version {
	case tls.VersionTLS10:
		return "tls version 1.0"
	case tls.VersionTLS11:
		return "tls version 1.1"
	case tls.VersionTLS12:
		return "tls version 1.2"
	case tls.VersionTLS13:
		return "tls version 1.3"
	default:
		return "unknown tls version"

	}
}
