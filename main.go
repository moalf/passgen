package main

import (
	"fmt"
	"log"

	"github.com/moalf/passgen/server"
)

const (
	address = "127.0.0.1"
	port    = 8080
)

func main() {
	server := server.NewHttpServer(address, port)

	fmt.Printf("Starting server on %s:%d\n", address, port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
