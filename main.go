package main

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/sony/gobreaker"
)

// Fake external service simulation
func unreliableService() error {
	randFloat := rand.Float32()
	fmt.Println("randFloat:", randFloat)
	if randFloat < 0.7 {
		return errors.New("service failed")
	}
	fmt.Println("Service succeeded")
	return nil
}

// Retry with exponential backoff
func retryWithBackoff(operation func() error, maxRetries int) error {
	var err error
	for i := range maxRetries {
		err = operation()
		if err == nil {
			return nil
		}
		wait := time.Duration((1 << i)) * time.Millisecond * 100 // Exponential: 100ms, 200ms, 400ms...
		fmt.Printf("Retry %d failed, waiting %v...\n", i+1, wait)
		time.Sleep(wait)
	}
	return fmt.Errorf("all retries failed: %w", err)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Define circuit breaker settings
	cbSettings := gobreaker.Settings{
		Name:        "UnreliableServiceCB",
		MaxRequests: 2, // Allowed requests in half-open state
		Interval:    0, // No rolling window reset
		Timeout:     5 * time.Second,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			return counts.ConsecutiveFailures >= 3
		},
	}

	breaker := gobreaker.NewCircuitBreaker(cbSettings)

	http.HandleFunc("/call", func(w http.ResponseWriter, r *http.Request) {
		err := retryWithBackoff(func() error {
			// Wrap the external service with the circuit breaker
			_, err := breaker.Execute(func() (interface{}, error) {
				return nil, unreliableService()
			})
			return err
		}, 3)

		if err != nil {
			http.Error(w, "Service unavailable: "+err.Error(), http.StatusServiceUnavailable)
			return
		}

		fmt.Fprintln(w, "Success!")
	})

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
