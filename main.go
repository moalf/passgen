package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/moalf/passgen/api"
)

const port = "8080"

func main() {
	http.HandleFunc("/status", api.Status)
	http.HandleFunc("/", api.GetPassword)

	fmt.Printf("Starting server at port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
