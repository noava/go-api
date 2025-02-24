package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const rateLimit = time.Second / 10 // per second

// Rate limiter
func rateLimitMiddleware(next http.Handler) http.Handler {
	limiter := time.NewTicker(rateLimit)
	
	requestChannel := make(chan struct{}, 1)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		select {
		case requestChannel <- struct{}{}:
			defer func() { <-requestChannel }()
			<-limiter.C
			next.ServeHTTP(w, r)
		default:
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
		}
	})
}

func main() {
	// Create a router using mux
	mux := http.NewServeMux()

	// GET /hello
	mux.Handle("GET /hello", rateLimitMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})))

	// GET /bye?name={name}
	mux.Handle("GET /bye", rateLimitMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "World"
		}
		fmt.Fprintf(w, "Bye, %s!", name)
	})))

	log.Print("Listening...")

	if err := http.ListenAndServe(":5006", mux); err != nil {
		fmt.Println(err.Error())
	}
}