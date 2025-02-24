package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/noava/go-api/db"
	"github.com/noava/go-api/gardening"
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
	db.InitDB()
	db.SeedDB()

	// Create a router using mux
	mux := http.NewServeMux()

	mux.Handle("/severity", rateLimitMiddleware(http.HandlerFunc(pollen.SeverityHandler)))
	mux.Handle("/pollen-info", rateLimitMiddleware(http.HandlerFunc(pollen.PollenInfoHandler)))
	mux.Handle("/when-to-plant", rateLimitMiddleware(http.HandlerFunc(gardening.WhenToPlantHandler)))

	log.Print("Listening...")

	if err := http.ListenAndServe(":5006", mux); err != nil {
		fmt.Println(err.Error())
	}
}