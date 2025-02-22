package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Create a router using mux
	mux := http.NewServeMux()

	// GET /hello
	mux.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	// GET /bye?name={name}
	mux.HandleFunc("GET /bye", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "World"
		}
		fmt.Fprintf(w, "Bye, %s!", name)
	})

	log.Print("Listening...")

	if err := http.ListenAndServe(":5006", mux); err != nil {
		fmt.Println(err.Error())
	}
}