package main

import (
	"fmt"
	"log"
	"time"

	"github.com/moalf/passgen/server"
)

const (
	address = "0.0.0.0"
	port    = 8080
)

func main() {
	server := server.NewHttpServer(address, port)

	fmt.Printf("%s passgen started...\n", time.Now().Format("2006-01-02 15:04:05"))
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
