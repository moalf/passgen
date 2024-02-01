package server

import (
	"fmt"
	"net/http"

	"github.com/moalf/passgen/api"
)

func NewHttpServer(address string, port int) *http.Server {
	http.HandleFunc("/status", api.LogRequest(api.Status))
	http.HandleFunc("/", api.LogRequest(api.GetPassword))

	listener := &http.Server{
		Addr: fmt.Sprintf("%s:%d", address, port),
	}

	return listener
}
