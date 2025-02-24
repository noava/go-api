package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/noava/go-api/pollen"
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

	mux.Handle("/severity", rateLimitMiddleware(http.HandlerFunc(pollen.SeverityHandler)))

	log.Print("Listening...")

	if err := http.ListenAndServe(":5006", mux); err != nil {
		fmt.Println(err.Error())
	}
}