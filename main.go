package main

import (
	"fmt"
	"github.com/Superm4n97/book-server-backend/router"
	"net/http"
)

const (
	port = 8080
)

func main() {
	rt := router.Register()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: rt,
	}

	fmt.Println(fmt.Sprintf("starting server on port %d", port))
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to start the server. %s", err.Error()))
	}
}
