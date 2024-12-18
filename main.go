package main

import (
	"net/http"
)

func main() {
	api := &api{addr: ":8080"}

	// Initializa Mux router
	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", api.getUserHandler)
	mux.HandleFunc("POST /user", api.createUserHandler)

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
